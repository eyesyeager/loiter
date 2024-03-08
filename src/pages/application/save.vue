<template>
    <el-dialog class="save" v-model="dialogVisible" :title="props.flag == SaveAppDialog.add ? '新增应用' : '编辑应用'" width="500">
        <div>
            <div class="inputGroup">
                <span class="label">应用名</span>
                <el-input class="input" v-model="inputValue.appName" clearable />
            </div>
            <div class="inputGroup">
                <span class="label">代理地址</span>
                <el-input class="input" v-model="inputValue.host" clearable>
                    <template #prepend>http://</template>
                </el-input>
            </div>
            <div class="inputGroup">
                <span class="label">责任人</span>
                <el-select class="input" v-model="inputValue.ownerId" clearable placeholder="请选择">
                    <el-option v-for="item in userOptions" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </div>
            <div class="inputGroup">
                <span class="label">代理实例</span>
                <div class="input">
                    <el-input class="serverInput" v-for="(_, index) in inputValue.serverList" 
                        v-model="inputValue.serverList[index].address" placeholder="地址" clearable>
                        <template #prepend>
                            <el-input class="weightInput" v-model="inputValue.serverList[index].weight" placeholder="权重" clearable />
                        </template>
                        <template #append>
                            <el-button :icon="Minus" circle size="small" @click="deleteServer(index)" />
                        </template>
                    </el-input>
                </div>
            </div>
            <div class="addServer">
                <el-button :icon="Plus" circle size="small" @click="addServer" />
            </div>
            <div class="inputGroup">
                <span class="label">备注</span>
                <el-input class="input" type="textarea" v-model="inputValue.remarks" autosize clearable />
            </div>
        </div>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="saveApp">保存</el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, watch } from "vue";
import { Plus, Minus } from "@element-plus/icons-vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode, role } from "@/config";
import { OptionsInterface } from "@/d.ts/common";
import { useRoleStore } from "@/store";
import { SaveAppDialog } from "@/constants";

interface ServerInterface {
    address: string,
    weight: number | null
}

interface InputValueInterface {
    appId: number,
    appName: string,
    host: string,
    ownerId: number | null | string,
    serverList: ServerInterface[],
    remarks: string
}

const emit = defineEmits(["reload"]);
const props = defineProps({
    show: Number,
    flag: Number,
    appId: Number,
});
const roleStore = useRoleStore();
const dialogVisible = ref(false);
const userOptions = reactive<OptionsInterface[]>([]);
const inputValue = reactive<InputValueInterface>({
    appId: 0,
    appName: "",
    host: "",
    serverList: [{
        address: "",
        weight: null
    }],
    ownerId: null,
    remarks: ""
});

// 监听父组件传值变化，控制弹窗展示
watch(
    () => props.show,
    _ => {
        clearCondition();
        if (props.flag == SaveAppDialog.update) {
            inputValue.appId = props.appId!;
            initUpdateCondition();
        }
        dialogVisible.value = true;
    }
);

// 增加服务实例数量
function addServer() {
    inputValue.serverList.push({
        address: "",
        weight: null
    });
}

// 删除指定服务实例
function deleteServer(index: number) {
    if (inputValue.serverList.length <= 1) {
        ElMessage({ type: "warning", message: "服务实例不能为空" });
        return;
    }
    inputValue.serverList.splice(index, 1);
}

// 获取所有可选用户
function getAllUser() {
    api.getAllUser().then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "用户信息获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            userOptions[index] = {
                "label": item.label,
                "value": item.value
            };
        });
    });
}

// 新增/更新应用
function saveApp() {
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足！" });
        return;
    }
    for (let index in inputValue.serverList) {
        let weight = Number(inputValue.serverList[index].weight);
        if (weight == null && weight <= 0) {
            ElMessage({ type: "error", message: "权重值必填且不能小于1" });
            return;
        }
        inputValue.serverList[index].weight = weight;
    }
    inputValue.ownerId = Number(inputValue.ownerId);
    api.saveApp({ ...inputValue }).then(({ code, msg }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "保存失败：" + msg });
            return;
        }
        dialogVisible.value = false;
        emit("reload");
    });
}

// 清空条件
function clearCondition() {
    inputValue.appId = 0;
    inputValue.appName = "";
    inputValue.host = "";
    inputValue.serverList = [{
        address: "",
        weight: null
    }];
    inputValue.ownerId = null;
    inputValue.remarks = "";
}

// 初始化编辑条件
function initUpdateCondition() {
    api.getAppInfoById([props.appId]).then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "获取应用详细信息失败" + msg });
            return;
        }
        inputValue.appName = data.appName;
        inputValue.host = data.host;
        inputValue.serverList = data.serverList;
        inputValue.ownerId = String(data.ownerId);
        inputValue.remarks = data.remarks;
    });
}

onMounted(() => {
    getAllUser();
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
            text-align: right;
        }

        .serverInput {
            margin-bottom: 10px;
            width: 380px;

            .weightInput {
                width: 60px;
            }
        }
    }

    .addServer {
        text-align: right;
        transform: translateY(-10px);
    }
}
</style>