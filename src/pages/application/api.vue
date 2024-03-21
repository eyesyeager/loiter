<template>
    <el-dialog class="api" v-model="dialogVisible" title="应用api配置项" width="500">
        <div>
            暂无
        </div>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="saveApiApp">保存</el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode, role } from "@/config";
import { useRoleStore } from "@/store";

const props = defineProps({
    show: Number,
    appId: Number,
});
const roleStore = useRoleStore();
const dialogVisible = ref(false);

// 监听父组件传值变化，控制弹窗展示
watch(
    () => props.show,
    _ => {
        getAppApiInfoById();
        dialogVisible.value = true;
    }
);

// 获取应用api配置
function getAppApiInfoById() {
    api.getAppApiInfoById([props.appId]).then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "获取应用api配置失败：" + msg });
            return;
        }
        console.log(data);
    });
}

// 新增/更新动态应用配置项
function saveApiApp() {
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足！" });
        return;
    }
    ElMessage({ type: "success", message: "保存成功" });
    dialogVisible.value = false;
}
</script>