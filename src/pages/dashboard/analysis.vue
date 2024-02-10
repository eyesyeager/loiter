<template>
    <div class="analysis">
        <div class="condition">
            <span class="label">应用名</span>
            <el-select class="appName" v-model="appName" clearable placeholder="请选择">
                <el-option v-for="item in appNameOptions" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
            <span class="label">时间范围</span>
            <el-date-picker class="dateRange" v-model="dateRange" clearable type="daterange" range-separator="至" start-placeholder="开始时间" end-placeholder="结束时间" />
            <el-button class="button search" plain>查询</el-button>
        </div>
        <Extremum class="extremum" />
        <div class="echartsGroup">
            
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import Extremum from "./components/extremum.vue";

const appName = ref("");
const appNameOptions = ref([
    {
        "label": "eyesSpace",
        "value": "eyesSpace"
    },
    {
        "label": "eyesStars",
        "value": "eyesStars"
    },
    {
        "label": "eyesUser",
        "value": "eyesUser"
    },
    {
        "label": "eyesHunt",
        "value": "eyesHunt"
    }
]);
const dateRange = ref("")

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";
@import "@/assets/css/size.scss";

.analysis {
    background: $baseBackground;
    width: calc(100% - 30px - 40px); // 减去 margin 和 padding
    margin: 15px auto;
    padding: 10px 20px;
    border-radius: $cardBorder;

    .condition {
        height: 50px;
        line-height: 50px;

        .label {
            margin-right: 10px;
        }

        .appName {
            width: 150px;
            margin-right: 10px;
        }

        .button {
            margin-left: 10px;
        }
    }
}
</style>