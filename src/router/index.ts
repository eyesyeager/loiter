import { createRouter, RouteRecordRaw, createWebHashHistory } from 'vue-router';
import { publicPath, content } from './path';

// 定义路由规则
const routes: Array<RouteRecordRaw> = [
    {
        path: publicPath.home,
        component: () => import("@/App.vue"),
        children: content
    },
]

// 创建路由
const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

// 前置拦截器
router.beforeEach((to, from, next) => {
    next();
});

// 后置拦截器
router.afterEach((to, from) => {
});

export default router;