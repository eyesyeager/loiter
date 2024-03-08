<template>
    <div class="reject">
        <Polyline class="polyline" :option="option" :key="refreshKey" />
    </div>
</template>

<script setup lang="ts">
import Polyline from "@/components/echarts/polyline.vue";
import { ref, reactive, onMounted } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";

const props = defineProps({
    condition: Object
});
const refreshKey = ref(0);
const option = reactive({
    title: {
        show: true,
        text: "请求拒绝数",
        x: "center",
        y: "top",
    },
    xAxis: {
        type: "category",
        data: []
    },
    yAxis: {
        name: "次",
        type: "value"
    },
    legend: {
        data: ["拦截数", "错误数"],
        y: "bottom",
    },
    tooltip: {
        trigger: "axis"
    },
    series: [
        {
            name: "拦截数",
            type: "line",
            data: []
        },
        {
            name: "错误数",
            type: "line",
            data: []
        }]
});

// 获取请求拒绝统计信息
function getDetailedRequestRejectLog() {
    api.getDetailedRequestRejectLog(props.condition).then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "请求拒绝统计信息获取失败：" + msg });
            return;
        }
        // 数据赋值
        option.xAxis.data = data.xAxis;
        option.series[0].data = data.rejectSeries;
        option.series[1].data = data.errorSeries;
        refreshKey.value++;
    });
}

onMounted(() => {
    getDetailedRequestRejectLog();
});

</script>

<style lang="scss" scoped>
.reject {
    height: 300px;
    padding: 15px 0;

    .polyline {
        height: 100%;
    }
}
</style>