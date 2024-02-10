// 请求时，url携带参数类型
export enum UrlReqType {
    "param",  // 拼接在 ? 后的参数
    "path",   // 拼接在路径中的参数
    "mix",    // 以上两种方式的混合
}