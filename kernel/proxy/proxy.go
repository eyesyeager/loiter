package proxy

import (
	"fmt"
	"loiter/constants"
	"loiter/global"
	"loiter/kernel/balancer"
	"loiter/kernel/container"
	"loiter/utils"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
)

/**
 * @author eyesYeager
 * @date 2023/9/25 16:58
 */

func StartProxy() {
	// 制定路由规则
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		host := req.Host
		// 获取应用代理类型
		genre, ok := container.GenreByAppMap[host]
		if !ok {
			errMsg := fmt.Sprintf("the application with host %s has not been registered yet", host)
			statusCode, contentType, content := utils.HtmlSimpleTemplate(constants.ResponseTitle.BadGateway, errMsg)
			utils.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(errMsg)
			return
		}

		// 请求预处理
		if err, allow := pre(w, req, host, genre); !allow {
			if err == nil {
				post(w, req, nil, host, constants.PostEntrance.Reject, "", genre)
			} else {
				post(w, req, nil, host, constants.PostEntrance.Error, err.Error(), genre)
			}
			return
		}

		// 获取代理信息
		err, targetUrl := balancer.Entry(host)
		if err != nil {
			errMsg := fmt.Sprintf("the load balancing policy execution failed and the proxy could not be used. Error message: %s", err.Error())
			statusCode, contentType, content := utils.ResponseTemplate(constants.ResponseTitle.BadGateway, errMsg, genre)
			utils.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(errMsg)
			return
		}

		// 根据用户代理类型选择服务
		if genre == constants.AppGenre.Api {
			apiProxy(w, req, host, targetUrl)
		} else if genre == constants.AppGenre.Static {
			staticProxy(w, req, host, targetUrl)
		} else {
			errMsg := fmt.Sprintf("the application with host %s has illegal type: %s", host, genre)
			statusCode, contentType, content := utils.HtmlSimpleTemplate(constants.ResponseTitle.BadGateway, errMsg)
			utils.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(errMsg)
			return
		}
	})
}

// apiProxy Api服务代理
func apiProxy(w http.ResponseWriter, req *http.Request, host string, targetUrl string) {
	// 创建代理
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http", // 暂时只实现http代理
		Host:   targetUrl,
	})

	// 响应处理
	proxy.ModifyResponse = func(resp *http.Response) error {
		post(w, req, resp, host, constants.PostEntrance.Post, "", constants.AppGenre.Api)
		return nil
	}

	// 执行反向代理
	proxy.ServeHTTP(w, req)
}

// staticProxy 静态服务代理
func staticProxy(w http.ResponseWriter, req *http.Request, host string, targetUrl string) {
	// 构建请求路径
	buildPathErr, rootPath, fullPath := buildStaticPath(targetUrl, req)
	if buildPathErr != nil {
		statusCode, contentType, content := utils.HtmlSimpleTemplate(constants.ResponseTitle.BadGateway, buildPathErr.Error())
		utils.Response(w, statusCode, contentType, content)
		global.GatewayLogger.Warn(buildPathErr.Error())
		post(w, req, nil, host, constants.PostEntrance.Error, buildPathErr.Error(), constants.AppGenre.Static)
		return
	}
	// 资源嗅探
	if !checkStaticExist(fullPath) {
		// 获取重定向信息
		appStatic, ok := container.StaticByAppMap[host]
		if !ok {
			errMsg := fmt.Sprintf("no AppStatic container with host %s application found", host)
			statusCode, contentType, content := utils.HtmlSimpleTemplate(constants.ResponseTitle.BadGateway, errMsg)
			utils.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(errMsg)
			post(w, req, nil, host, constants.PostEntrance.Error, errMsg, constants.AppGenre.Static)
			return
		}
		// 若配置重定向信息，则按配置获取资源，否则返回404
		if appStatic.ErrorRoute == "" {
			errMsg := fmt.Sprintf("resource not found for application with host %s: %s", host, fullPath)
			statusCode, contentType, content := utils.HtmlSimpleTemplate(constants.ResponseTitle.NotFound, errMsg)
			utils.Response(w, statusCode, contentType, content)
			global.GatewayLogger.Warn(errMsg)
		} else {
			http.ServeFile(w, req, rootPath+string(filepath.Separator)+appStatic.ErrorRoute)
		}
		post(w, req, nil, host, constants.PostEntrance.Post, "", constants.AppGenre.Static)
		return
	}
	// 执行静态资源服务
	http.FileServer(http.Dir(rootPath)).ServeHTTP(w, req)
	post(w, req, nil, host, constants.PostEntrance.Post, "", constants.AppGenre.Static)
}
