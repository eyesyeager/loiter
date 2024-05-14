import { get, post, put, del } from "./ajax";
import { UrlReqType } from "@/constants";

// 用户相关接口
const user = {
    doLogin: async (req: any) => {
        return await post("/user/doLogin", req);
    },
    getUserInfo: async () => {
        return await get("/user/getUserInfo");
    },
    getAllUser: async () => {
        return await get("/user/getAllUser");
    }
};

// 应用相关接口
const app = {
    saveApp: async (req: any) => {
        return await post("/app/saveApp", req);
    },
    activateApp: async (req: any) => {
        return await post("/app/activateApp", req);
    },
    deleteApp: async (req: any) => {
        return await post("/app/deleteApp", req);
    },
    getAppInfoByPage: async (req: any) => {
        return await post("/app/getAppInfoByPage", req);
    },
    getAppInfoById: async (req: any) => {
        return await get("/app/getAppInfoById", req, UrlReqType.path);
    },
    getAppApiInfoById: async (req: any) => {
        return await get("/app/getAppApiInfoById", req, UrlReqType.path);
    },
    getAppStaticInfoById: async (req: any) => {
        return await get("/app/getAppStaticInfoById", req, UrlReqType.path);
    },
    saveStaticApp: async (req: any) => {
        return await post("/app/saveStaticApp", req);
    },
};

// 负载均衡相关接口
const balancer = {
    getBalancerByPage: async (req: any) => {
        return await post("/balancer/getBalancerByPage", req);
    },
    updateAppBalancer: async (req: any) => {
        return await post("/balancer/updateAppBalancer", req);
    }
}

// 处理器相关接口
const processor = {
    // --------------------------- processor ------------------------------
    getProcessorByPage: async (req: any) => {
        return await post("/processor/getProcessorByPage", req);
    },
    saveAppProcessor: async (req: any) => {
        return await post("/processor/saveAppProcessor", req);
    },
    getProcessorByGenre: async (req: any) => {
        return await get("/processor/getProcessorByGenre", req, UrlReqType.path);
    },
    // --------------------------- limiter ------------------------------
    getLimiterByPage: async (req: any) => {
        return await post("/processor/getLimiterByPage", req);
    },
    saveAppLimiter: async (req: any) => {
        return await post("/processor/saveAppLimiter", req);
    },
    deleteAppLimiter: async (req: any) => {
        return await post("/processor/deleteAppLimiter", req);
    },
    // --------------------------- namelist ------------------------------
    getAppNameList: async (req: any) => {
        return await get("/processor/getAppNameList", req, UrlReqType.path);
    },
    updateAppNameList: async (req: any) => {
        return await post("/processor/updateAppNameList", req);
    },
    addNameListIp: async (req: any) => {
        return await post("/processor/addNameListIp", req);
    },
    // --------------------------- requestLog ------------------------------
    getOverviewRequestLog: async () => {
        return await get("/processor/getOverviewRequestLog");
    },
    getDetailedRequestExtremumLog: async (req: any) => {
        return await post("/processor/getDetailedRequestExtremumLog", req);
    },
    getDetailedRequestNumLog: async (req: any) => {
        return await post("/processor/getDetailedRequestNumLog", req);
    },
    getDetailedRequestRuntimeLog: async (req: any) => {
        return await post("/processor/getDetailedRequestRuntimeLog", req);
    },
    getDetailedRequestQPSLog: async (req: any) => {
        return await post("/processor/getDetailedRequestQPSLog", req);
    },
    getDetailedRequestVisitorLog: async (req: any) => {
        return await post("/processor/getDetailedRequestVisitorLog", req);
    },
    getDetailedRequestTopApiLog: async (req: any) => {
        return await post("/processor/getDetailedRequestTopApiLog", req);
    },
    getDetailedRequestRejectLog: async (req: any) => {
        return await post("/processor/getDetailedRequestRejectLog", req);
    }
};

// 日志相关接口
const log = {
    getUniversalLog: async (req: any) => {
        return await post("/log/getUniversalLog", req);
    },
};

// 通用接口
const common = {
    getStatusDictionary: async () => {
        return await get("/common/getStatusDictionary");
    },
    getAppDictionary: async () => {
        return await get("/common/getAppDictionary");
    },
    getBalancerDictionary: async () => {
        return await get("/common/getBalancerDictionary");
    },
    getNoticeDictionary: async () => {
        return await get("/common/getNoticeDictionary");
    },
    getRoleDictionary: async () => {
        return await get("/common/getRoleDictionary");
    },
    getProcessorDictionary: async () => {
        return await get("/common/getProcessorDictionary");
    },
    getAppGenreDictionary: async () => {
        return await get("/common/getAppGenreDictionary");
    },
    getLimiterDictionary: async () => {
        return await get("/common/getLimiterDictionary");
    },
    getLimiterModeDictionary: async () => {
        return await get("/common/getLimiterModeDictionary");
    },
    getNameListDictionary: async () => {
        return await get("/common/getNameListDictionary");
    }
}

// 通知接口
const notice = {
    getNoticeList: async (req: any) => {
        return await post("/notice/getNoticeList", req);
    },
    getEmailNoticeContent: async (req: any) => {
        return await get("/notice/getEmailNoticeContent", req, UrlReqType.path);
    }
}

// 容器接口
const container = {
    refreshAppContainer: async (req: any) => {
        return await get("/container/refreshAppContainer", req, UrlReqType.path);
    },
    refreshBalancer: async (req: any) => {
        return await get("/container/refreshBalancer", req, UrlReqType.path);
    },
    refreshProcessor: async (req: any) => {
        return await get("/container/refreshProcessor", req, UrlReqType.path);
    },
    refreshLimiter: async (req: any) => {
        return await get("/container/refreshLimiter", req, UrlReqType.path);
    }
}

// 接口整合，并对外暴露
const api = {
    ...user,
    ...app,
    ...balancer,
    ...processor,
    ...log,
    ...common,
    ...notice,
    ...container,
};

export default api;