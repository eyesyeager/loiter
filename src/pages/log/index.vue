<template>
    <div class="log">
        <Condition @search="search" />
        <div class="data">
            <el-table class="table" :border="true" :data="tableData">
                <el-table-column type="index" :align="'center'"/>
                <el-table-column v-for="item in tableColumn" :prop="item.prop" :label="item.label" :width="item.width" :align="item.align"/>
            </el-table>
            <el-pagination class="pagination" :layout="layout" :page-sizes="pageSizes" :current-page="pageNo"
                :page-size="pageSize" :total="totalNum" @current-change="handlePageChange" @size-change="handleSizeChange" 
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";
import moment from "moment";
import Condition from "./condition.vue";

let condition = {
    operatorName: "",
    title: "",
    content: "",
    timeBegin: "",
    timeEnd: ""
};
const pageNo = ref(1);
const pageSize = ref(10);
const totalNum = ref(0);
const tableData = ref([]);
const layout = "total, sizes, prev, pager, next";
const pageSizes = [10, 50, 100, 200];
const tableColumn = [
    {
        prop: "title",
        label: "操作类型",
        width: 230,
        align: "left"
    },
    {
        prop: "content",
        label: "操作内容",
        align: "left"
    },
    {
        prop: "operator",
        label: "操作人",
        width: 100,
        align: "center"
    },
    {
        prop: "ip",
        label: "操作人IP",
        width: 100,
        align: "center"
    },
    {
        prop: "createdAt",
        label: "操作时间",
        width: 200,
        align: "center"
    },
];

// 处理请求参数，再获取通知数据
function search(inputValue: any, timeRange: any) {
    condition.operatorName = inputValue.operatorName;
    condition.title = inputValue.title;
    condition.content = inputValue.content;
    condition.timeBegin = timeRange[0] ? moment(timeRange.value[0]).format("YYYY-MM-DD") : "";
    condition.timeEnd = timeRange[1] ? moment(timeRange.value[1]).format("YYYY-MM-DD") : "";
    getUniversalLog();
}

// 发送日志查询请求
function getUniversalLog() {
    api.getUniversalLog({
        ...condition,
        pageNo: pageNo.value,
        pageSize: pageSize.value
    }).then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "操作日志获取失败：" + msg });
            return;
        }
        tableData.value = data.data;
        totalNum.value = data.total;
    });
}

// 单页大小变化
function handleSizeChange(size: number) {
    pageSize.value = size;
    getUniversalLog();
}

// 页码变化
function handlePageChange(page: number) {
    pageNo.value = page;
    getUniversalLog();
}

onMounted(() => {
    getUniversalLog();
});

</script>

<style lang="scss" scoped>
.log {
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