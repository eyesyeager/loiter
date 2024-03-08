<template>
    <div class="balancer">
        <Condition @search="search" />
        <div class="data">
            <el-table class="table" :border="true" :data="tableData">
                <el-table-column type="index" :align="'center'" />
                <el-table-column v-for="item in tableColumn" :prop="item.prop" :label="item.label" :align="'center'" />
                <el-table-column label="操作" :align="'center'">
                    <template #default="scope">
                        <el-button size="small" @click="openEditDialog(scope)">编辑</el-button>
                        <el-popconfirm title="确认刷新容器吗?" @confirm="refreshBalancer(scope)">
                            <template #reference><el-button size="small">刷新容器</el-button></template>
                        </el-popconfirm>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination class="pagination" :current-page="pageNo" :page-size="pageSize" :total="totalNum"
                :layout="layout" @current-change="handlePageChange" />
        </div>
        <Edit :data="balancerData" :show="showDialog" @reload="getBalancerByPage"/>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import Condition from "./condition.vue";
import Edit from "./edit.vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode, role } from "@/config";
import { useRoleStore } from "@/store";

let condition = {
    appName: "",
    balancer: "",
};
const roleStore = useRoleStore();
const pageNo = ref(1);
const pageSize = ref(10);
const totalNum = ref(0);
const tableData = ref([]);
const layout = "prev, pager, next";
const tableColumn = [
    {
        prop: "appName",
        label: "应用名",
    },
    {
        prop: "balancerName",
        label: "负载均衡策略",
    },
    {
        prop: "operator",
        label: "操作人",
    },
    {
        prop: "updatedAt",
        label: "操作时间",
    },
];
const showDialog = ref(0);
const balancerData = ref({});

// 处理请求参数，再获取通知数据
function search(inputValue: any) {
    condition.appName = inputValue.appName;
    condition.balancer = inputValue.balancer;
    getBalancerByPage();
}

// 分页获取应用负载均衡策略
function getBalancerByPage() {
    api.getBalancerByPage({
        ...condition,
        pageNo: pageNo.value,
        pageSize: pageSize.value,
    }).then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用负载均衡策略获取失败：" + msg });
            return;
        }
        totalNum.value = data.total;
        tableData.value = data.data;
    });
}

// 页码变化
function handlePageChange(page: number) {
    pageNo.value = page;
    getBalancerByPage();
}

// 打开编辑弹窗
function openEditDialog(data: any) {
    balancerData.value = data.row;
    showDialog.value++;
}

// 刷新负载均衡容器
function refreshBalancer(data: any) {
    if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足" });
        return;
    }
    api.refreshBalancer([data.row.id]).then(({code, msg}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用负载均衡容器刷新失败：" + msg });
            return;
        }
        ElMessage({ type: "success", message: "应用负载均衡容器刷新成功" });
    });
}

onMounted(() => {
    getBalancerByPage();
});
</script>

<style lang="scss" scoped>
.balancer {
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