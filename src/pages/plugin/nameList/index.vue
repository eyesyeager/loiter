<template>
    <div class="nameList">
        <Condition @search="search"/>
        <div class="ipTableGroup">
            <ip-table class="table" v-for="item in nameListOptions" :genre="item" :condition="condition" :refresh="refreshKey"/>
        </div>
        <Add :show="showAddDialog" />
        <Switch :show="showSwitchDialog" />
        <div class="btnGroup">
            <el-button class="btn" :icon="Plus" type="primary" circle size="large" @click="showAddDialog++" />
            <div class="segment" />
            <el-button class="btn" :icon="Setting" type="primary" circle size="large" @click="showSwitchDialog++" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive } from "vue";
import { Plus, Setting } from "@element-plus/icons-vue";
import Condition from "./condition.vue";
import IpTable from "./ipTable.vue";
import Add from "./add.vue";
import Switch from "./switch.vue";
import moment from "moment";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";
import { OptionsInterface } from "@/d.ts/common";

const showAddDialog = ref(0);
const showSwitchDialog = ref(0);
const nameListOptions = reactive<OptionsInterface[]>([]);
const refreshKey = ref(0);
let condition = {
    appId: 0,
    ip: "",
    remarks: "",
    timeBegin: "",
    timeEnd: ""
}

// 处理请求参数，再获取通知数据
function search(inputValue: any, timeRange: any) {
    condition.appId = inputValue.appId ? Number(inputValue.appId) : 0;
    condition.ip = inputValue.ip;
    condition.remarks = inputValue.remarks;
    condition.timeBegin = timeRange[0] ? moment(timeRange[0]).format("YYYY-MM-DD") : "";
    condition.timeEnd = timeRange[1] ? moment(timeRange[1]).format("YYYY-MM-DD") : "";
    refreshKey.value++;
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

onMounted(() => {
    getNameListDictionary();
})

</script>

<style lang="scss" scoped>
.nameList {
    .ipTableGroup {
        width: calc(100% - 30px);
        margin: 15px auto;
        display: flex;
        .table {
            width: 50%;
        }
    }

    .btnGroup {
        position: fixed;
        right: 25px;
        bottom: 25px;

        .segment {
            height: 10px;
        }
    }
}
</style>