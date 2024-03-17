<template>
    <div class="edit">
        <el-dialog class="save" v-model="dialogVisible" title="编辑应用插件" width="480">
            <div>
                <div class="inputGroup">
                    <span class="label">应用名</span>
                    <el-input class="input" v-model="appName" disabled />
                </div>
                <div class="inputGroup" v-for="item in processorOptions">
                    <span class="label">{{ item.label }}</span>
                    <el-select class="input" v-model="item.model" multiple clearable>
                        <el-option v-for="option in item.instansts" :key="option.value" :label="option.label" :value="option.value" />
                    </el-select>
                </div>
            </div>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="dialogVisible = false">取消</el-button>
                    <el-button type="primary" @click="saveAppProcessor">保存</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode, role } from "@/config";
import { OptionsInterface } from "@/d.ts/common";
import { useRoleStore } from "@/store";

interface ProcessorOptionsInterface {
    label: string,
    value: string,
    instansts: OptionsInterface[],
    model: string[]
}

const emit = defineEmits(["reload"]);
const roleStore = useRoleStore();
const props = defineProps({
    show: Number,
    data: Object
});
const appName = ref("");
const dialogVisible = ref(false);
const processorOptions = ref<ProcessorOptionsInterface[]>([]);

// 监听父组件传值变化，控制弹窗展示
watch(
    () => props.show,
    _ => {
        // 绑定值
        appName.value = props.data!.appName;
        if (processorOptions.value != null) {
            processorOptions.value.forEach((v: any) => {
                v.model = props.data![v.value + "Code"];
            });
        }
        // 打开弹窗
        dialogVisible.value = true;
    }
);

// 获取处理器字典
function getProcessorDictionary() {
    api.getProcessorDictionary().then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "处理器字典获取失败：" + msg });
            return;
        }
        Promise.all(data.map((v: any) => getProcessorByGenre(v.value))).then(values => {
            values.forEach((v: any, i: number) => {
                if (v.code != responseCode.success) {
                    ElMessage({ type: "error", message: `获取类别为${v.value}的处理器失败：` + v.msg });
                    return;
                }
                let item = v.data;
                processorOptions.value.push({
                    label: data[i].label,
                    value: data[i].value,
                    instansts: item,
                    model:[]
                });
            });
        });
    });
}

// 按类别获取处理器
async function getProcessorByGenre(genre: string) {
    return await api.getProcessorByGenre([genre]);
}

// 更新应用处理器
function saveAppProcessor() {
    // 权限校验
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足！" });
        return;
    }
    // 构建参数
    let params: any = {
        appId: props.data!.appId,
    };
    processorOptions.value.forEach(v => {
        params[v.value] = v.model || [];
    });
    
    // 执行请求
    api.saveAppProcessor(params).then(({ code, msg }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "更新应用处理器失败：" + msg });
            return;
        }
        ElMessage({ type: "success", message: "更新应用处理器成功" });
        dialogVisible.value = false;
        emit("reload");
    });
}

onMounted(() => {
    getProcessorDictionary();
});

</script>

<style lang="scss" scoped>
.edit {
    .inputGroup {
        margin: 10px 0;
        display: flex;
        justify-content: space-between;

        .label {
            display: inline-block;
            width: 80px;
            height: 30px;
            text-align: right;
            line-height: 30px;
        }

        .input {
            width: 350px;
        }
    }
}
</style>