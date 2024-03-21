package proxy

import (
	"loiter/config"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/**
 * 静态资源代理工具
 * @author eyesYeager
 * @date 2024/3/21 11:29
 */

// buildStaticPath 构建完整文件路径
func buildStaticPath(targetUrl string, r *http.Request) (error, string, string) {
	dir, _ := os.Getwd()
	rootPath := dir + config.Program.StaticDirPath + targetUrl
	// 处理url
	filePath := path.Clean(r.URL.Path)
	if strings.HasSuffix(filePath, "/") {
		filePath += config.Program.StaticDefaultMainFile
	}
	filePath = strings.ReplaceAll(filePath, "/", string(filepath.Separator))
	return nil, rootPath, rootPath + filePath
}

// 校验静态资源是否存在
func checkStaticExist(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}
