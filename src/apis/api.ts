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
    getAllApp: async () => {
        return await get("/app/getAllApp");
    },
    getAppInfoByPage: async (req: any) => {
        return await post("/app/getAppInfoByPage", req);
    },
    getAppInfoById: async (req: any) => {
        return await get("/app/getAppInfoById", req, UrlReqType.path);
    }
};

// 负载均衡相关接口
const balancer = {
    getAllBalancer: async () => {
        return await get("/balancer/getAllBalancer");
    },
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
    getNoticeDictionary: async () => {
        return await get("/common/getNoticeDictionary");
    },
    getRoleDictionary: async () => {
        return await get("/common/getRoleDictionary");
    },
    getProcessorDictionary: async () => {
        return await get("/common/getProcessorDictionary");
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
    refreshAppServer: async (req: any) => {
        return await get("/container/refreshAppServer", req, UrlReqType.path);
    },
    refreshBalancer: async (req: any) => {
        return await get("/container/refreshBalancer", req, UrlReqType.path);
    },
    refreshProcessor: async (req: any) => {
        return await get("/container/refreshProcessor", req, UrlReqType.path);
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