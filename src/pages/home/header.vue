<template>
    <div class="header">
        <div class="logo"><span @click="jumpHome">{{ site.name }}</span></div>
        <div class="tool">
            <div class="github"><span @click="jumpSite(site.github)">Github</span></div>
            <div class="document"><span @click="jumpSite(site.document)">文档</span></div>
            <div class="user"><span @click="openUserInfo">{{ "耶瞳" }}</span></div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { site } from "@/config";
import { useRouter } from 'vue-router';
import { publicPath } from '@/router/path';
import { jumpSite } from '@/utils/utils';

const router = useRouter();

// 点击 logo，跳转到首页
function jumpHome() {
    router.push(publicPath.home);
}

// 打开用户信息窗
function openUserInfo() {

}

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";

.header {
    background-color: $baseBackground;
    border-bottom: 1px solid $baseBorder;
    display: flex;
    justify-content: space-between;
    .logo {
        width: 150px;
        text-align: center;
        line-height: 60px;
        span {
            font-size: 24px;
            color: $primaryText;
            cursor: pointer;
        }
    }
    .tool {
        display: flex;
        div {
            width: 80px;
            text-align: center;
            line-height: 60px;
            span {
                color: $regularText;
                cursor: pointer;
            }
        }
    }
}</style>