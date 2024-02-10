import { RouteRecordRaw } from 'vue-router';

const publicPath = {
    home: "/dashboard",
    dashboard: "/dashboard",
    application: "/application",
    plugin: "/plugin",
    message: "/message",
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
                path: publicPath.plugin,
                component: () => import("@/pages/plugin/index.vue"),
            },
            {
                path: publicPath.message,
                component: () => import("@/pages/message/index.vue"),
            },
            {
                path: publicPath.log,
                component: () => import("@/pages/log/index.vue"),
            },
        ]
    }
];

export { publicPath, content };