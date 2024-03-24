<template>
    <div class="switch">
        <el-dialog class="save" v-model="dialogVisible" title="切换名单状态" width="430">
            <div>
                <div class="inputGroup">
                    <span class="label">应用名</span>
                    <el-select class="input" v-model="inputValue.appId" clearable @change="getAppNameList">
                        <el-option v-for="option in appOptions" :key="option.value" :label="option.label" :value="option.value" />
                    </el-select>
                </div>
                <div class="switchGroup">
                    <div class="switchItem">
                        <span>黑名单：</span>
                        <el-switch v-model="inputValue.black" />
                    </div>
                    <div class="switchItem">
                        <span>白名单：</span>
                        <el-switch v-model="inputValue.white" />
                    </div>
                </div>
            </div>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="dialogVisible = false">取消</el-button>
                    <el-button type="primary" @click="updateAppNameList">保存</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, reactive } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode, role } from "@/config";
import { OptionsInterface } from "@/d.ts/common";
import { useRoleStore } from "@/store";

const emit = defineEmits(["reload"]);
const roleStore = useRoleStore();
const appOptions = reactive<OptionsInterface[]>([]);
const props = defineProps({
    show: Number,
});
const dialogVisible = ref(false);
const inputValue = reactive({
    appId: "",
    black: false,
    white: false
});

// 监听父组件传值变化，控制弹窗展示
watch(
    () => props.show,
    _ => {
        clearInputValue();
        dialogVisible.value = true;
    }
);

// 获取所有应用信息
function getAppDictionary() {
    api.getAppDictionary().then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用信息获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            appOptions[index] = {
                "label": item.label,
                "value": item.value,
            };
        });
    });
}

// 获取应用名单状态
function getAppNameList(appId: any) {
    if (!appId) {
        clearInputValue();
    }
    api.getAppNameList([appId]).then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用黑白名单状态获取失败：" + msg });
            return;
        }
        inputValue.black = data.black;
        inputValue.white = data.white;
    });
}

// 更新应用黑白名单状态
function updateAppNameList() {
    // 权限校验
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足！" });
        return;
    }
    if (!inputValue.appId) {
        ElMessage({ type: "error", message: "请选择目标应用！" });
        return;
    }
    api.updateAppNameList({
        ...inputValue,
        appId: inputValue.appId ? Number(inputValue.appId) : null
    }).then(({code, msg}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用黑白名单状态更新失败：" + msg });
            return;
        }
        ElMessage({ type: "success", message: "应用黑白名单状态更新成功！" });
        dialogVisible.value = false;
    });
}

// 清空输入框值
function clearInputValue() {
    inputValue.appId = "";
    inputValue.black = false;
    inputValue.white = false;
}

onMounted(() => {
    getAppDictionary();
});

</script>

<style lang="scss" scoped>
.switch {
    .inputGroup {
        margin: 10px 0;
        display: flex;
        justify-content: space-between;

        .label {
            display: inline-block;
            width: 50px;
            height: 30px;
            text-align: right;
            line-height: 30px;
        }

        .input {
            width: 330px;
        }
    }

    .switchGroup {
        display: flex;
        justify-content: center;
        height: 30px;
        .switchItem {
            width: 40%;
            text-align: center;
        }
    }
}
</style>