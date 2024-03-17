<template>
    <div class="appCard">
        <div class="header">
            <div class="title">{{ props.data!.appName + "（" + props.data!.host + "）" }}</div>
        </div>
        <div class="body">
            <div class="line">
                <div>应用类型: {{ props.data!.appGenre }}</div>
                <div>服务实例: {{ props.data!.validServerNum + "/" + props.data!.serverNum }}</div>
            </div>
            <div class="line">
                <div>插件数: {{ props.data!.plugins }}</div>
                <div>责任人: {{ props.data!.owner }}</div>
            </div>
            <div class="line">
                <div>状态: {{ props.data!.status }}</div>
                <div>创建时间: {{ props.data!.createdAt }}</div>
            </div>
            <div class="remarks">
                备注: {{ props.data!.remarks }}
            </div>
        </div>
        <div class="footer">
            <el-popconfirm title="确认刷新容器吗?" @confirm="refreshAppContainer">
                <template #reference><el-button class="appBtn" plain>刷新容器</el-button></template>
            </el-popconfirm>
            <el-popconfirm title="确认应用状态吗?" @confirm="activateApp">
                <template #reference><el-button class="appBtn" plain>切换状态</el-button></template>
            </el-popconfirm>
            <el-button class="appBtn" plain @click="updateApp">编辑</el-button>
            <el-popconfirm title="确认删除该应用吗?" @confirm="deleteApp">
                <template #reference><el-button class="appBtn" plain>删除</el-button></template>
            </el-popconfirm>
        </div>
    </div>
</template>

<script setup lang="ts">
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode, role } from "@/config";
import { useRoleStore } from "@/store";

const emit = defineEmits(["reload", "activate", "edit"]);
const props = defineProps({
    data: Object,
});
const roleStore = useRoleStore();

// 刷新应用服务实例容器
function refreshAppContainer() {
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足" });
        return;
    }
    api.refreshAppContainer([props.data!.appId]).then(({code, msg}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用服务实例容器刷新失败：" + msg });
            return;
        }
        ElMessage({ type: "success", message: "应用服务实例容器刷新成功" });
    });
}

// 激活/使失效 app
function activateApp() {
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足" });
        return;
    }
    api.activateApp({ appId: props.data!.appId }).then(({ code, msg }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用状态变更失败：" + msg });
            return;
        }
        emit("reload");
    });
}

// 编辑应用
function updateApp() {
    emit("edit", props.data!.appId);
}

// 删除应用
function deleteApp() {
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足" });
        return;
    }
    api.deleteApp({ appId: props.data!.appId }).then(({ code, msg }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "删除应用失败：" + msg });
            return;
        }
        emit("reload");
    });
}

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";
@import "@/assets/css/size.scss";

.appCard {
    background: $baseBackground;
    border-radius: $cardBorder;

    .header {
        height: 40px;
        color: $primaryText;
        line-height: 40px;

        .title {
            font-size: 18px;
            text-align: center;
        }
    }

    .body {
        color: $regularText;
        border-bottom: 1px solid $baseBorder;
        padding: 15px;

        .line {
            display: flex;
            justify-content: space-between;

            div {
                width: 50%;
                height: 25px;
                line-height: 25px;
            }
        }

        .remarks {
            line-height: 25px;
        }
    }

    .footer {
        height: 50px;
        line-height: 50px;
        text-align: right;

        &:last-child {
            margin-right: 12px;
        }
    }
}
</style>