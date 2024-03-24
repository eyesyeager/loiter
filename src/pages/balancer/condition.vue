<template>
    <div class="condition">
        <span class="label">应用</span>
        <el-select class="appName" v-model="inputValue.appId" filterable clearable>
            <el-option v-for="item in appOptions" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
        <span class="label">负载策略</span>
        <el-select class="status" v-model="inputValue.balancer" clearable>
            <el-option v-for="item in balancerOptions" :key="item.value" :label="item.label" :value="item.value" />
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
const appOptions = reactive<OptionsInterface[]>([]);
const balancerOptions = reactive<OptionsInterface[]>([]);
const inputValue = reactive({
    appId: "",
    balancer: ""
});

// 获取所有应用信息
function getAppDictionary() {
    api.getAppDictionary().then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用信息获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            appOptions[index] = {
                "label": item.label,
                "value": item.value
            };
        });
    });
}

// 获取状态字典
function getBalancerDictionary() {
    api.getBalancerDictionary().then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "负载均衡策略获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            balancerOptions[index] = item;
        });
    });
}

// 通知父组件进行查询
function search() {
    emit("search", inputValue);
}

onMounted(() => {
    getAppDictionary();
    getBalancerDictionary();
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