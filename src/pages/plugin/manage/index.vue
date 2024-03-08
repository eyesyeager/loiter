<template>
    <div class="manage">
        <Condition @search="search"/>
        <div class="data">
            <el-table class="table" :border="true" :data="tableData">
                <el-table-column type="index" :align="'center'" />
                <el-table-column v-for="item in tableColumn" :prop="item.prop" :label="item.label" :align="'center'" />
                <el-table-column prop="content" label="操作" :align="'center'">
                    <template #default="scope">
                        <el-button size="small" @click="openEditDialog(scope)">编辑</el-button>
                        <el-popconfirm title="确认刷新容器吗?" @confirm="refreshProcessor(scope)">
                            <template #reference><el-button size="small">刷新容器</el-button></template>
                        </el-popconfirm>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination class="pagination" :current-page="pageNo" :page-size="pageSize" :total="totalNum"
                :layout="layout" @current-change="handlePageChange" />
        </div>
        <Edit :show="showDialog" :data="processorData" @reload="getProcessorByPage" />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive } from "vue";
import Condition from "./condition.vue";
import Edit from "./edit.vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode, role } from "@/config";
import { useRoleStore } from "@/store";

let condition = {
    appName: "",
};
const roleStore = useRoleStore();
const pageNo = ref(1);
const pageSize = ref(10);
const totalNum = ref(0);
const tableData = ref([]);
const layout = "prev, pager, next";
const tableColumn = ref([{
    prop: "appName",
    label: "应用名",
}]);
const showDialog = ref(0);
const processorData = ref();

// 获取处理器字典
function getProcessorDictionary() {
    api.getProcessorDictionary().then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "处理器字典获取失败：" + msg });
            return;
        }
        data.forEach((item: any) => {
            tableColumn.value.push({
                prop: item.value,
                label: item.label,
            });
        });
    });
}

// 分页获取应用插件信息
function getProcessorByPage() {
    api.getProcessorByPage({
        ...condition,
        pageNo: pageNo.value,
        pageSize: pageSize.value
    }).then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用插件列表获取失败：" + msg });
            return;
        }
        totalNum.value = data.total;
        tableData.value = data.data;
    });
}

// 处理请求参数，再获取通知数据
function search(inputValue: any) {
    condition.appName = inputValue.appName;
    getProcessorByPage();
}

// 页码变化
function handlePageChange(page: number) {
    pageNo.value = page;
    getProcessorByPage();
}

// 打开编辑弹窗
function openEditDialog(data: any) {
    processorData.value = data.row;
    showDialog.value++;
}

// 刷新处理器容器
function refreshProcessor(data: any) {
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足" });
        return;
    }
    api.refreshProcessor([data.row.appId]).then(({code, msg}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "处理器容器刷新失败：" + msg });
            return;
        }
        ElMessage({ type: "success", message: "处理器容器刷新成功" });
    });
}

onMounted(() => {
    getProcessorByPage();
    getProcessorDictionary();
});
</script>

<style lang="scss" scoped>
.manage {
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
}
</style>