<template>
    <div class="application">
        <Condition @search="search" />
        <div class="appGroup">
            <app-card class="appCard" v-for="(item, index) in appData" :key="index" :data="item" @reload="getAppInfoByPage" @edit="showUpdateDialog"/>
        </div>
        <el-empty description="暂无应用配置" v-show="isEmpty"/>
        <el-pagination class="pagination" :layout="layout" 
            :current-page="pageNo" :page-size="pageSize" :total="totalNum" @current-change="handlePageChange" />
        <el-button class="saveBtn" :icon="Plus" type="primary" circle size="large" @click="showAddDialog" />
        <Save :show="showDialog" :flag="flagDialog" :appId="appId" @reload="getAppInfoByPage" />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import Condition from "./condition.vue";
import AppCard from "./appCard.vue";
import Save from "./save.vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";
import { Plus } from "@element-plus/icons-vue";
import { SaveAppDialog } from "@/constants";

let condition = {
    appName: "",
    status: "",
};
const pageNo = ref(1);
const pageSize = ref(9);
const appData = ref([]);
const layout = "prev, pager, next";
const totalNum = ref(0);
const isEmpty = ref(true);
const showDialog = ref(0);
const flagDialog = ref();
const appId = ref(0);

// 执行条件搜索
function search(inputValue: any) {
    condition.appName = inputValue.appName;
    condition.status = inputValue.status;
    getAppInfoByPage();
}

// 分页获取应用详细信息
function getAppInfoByPage() {
    api.getAppInfoByPage({
        ...condition,
        pageNo: pageNo.value,
        pageSize: pageSize.value
    }).then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用详细信息获取失败：" + msg });
            return;
        }
        isEmpty.value = data.data == null || data.data.length == 0;
        appData.value = data.data;
        totalNum.value = data.total;
    });
}

// 页码变化
function handlePageChange(page: number) {
    pageNo.value = page;
    getAppInfoByPage();
}

// 打开新增弹窗
function showAddDialog() {
    flagDialog.value = SaveAppDialog.add;
    appId.value = 0;
    showDialog.value++;
}

// 打开编辑弹窗
function showUpdateDialog(id: number) {
    flagDialog.value = SaveAppDialog.update;
    appId.value = id;
    showDialog.value++;
}

onMounted(() => {
    getAppInfoByPage();
});

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";
@import "@/assets/css/size.scss";

.application {
    .appGroup {
        width: calc(100% - 30px);
        margin: 0px auto;
        display: flex;
        flex-wrap: wrap;

        .appCard {
            width: calc((100% - 30px) / 3);
            margin-right: 15px;
            margin-bottom: 15px;
            &:nth-child(3n){
                margin-right: 0;
            }
        }
    }

    .pagination {
        position: absolute;
        right: 15px;
    }

    .saveBtn {
        position: fixed;
        right: 25px;
        bottom: 25px;
    }
}

@media screen and (max-width: 1400px) {
    .application {
        .appGroup {
            .appCard {
                width: calc((100% - 15px) / 2) !important;
                &:nth-child(3n){
                    margin-right: 15px;
                }
                &:nth-child(even){
                    margin-right: 0px;
                }
            }
        }
    }
}
</style>