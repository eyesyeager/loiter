<template>
    <el-dialog class="static" v-model="dialogVisible" title="应用static配置项" width="500">
        <div>
            <div class="inputGroup">
                <span class="label">默认错误定向</span>
                <el-input class="input" v-model="inputValue.errorRoute" placeholder="默认错误定向" clearable />
            </div>
        </div>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="saveStaticApp">保存</el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, watch } from "vue";
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
const inputValue = reactive({
    errorRoute: ""
});

// 监听父组件传值变化，控制弹窗展示
watch(
    () => props.show,
    _ => {
        clearInputValue();
        getAppStaticInfoById();
        dialogVisible.value = true;
    }
);

// 获取应用static配置
function getAppStaticInfoById() {
    api.getAppStaticInfoById([props.appId]).then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "获取应用static配置失败：" + msg });
            return;
        }
        inputValue.errorRoute = data.errorRoute;
    });
}

// 新增/更新静态应用配置项
function saveStaticApp() {
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足！" });
        return;
    }
    api.saveStaticApp({ 
        ...inputValue,
        appId: props.appId ? Number(props.appId) : null
    }).then(({ code, msg }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "保存失败：" + msg });
            return;
        }
        dialogVisible.value = false;
        ElMessage({ type: "success", message: "保存成功" });
    });
}

// 清空表单值
function clearInputValue() {
    inputValue.errorRoute = "";
}
</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";

.static {
    .inputGroup {
        margin: 10px 0;
        display: flex;
        justify-content: space-between;

        .label {
            display: inline-block;
            width: 100px;
            height: 30px;
            text-align: right;
            line-height: 30px;
        }

        .input {
            width: 340px;
        }
    }
}
</style>