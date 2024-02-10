// 网站相关配置
const site = {
    name: "Loiter",                                         // 网站名称
    slogen: "尘世闲游，无意逐鹿",                              // 网站标语
    github: "https://www.github.com/YuJiZhao/loiter",       // 项目Github仓库地址
    document: "http://space.eyescode.top",                  // 项目文档地址
};

// 接口相关配置
const server = {
    baseUrl: "http://127.0.0.1:9510",   // 后端服务地址
    timeOut: 6 * 1000,                  // 请求超时时间(ms)
}

// 状态码
const responseCode = {
    success: 200,
    fail: 400,
    tooManyRequest: 429,
    badGateway: 502,
}

export { site, server, responseCode }