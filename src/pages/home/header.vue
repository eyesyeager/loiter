<template>
    <div class="header">
        <div class="logo"><span>{{ site.name }}</span></div>
        <div class="tool">
            <div class="github"><span @click="jumpSite(site.github)">Github</span></div>
            <div class="document"><span @click="jumpSite(site.document)">文档</span></div>
            <div class="user">
                <el-popconfirm title="是否退出登录?" @confirm="logout">
                    <template #reference> <span>{{ userName }}</span> </template>
                </el-popconfirm>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, watch } from "vue";
import { site } from "@/config";
import { jumpSite } from '@/utils/utils';
import { useUserStore } from "@/store";

const userStore = useUserStore();
const userName = ref("");

// 用户信息是异步获取的，因此需要添加监控
watch(
    () => userStore.$state.username,
    value => userName.value = value,
    { immediate: true }
);

function logout() {
    localStorage.clear();
    location.reload();
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
}
</style>