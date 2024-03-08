import progress from "@/utils/nprogress";
import { AxiosResponse } from "axios";
import { server, storage, responseCode } from "@/config";
import { publicPath } from "@/router/path";

export default (response: AxiosResponse<any, any>) => {
    // 关闭api进度条
    progress.close();
    // 刷新令牌
    refreshToken(response);
    // 特殊状态码处理
    let code = response.data.code;
    switch (code) {
        case responseCode.tooManyRequest: 
            requestLimit();
            return;
        case responseCode.badGateway:
            badGateway();
            return;
        case responseCode.authError:
            goLogin();
            return;
    }
}

// 刷新令牌
function refreshToken(response: AxiosResponse<any, any>) {
    let headers = response.headers;
    if (headers == null) {
        return;
    }
    let token = headers[server.token];
    if (token == null) {
        return;
    }
    localStorage.setItem(storage.token, token);
}

// 限制请求
function requestLimit() {
    alert("请求频率过高，请稍后再试")
}

// 网关异常
function badGateway() {
    alert("网关异常，请联系管理员处理")
}

// 权限异常，前往登录
function goLogin() {
    localStorage.removeItem(storage.token);
    window.location.replace("/#" + publicPath.common.login);
}