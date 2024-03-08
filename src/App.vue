<template>
    <router-view v-slot="{ Component }">
        <template v-if="Component">
            <component :is="Component" :v-if="show"/>
        </template>
    </router-view>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { publicPath } from "@/router/path";
import { storage, responseCode } from "@/config";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { useUserStore, useRoleStore } from "@/store";

const router = useRouter();
const userStore = useUserStore();
const roleStore = useRoleStore();
const show = ref(false);

router.isReady().then(() => {
    // 权限判断
    if (permissionMissJudge()) {
        showPage(publicPath.common.login);
        return;
    }
    // 获取用户信息
    getUserInfo();
    // 获取角色信息
    getRoleInfo();
});

// 判断是否无权限
function permissionMissJudge(): boolean {
    return localStorage.getItem(storage.token) == null;
}

// 获取用户信息
function getUserInfo() {
    api.getUserInfo().then(({code, msg, data}) => {
        if (code == responseCode.success) {
            userStore.init(data.uid, data.username, data.weight);
            showPage(publicPath.home);
        } else {
            ElMessage({ type: "error", message: msg });
            localStorage.removeItem(storage.token);
            showPage(publicPath.common.login);
        }
    });
}

// 获取角色信息
function getRoleInfo() {
    api.getRoleDictionary().then(({code, msg, data}) => {
        if (code == responseCode.success) {
            roleStore.init(data);
        } else {
            ElMessage({ type: "error", message: "角色信息获取失败！" + msg });
        }
    });
}

// 完成初始处理，开始展示页面
function showPage(target: string) {
    router.replace(target);
    show.value = true;
}
</script>