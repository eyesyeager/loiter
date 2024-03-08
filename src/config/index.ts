// 网站相关配置
const site = {
    name: "Loiter",                                         // 网站名称
    slogen: "尘世闲游，无意逐鹿",                              // 网站标语
    github: "https://www.github.com/YuJiZhao/loiter",       // 项目Github仓库地址
    document: "http://space.eyescode.top",                  // 项目文档地址
    establishment: "2023",                                  // 项目成立年份
    author: "耶瞳",                                          // 项目作者

    usernameMaxLen: 10,  // 用户名最大长度
    usernameMinLen: 2,   // 用户名最小长度
    userPasMaxLen: 20,   // 用户密码最大长度
    userPasMinLen: 6,    // 用户密码最小长度
};

// 接口相关配置
const server = {
    baseUrl: "http://127.0.0.1:9500",   // 后端服务地址
    timeOut: 6 * 1000,                  // 请求超时时间(ms)
    token: "token",                     // 响应头中的令牌字段名
}

// 状态码
const responseCode = {
    success: 200,
    fail: 400,
    tooManyRequest: 429,
    badGateway: 502,
    authError: 20001,
}

// 角色信息
const role = {
    user: "user",
    admin: "admin",
    superAdmin: "super_admin"
}

// 本地存储
const storage = {
    token: "token"
}

export { site, server, responseCode, storage, role }