<template>
    <div class="condition">
        <span class="label">应用名</span>
        <el-select class="appName" v-model="inputValue.appName" filterable clearable placeholder="请选择">
            <el-option v-for="item in appNameOptions" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
        <span class="label">状态</span>
        <el-select class="status" v-model="inputValue.status" clearable placeholder="请选择">
            <el-option v-for="item in statusOptions" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
        <el-button class="search" plain @click="search">查询</el-button>
    </div>
</template>

<script setup lang="ts">
import { onMounted, reactive } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";
import { OptionsInterface } from "@/d.ts/common";

const emit = defineEmits([ "search" ]);
const appNameOptions = reactive<OptionsInterface[]>([]);
const statusOptions = reactive<OptionsInterface[]>([]);
const inputValue = reactive({
    appName: "",
    status: ""
});

// 获取所有应用信息
function getAllApp() {
    api.getAllApp().then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用信息获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            appNameOptions[index] = {
                "label": item,
                "value": item
            };
        });
    });
}

// 获取状态字典
function getStatusDictionary() {
    api.getStatusDictionary().then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "状态字典获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            statusOptions[index] = item;
        });
    });
}

// 通知父组件进行查询
function search() {
    emit("search", inputValue);
}

onMounted(() => {
    getAllApp();
    getStatusDictionary();
});

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";
@import "@/assets/css/size.scss";

.condition {
    width: calc(100% - 60px);
    margin: 15px auto;
    height: 35px;
    line-height: 35px;
    padding: 15px;
    background: $baseBackground;
    border-radius: $cardBorder;

    .label {
        margin-right: 10px;
    }

    .appName,
    .status {
        width: 200px;
        margin-right: 10px;
    }

    .search {
        margin-left: 10px;
    }
}
</style>