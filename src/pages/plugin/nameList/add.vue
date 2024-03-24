<template>
    <div class="add">
        <el-dialog class="add" v-model="dialogVisible" title="增加名单ip" width="480">
            <div>
                <div class="inputGroup">
                    <span class="label">应用名</span>
                    <el-select class="input" v-model="inputValue.appId" clearable>
                        <el-option v-for="option in appOptions" :key="option.value" :label="option.label" :value="option.value" />
                    </el-select>
                </div>
                <div class="inputGroup">
                    <span class="label">名单类型</span>
                    <el-select class="input" v-model="inputValue.genre" clearable>
                        <el-option v-for="option in nameListOptions" :key="option.value" :label="option.label" :value="option.value" />
                    </el-select>
                </div>
                <div class="inputGroup">
                    <span class="label">IP 地址</span>
                    <el-input class="input" type="textarea" v-model="inputValue.ipList" :rows="6" clearable placeholder="不同ip请换行区分"/>
                </div>
            </div>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="dialogVisible = false">取消</el-button>
                    <el-button type="primary" @click="addNameListIp">保存</el-button>
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
import { getListByWrap, checkIPv4 } from "@/utils/utils";

const emit = defineEmits(["reload"]);
const roleStore = useRoleStore();
const appOptions = reactive<OptionsInterface[]>([]);
const nameListOptions = reactive<OptionsInterface[]>([]);
const props = defineProps({
    show: Number,
});
const dialogVisible = ref(false);
const inputValue = reactive({
    appId: "",
    genre: "",
    ipList: ""
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

// 获取黑白名单字典
function getNameListDictionary() {
    api.getNameListDictionary().then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "黑白名单字典获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            nameListOptions[index] = {
                "label": item.label,
                "value": item.value,
            };
        });
    });
}

// 清空输入框
function clearInputValue() {
    inputValue.appId = "";
    inputValue.genre = "";
    inputValue.ipList = "";
}

// 更新应用黑白名单状态
function addNameListIp() {
    // 权限校验
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足！" });
        return;
    }
    // 参数校验
    if (inputValue.ipList == null) {
        ElMessage({ type: "error", message: "IP地址必填！" });
        return;
    }
    let ipList = getListByWrap(inputValue.ipList);
    for (let index in ipList) {
        if (!checkIPv4(ipList[index])) {
            ElMessage({ type: "error", message: "IP格式不对：" + ipList[index] });
            return;
        }
    }
    // 添加ip
    api.addNameListIp({
        appId: inputValue.appId ? Number(inputValue.appId) : null,
        genre: inputValue.genre,
        ipList: inputValue.ipList.split(/\r?\n/)
    }).then(({code, msg}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "添加IP失败：" + msg });
            return;
        }
        ElMessage({ type: "success", message: "IP添加成功" });
        dialogVisible.value = false;
        // TODO：通知父组件刷新IP列表
    });
}

onMounted(() => {
    getAppDictionary();
    getNameListDictionary();
});

</script>

<style lang="scss" scoped>
.add {
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