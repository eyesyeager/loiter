<template>
    <el-menu :default-active="publicPath.dashboard" router>
        <template v-for="item in menuData">
            <el-menu-item :index="item.path" v-if="item.children == null">
                <el-icon><component :is="item.icon" /></el-icon>
                <span>{{ item.name }}</span>
            </el-menu-item>
            <el-sub-menu :index="item.name" v-else>
                <template #title>
                    <el-icon><component :is="item.icon" /></el-icon>
                    <span>{{ item.name }}</span>
                </template>
                <el-menu-item v-for="i in item.children" :index="i.path">{{ i.name }}</el-menu-item>
            </el-sub-menu>
        </template>
    </el-menu>
</template>
  
<script lang="ts" setup>
import { DataAnalysis, Compass, Rank, Bell, Help, Document } from '@element-plus/icons-vue';
import { publicPath } from '@/router/path';

const menuData = [
    {
        name: "总控台",
        path: publicPath.dashboard,
        icon: DataAnalysis
    },
    {
        name: "应用管理",
        path: publicPath.application,
        icon: Compass
    },
    {
        name: "负载均衡",
        path: publicPath.balancer,
        icon: Rank
    },
    {
        name: "网关插件",
        icon: Help,
        children: [
            {
                name: "插件管理",
                path: publicPath.plugin.manage,
            },
            {
                name: "应用限流",
                path: publicPath.plugin.limiter,
            },
            {
                name: "黑白名单",
                path: publicPath.plugin.nameList,
            },
        ]
    },
    {
        name: "消息通知",
        path: publicPath.notice,
        icon: Bell
    },
    {
        name: "操作日志",
        path: publicPath.log,
        icon: Document
    },
];

</script>
  