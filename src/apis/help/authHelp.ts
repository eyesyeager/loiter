import { storage, server } from "@/config";

// 请求头处理
export let addHeader = (existent?: any) => {
    let config = {...existent};
    config.headers = config.headers || {};
    // 如果存在token，则携带
    let token = localStorage.getItem(storage.token);
    if (token != null) {
        config.headers[server.token] = token;
    }
    return config;
};