<template>
    <div class="nameList">
        <Condition @search="search"/>
        <div class="ipTable">
            <Table />
            <Table />
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
import { ref, onMounted } from "vue";
import { Plus, Setting } from "@element-plus/icons-vue";
import Condition from "./condition.vue";
import Table from "./table.vue";
import Add from "./add.vue";
import Switch from "./switch.vue";
import moment from "moment";

const showAddDialog = ref(0);
const showSwitchDialog = ref(0);
let condition = {
    appId: "",
    ip: "",
    remarks: "",
    timeBegin: "",
    timeEnd: ""
}

// 处理请求参数，再获取通知数据
function search(inputValue: any, timeRange: any) {
    condition.appId = inputValue.appId;
    condition.ip = inputValue.ip;
    condition.remarks = inputValue.remarks;
    condition.timeBegin = timeRange[0] ? moment(timeRange[0]).format("YYYY-MM-DD") : "";
    condition.timeEnd = timeRange[1] ? moment(timeRange[1]).format("YYYY-MM-DD") : "";
}

</script>

<style lang="scss" scoped>
.nameList {
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