<template>
    <div class="statistic">
        <el-statistic class="item" v-for="item in statisticData" :title="item.title" :value="item.value" />
    </div>
</template>

<script setup lang="ts">
import { onMounted, reactive } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";

interface statisticDataInterface {
    key: string,
    title: string
    value: number,
}

let statisticData = reactive<statisticDataInterface[]>([
    {
        key: "requestNum",
        title: "今日接口调用量",
        value: 0
    },
    {
        key: "avgRunTime",
        title: "今日接口平均耗时(ms)",
        value: 0
    },
    {
        key: "visitorNum",
        title: "今日访客总数",
        value: 0
    },
    {
        key: "rejectNum",
        title: "今日请求拒绝总数",
        value: 0
    },
])

// 获取请求信息概览
function getOverviewRequestLog() {
    api.getOverviewRequestLog().then(({code, msg, data}) => {
        if (code == responseCode.success) {
            for (let index in statisticData) {
                statisticData[index].value = data[statisticData[index].key];
            }
        } else {
            ElMessage({ type: "error", message: "请求信息概览获取失败：" + msg });
        }
    })
}

onMounted(() => {
    getOverviewRequestLog();
})
</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";
@import "@/assets/css/size.scss";

.statistic {
    background: $baseBackground;
    width: calc(100% - 30px);
    height: 140px;
    margin: 15px auto;
    border-radius: $cardBorder;
    display: flex;
    justify-content: space-around;
    align-items: center;
    .item {
        text-align: center;
        :deep(.el-statistic__head) {
            font-size: 18px;
            margin-bottom: 10px;
        }
    }
}
</style>