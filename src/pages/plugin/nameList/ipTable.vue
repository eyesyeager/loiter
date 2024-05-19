<template>
    <div class="ipTable">
        <el-table class="table" :border="true" :data="tableData">
            <el-table-column type="index" :align="'center'" :label="props.genre!.label.charAt(0)" />
            <el-table-column v-for="item in tableColumn" :prop="item.prop" :label="item.label" :align="'center'" />
            <el-table-column prop="content" label="操作" :align="'center'">
                <template #default="scope">
                    <el-popconfirm title="操作将立即生效，确认删除吗?" @confirm="deleteNameListIp(scope)">
                        <template #reference><el-button size="small">删除</el-button></template>
                    </el-popconfirm>
                </template>
            </el-table-column>
        </el-table>
        <el-pagination class="pagination" :layout="layout" :page-sizes="pageSizes" :current-page="pageNo"
                :page-size="pageSize" :total="totalNum" @current-change="handlePageChange" @size-change="handleSizeChange" 
            />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode, role } from "@/config";
import { useRoleStore } from "@/store";

const props = defineProps({
    genre: Object,
    condition: Object,
    refresh: Number
});
const roleStore = useRoleStore();
const tableData = ref([]);
const tableColumn = ref([
    { prop: "appName", label: "应用名" },
    { prop: "ip", label: "IP地址" },
    { prop: "remarks", label: "备注" },
    { prop: "createdAt", label: "创建时间" },
]);
const pageNo = ref(1);
const pageSize = ref(10);
const totalNum = ref(0);
const layout = "total, sizes, prev, pager, next";
const pageSizes = [10, 50, 100, 200];

// 监听父组件传值变化，控制表格刷新
watch(
    () => props.refresh,
    _ => {
        pageNo.value = 1;
        getNameList();
    }
);

// 分页获取 ip 信息
function getNameList() {
    api.getNameList({
        ...props.condition,
        genre: props.genre!.value,
        pageNo: pageNo.value,
        pageSize: pageSize.value
    }).then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "IP 名单获取失败：" + msg });
            return;
        }
        tableData.value = data.data;
        totalNum.value = data.total;
    });
}

// 单页大小变化
function handleSizeChange(size: number) {
    pageSize.value = size;
    getNameList();
}

// 页码变化
function handlePageChange(page: number) {
    pageNo.value = page;
    getNameList();
}

// 删除名单ip
function deleteNameListIp(data: any) {
    // 权限校验
        if (!roleStore.checkAuth(role.admin)) {
        ElMessage({ type: "error", message: "权限不足！" });
        return;
    }
    api.deleteNameListIp({
        appId: data.row.appId,
        id: data.row.id,
        ip: data.row.ip,
        genre: props.genre!.value
    }).then(({code, msg}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "IP 删除失败：" + msg });
            return;
        }
        getNameList();
    });
}

onMounted(() => {
    getNameList();
});

</script>

<style lang="scss" scoped>
.ipTable {
    position: relative;
    .pagination {
        position: absolute;
        right: 0px;
        margin: 10px 0;
    }
}
</style>