<template>
    <div class="notice">
        <Condition @search="search" />
        <div class="data">
            <el-table class="table" :border="true" :data="tableData">
                <el-table-column type="index" :align="'center'" />
                <el-table-column prop="appName" label="应用名" :width=150 :align="'center'" />
                <el-table-column prop="genre" label="类型" :width=100 :align="'center'" />
                <el-table-column prop="title" label="标题" :width=230 :align="'center'" />
                <el-table-column prop="content" label="内容" :align="'center'">
                    <template #default="scope">
                        <span v-if="scope.row.genre == 'site'">{{ scope.row.content }}</span>
                        <el-button size="small" v-else @click="viewEmail(scope.row.id)">查看</el-button>
                    </template>
                </el-table-column>
                <el-table-column prop="remarks" label="备注" :align="'center'" />
                <el-table-column prop="createdAt" label="通知时间" :width=200 :align="'center'" />
            </el-table>
            <el-pagination class="pagination" :layout="layout" :current-page="pageNo" :page-size="pageSize"
                :total="totalNum" @current-change="handlePageChange" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import Condition from "./condition.vue";
import api from "@/apis/api";
import { ElMessage, ElMessageBox } from "element-plus";
import { responseCode } from "@/config";
import moment from "moment";

let condition = {
    appName: "",
    genre: "",
    title: "",
    timeBegin: "",
    timeEnd: "",
}
const pageNo = ref(1);
const pageSize = ref(10);
const totalNum = ref(0);
const tableData = ref([]);
const layout = "prev, pager, next";

// 处理请求参数，再获取通知数据
function search(inputValue: any, timeRange: any) {
    condition.appName = inputValue.appName;
    condition.genre = inputValue.genre;
    condition.title = inputValue.title;
    condition.timeBegin = timeRange[0] ? moment(timeRange.value[0]).format("YYYY-MM-DD") : "";
    condition.timeEnd = timeRange[1] ? moment(timeRange.value[1]).format("YYYY-MM-DD") : "";
    getNoticeList();
}

// 获取通知数据
function getNoticeList() {
    api.getNoticeList({
        ...condition,
        pageNo: pageNo.value,
        pageSize: pageSize.value
    }).then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "消息通知获取失败：" + msg });
            return;
        }
        totalNum.value = data.total;
        tableData.value = data.data;
    });
}

// 页码更改
function handlePageChange(page: number) {
    pageNo.value = page;
    getNoticeList();
}

// 查看email详情
function viewEmail(id: number) {
    api.getEmailNoticeContent([id]).then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "邮件信息获取失败：" + msg });
            return;
        }
        ElMessageBox.alert(data, "", { dangerouslyUseHTMLString: true });
    });
}

onMounted(() => {
    getNoticeList();
});

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";
@import "@/assets/css/size.scss";

.notice {
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