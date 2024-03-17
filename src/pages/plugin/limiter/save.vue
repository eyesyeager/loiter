<template>
    <el-dialog class="save" v-model="dialogVisible" :title="props.flag == SaveDialog.add ? '新增应用限流器' : '编辑应用限流器'" width="500">
        <div>
            <div class="inputGroup">
                <span class="label">应用名</span>
                <el-select class="input" v-model="inputValue.appName" :disabled="props.flag != SaveDialog.add" filterable clearable>
                    <el-option v-for="item in appOptions" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </div>
            <div class="inputGroup">
                <span class="label">限流器</span>
                <el-select class="input" v-model="inputValue.limiterCode" clearable @change="changeCode">
                    <el-option v-for="item in limiterOptions" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </div>
            <div class="inputGroup">
                <span class="label">限流模式</span>
                <el-select class="input" v-model="inputValue.mode" clearable>
                    <el-option v-for="item in limiterModeOptions" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </div>
            <div class="inputGroup">
                <span class="label">参数示例</span>
                <el-input class="input" v-model="parameterEx" type="textarea" autosize disabled />
            </div>
            <div class="inputGroup">
                <span class="label">限流参数</span>
                <el-input class="input" v-model="inputValue.parameter" type="textarea" autosize />
            </div>
        </div>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="saveAppLimiter">保存</el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, watch } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode, role } from "@/config";
import { OptionsInterface } from "@/d.ts/common";
import { useRoleStore } from "@/store";
import { SaveDialog } from "@/constants";

const emit = defineEmits(["reload"]);
const props = defineProps({
    show: Number,
    flag: Number,
    data: Object,
});
const roleStore = useRoleStore();
const dialogVisible = ref(false);
const inputValue = reactive({
    appName: "",
    limiterCode: "",
    mode: "",
    parameter: ""
});
const parameterEx = ref("");
const appOptions = reactive<OptionsInterface[]>([]);
const limiterOptions = reactive<OptionsInterface[]>([]);
const limiterModeOptions = reactive<OptionsInterface[]>([]);

// 监听父组件传值变化，控制弹窗展示
watch(
    () => props.show,
    _ => {
        clearInputValue();
        if (props.flag == SaveDialog.update) {
            initUpdateInputValue();
        }
        dialogVisible.value = true;
    }
);

// 获取所有应用
function getAppDictionary() {
    api.getAppDictionary().then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用信息获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            appOptions[index] = {
                "label": item.label,
                "value": item.value
            };
        });
    });
}

// 获取限流器字典
function getLimiterDictionary() {
    api.getLimiterDictionary().then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "限流器信息获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            limiterOptions[index] = {
                "label": item.label,
                "value": item.value,
                "appendix": item.appendix
            };
        });
    });
}

// 获取限流器模式字典
function getLimiterModeDictionary() {
    api.getLimiterModeDictionary().then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "限流器模式信息获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            limiterModeOptions[index] = {
                "label": item.label,
                "value": item.value
            };
        });
    });
}

// changeCode 限流器选择发生变化，同步更改参数示例
function changeCode(value: any) {
    for(let index in limiterOptions) {
        if (limiterOptions[index].value == value) {
            parameterEx.value = limiterOptions[index].appendix || "";
            return;
        }
    }
}

// 新增/更新应用限流器
function saveAppLimiter() {
    // 权限校验
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足！" });
        return;
    }
    let param = {
        
    };
}

// 清空输入框
function clearInputValue() {
    inputValue.appName = "";
    inputValue.limiterCode = "";
    inputValue.mode = "";
    inputValue.parameter = "";
    parameterEx.value = "";
}

// 初始化输入框
function initUpdateInputValue() {
    inputValue.appName = props.data!.appName;
    inputValue.limiterCode = props.data!.limiterCode;
    inputValue.mode = props.data!.mode;
    inputValue.parameter = props.data!.parameter;
    changeCode(props.data!.limiterCode);
}

onMounted(() => {
    getAppDictionary();
    getLimiterDictionary();
    getLimiterModeDictionary();
});

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";

.save {
    .inputGroup {
        margin: 10px 0;
        display: flex;
        justify-content: space-between;

        .label {
            display: inline-block;
            width: 60px;
            height: 30px;
            text-align: right;
            line-height: 30px;
        }

        .input {
            width: 380px;
        }
    }
}
</style>