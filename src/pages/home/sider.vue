<template>
    <el-menu :default-active="publicPath.dashboard" router>
        <el-menu-item :index="publicPath.dashboard">
            <el-icon><DataAnalysis /></el-icon>
            <span>总控台</span>
        </el-menu-item>
        <el-menu-item :index="publicPath.application">
            <el-icon><Compass /></el-icon>
            <span>应用管理</span>
        </el-menu-item>
        <el-menu-item :index="publicPath.plugin">
            <el-icon><HelpFilled /></el-icon>
            <span>插件管理</span>
        </el-menu-item>
        <el-menu-item :index="publicPath.message">
            <el-icon><Bell /></el-icon>
            <span>消息通知</span>
        </el-menu-item>
        <el-menu-item :index="publicPath.log">
            <el-icon><Document /></el-icon>
            <span>操作日志</span>
        </el-menu-item>
    </el-menu>
</template>
  
<script lang="ts" setup>
import {
    DataAnalysis,
    Compass,
    HelpFilled,
    Bell,
    Document,
} from '@element-plus/icons-vue';
import { publicPath } from '@/router/path';
</script>
  