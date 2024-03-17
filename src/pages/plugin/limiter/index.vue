<template>
    <div class="limiter">
        <Condition @search="search" />
        <div class="data">
            <el-table class="table" :border="true" :data="tableData">
                <el-table-column type="index" :align="'center'" />
                <el-table-column v-for="item in tableColumn" :prop="item.prop" :label="item.label" :align="'center'" />
                <el-table-column prop="content" label="操作" :align="'center'">
                    <template #default="scope">
                        <el-button size="small" @click="openEditDialog(scope)">编辑</el-button>
                        <el-popconfirm title="确认刷新容器吗?" @confirm="refreshLimiter(scope)">
                            <template #reference><el-button size="small">刷新容器</el-button></template>
                        </el-popconfirm>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination class="pagination" :current-page="pageNo" :page-size="pageSize" :total="totalNum"
                :layout="layout" @current-change="handlePageChange" />
        </div>
        <el-button class="saveBtn" :icon="Plus" type="primary" circle size="large" @click="showAddDialog" />
        <Save :show="showDialog" :flag="flagDialog" :data="limiterData" @reload="getLimiterByPage"/>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import Condition from "./condition.vue";
import Save from "./save.vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode, role } from "@/config";
import { useRoleStore } from "@/store";
import { Plus } from "@element-plus/icons-vue";
import { SaveDialog } from "@/constants";

let condition = {
    appName: "",
    limit: "",
};
const roleStore = useRoleStore();
const pageNo = ref(1);
const pageSize = ref(10);
const totalNum = ref(0);
const tableData = ref([]);
const layout = "prev, pager, next";
const tableColumn = [
    { prop: "appName", label: "应用名" },
    { prop: "limiterName", label: "限流器" },
    { prop: "mode", label: "方式" },
    { prop: "parameter", label: "参数配置" },
    { prop: "updatedAt", label: "更新时间" },
];
const showDialog = ref(0);
const flagDialog = ref();
const limiterData = ref();

// 处理请求参数，再获取通知数据
function search(inputValue: any) {
    condition.appName = inputValue.appName;
    condition.limit = inputValue.limit;
    getLimiterByPage();
}

// 分页获取限流器配置
function getLimiterByPage() {
    api.getLimiterByPage({
        ...condition,
        pageNo: pageNo.value,
        pageSize: pageSize.value
    }).then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用限流器信息获取失败：" + msg });
            return;
        }
        totalNum.value = data.total;
        tableData.value = data.data;
    });
}

// 页码变化
function handlePageChange(page: number) {
    pageNo.value = page;
}

// 打开新增弹窗
function showAddDialog() {
    flagDialog.value = SaveDialog.add;
    limiterData.value = null;
    showDialog.value++;
}

// 打开编辑弹窗
function openEditDialog(data: any) {
    flagDialog.value = SaveDialog.update;
    limiterData.value = data.row;
    showDialog.value++;
}

// 刷新限流容器
function refreshLimiter(data: any) {
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足" });
        return;
    }
}

onMounted(() => {
    getLimiterByPage();
});
</script>

<style lang="scss" scoped>
.limiter {
    .data {
        margin-bottom: 150px;

        .table {
            width: calc(100% - 30px);
            margin: 15px auto;
        }

        .pagination {
            position: absolute;
            right: 15px;
        }
    }

    .saveBtn {
        position: fixed;
        right: 25px;
        bottom: 25px;
    }
}
</style>