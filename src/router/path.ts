import { RouteRecordRaw } from "vue-router";

const publicPath = {
    home: "/dashboard",
    dashboard: "/dashboard",
    application: "/application",
    balancer: "/balancer",
    plugin: {
        manage: "/plugin/manage",
        limiter: "/plugin/limiter",
        nameList: "/plugin/nameList",
    },
    notice: "/notice",
    log: "/log",
    common: {
        login: "/login",
        error: "/error",
    }
}

const common: Array<RouteRecordRaw> = [
    {
        path: publicPath.common.login,
        component: () => import("@/pages/login/index.vue"),
    },
    {
        path: publicPath.common.error,
        component: () => import("@/pages/error/index.vue"),
    },
    {
        path: "/:catchAll(.*)",
        redirect: publicPath.common.error
    }
];

const content: Array<RouteRecordRaw> = [
    ...common,
    {
        path: "",
        component: () => import("@/pages/home/index.vue"),
        children: [
            {
                path: publicPath.dashboard,
                component: () => import("@/pages/dashboard/index.vue"),
            },
            {
                path: publicPath.application,
                component: () => import("@/pages/application/index.vue"),
            },
            {
                path: publicPath.balancer,
                component: () => import("@/pages/balancer/index.vue"),
            },
            {
                path: publicPath.plugin.manage,
                component: () => import("@/pages/plugin/manage/index.vue"),
            },
            {
                path: publicPath.plugin.limiter,
                component: () => import("@/pages/plugin/limiter/index.vue"),
            },
            {
                path: publicPath.plugin.nameList,
                component: () => import("@/pages/plugin/nameList/index.vue"),
            },
            {
                path: publicPath.notice,
                component: () => import("@/pages/notice/index.vue"),
            },
            {
                path: publicPath.log,
                component: () => import("@/pages/log/index.vue"),
            },
        ]
    },
];

export { publicPath, content };