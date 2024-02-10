import { get, post, put, del } from "./ajax";

const user = {
    doLogin: async (req: any) => {
        return await post("/user/doLogin", req);
    },
    getUserInfo: async () => {
        return await get("/user/getUserInfo");
    },
}

const api = {
    ...user,
}

export default api;