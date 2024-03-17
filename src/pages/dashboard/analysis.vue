<template>
    <div class="analysis">
        <div class="condition">
            <span class="label">应用名</span>
            <el-select class="appName" v-model="condition.appId" filterable clearable>
                <el-option v-for="item in appOptions" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
            <span class="label">时间范围</span>
            <el-radio-group class="timeInterval" v-model="condition.timeInterval">
                <el-radio-button v-for="item in timeIntervalOptions" :label="item.label">{{ item.name }}</el-radio-button>
            </el-radio-group>
            <el-button class="search" plain @click="search">查询</el-button>
        </div>
        <Extremum class="extremum" :condition="condition" :key="refreshKey" />
        <div class="echartsGroup">
            <request-echarts class="chart" :condition="condition" :key="refreshKey"/>
            <visitor-echarts class="chart" :condition="condition" :key="refreshKey"/>
            <qps-echarts class="chart" :condition="condition" :key="refreshKey" />
            <runtime-echarts class="chart" :condition="condition" :key="refreshKey"/>
            <top-api-echarts class="chart" :condition="condition" :key="refreshKey"/>
            <reject-echarts class="chart" :condition="condition" :key="refreshKey"/>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";
import Extremum from "./components/extremum.vue";
import RequestEcharts from "./components/request.vue";
import RuntimeEcharts from "./components/runtime.vue";
import QpsEcharts from "./components/qps.vue";
import VisitorEcharts from "./components/visitor.vue";
import TopApiEcharts from "./components/topApi.vue";
import RejectEcharts from "./components/reject.vue";
import { OptionsInterface } from "@/d.ts/common";

const appOptions = reactive<OptionsInterface[]>([]);
const refreshKey = ref(0);
const timeIntervalOptions = [
    { label: "today", name: "今天" },
    { label: "yesterday", name: "昨天" },
    { label: "week", name: "最近7天" },
    { label: "month", name: "最近30天" },
];
const condition = reactive({
    appId: "",
    timeInterval: "today"
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

// 获取应用详细请求信息
function search() {
    refreshKey.value += 1;
}

onMounted(() => {
    getAppDictionary();
});

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";
@import "@/assets/css/size.scss";

.analysis {
    width: calc(100% - 30px);
    margin: 0 auto;

    .condition {
        height: 35px;
        line-height: 35px;
        padding: 15px;
        background: $baseBackground;
        border-radius: $cardBorder;
        
        .label {
            margin-right: 10px;
        }

        .appName {
            width: 200px;
            margin-right: 10px;
        }

        .timeInterval {
            transform: translateY(-4px);
        }

        .search {
            margin-left: 10px;
        }
    }

    .echartsGroup {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;

        .chart {
            width: calc(49%);
            margin-bottom: 15px;
            background: $baseBackground;
            border-radius: $cardBorder;
        }
    }
}
</style>