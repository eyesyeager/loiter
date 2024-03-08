<template>
    <div class="edit">
        <el-dialog class="save" v-model="dialogVisible" title="编辑负载均衡策略" width="400">
            <div>
                <div class="inputGroup">
                    <span class="label">应用名</span>
                    <el-input class="input" v-model="props.data!.appName" disabled />
                </div>
                <div class="inputGroup">
                    <span class="label">负载均衡</span>
                    <el-select class="input" v-model="balancer" clearable placeholder="请选择">
                        <el-option v-for="item in balancerOptions" :key="item.value" :label="item.label" :value="item.value" />
                    </el-select>
                </div>
            </div>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="dialogVisible = false">取消</el-button>
                    <el-button type="primary" @click="updateAppBalancer">保存</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script setup lang="ts">
import { ref, watch, reactive, onMounted } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode, role } from "@/config";
import { OptionsInterface } from "@/d.ts/common";
import { useRoleStore } from "@/store";

const emit = defineEmits(["reload"]);
const roleStore = useRoleStore();
const props = defineProps({
    show: Number,
    data: Object
});
const dialogVisible = ref(false);
const balancerOptions = reactive<OptionsInterface[]>([]);
const balancer = ref();

// 监听父组件传值变化，控制弹窗展示
watch(
    () => props.show,
    _ => {
        balancer.value = props.data!.balancerCode;
        dialogVisible.value = true;
    }
);

// 获取状态字典
function getAllBalancer() {
    api.getAllBalancer().then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "负载均衡策略获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            balancerOptions[index] = item;
        });
    });
}

// 更新应用负载均衡策略
function updateAppBalancer() {
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足！" });
        return;
    }
    api.updateAppBalancer({
        appId: props.data!.id,
        balancer: balancer.value
    }).then(({ code, msg }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用负载均衡策略更新失败：" + msg });
            return;
        }
        dialogVisible.value = false;
        emit("reload");
    });
}

onMounted(() => {
    getAllBalancer();
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
            width: 60px;
            height: 30px;
            text-align: right;
            line-height: 30px;
        }

        .input {
            width: 290px;
            text-align: right;
        }
    }
}
</style>