import axios from "axios";
import { UrlReqType } from "@/constants";
import requestFilter from "./filter/requestFilter";
import responseFilter from "./filter/responseFilter";
import requestReject from "./reject/requestReject";
import responseReject from "./reject/responseReject";
import urlRequest from ".//help/urlHelp";
import { addHeader } from "./help/authHelp";
import { server } from "@/config";

const service = axios.create({
    baseURL: server.baseUrl,
    timeout: server.timeOut,
});

service.interceptors.request.use(
    config => {
        requestFilter(config);
        return config;
    },
    error => {
        requestReject(error);
        return Promise.reject(error);
    }
);

service.interceptors.response.use(
    response => {
        responseFilter(response);
        return response.data;
    },
    error => {
        responseReject(error);
        return Promise.reject(error);
    }
);

// 接口返回格式
interface RespInterface {
    code: number;
    msg: string;
    [propName: string]: any;
}

const get = (url: string, req?: any, type?: UrlReqType): Promise<RespInterface> => {
    return service.get(urlRequest[type || UrlReqType.param](url, req), addHeader());
}

const post = (url: string, req?: any): Promise<RespInterface> => {
    return service.post(url, req, addHeader());
}

const put = (url: string, req?: any): Promise<RespInterface> => {
    return service.put(url, req, addHeader());
}

const del = (url: string, req?: any, type?: UrlReqType): Promise<RespInterface> => {
    return service.delete(urlRequest[type || UrlReqType.param](url, req), addHeader());
}

export { get, post, put, del };