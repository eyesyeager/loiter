<template>
    <div class="qps">
        <Polyline class="polyline" :option="option" :key="refreshKey" />
    </div>
</template>

<script setup lang="ts">
import Polyline from "@/components/echarts/polyline.vue";
import { ref, reactive, onMounted } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";
import { getDayFormat } from "@/utils/utils";

const props = defineProps({
    condition: Object
});
const refreshKey = ref(0);
const option = reactive<any>({
    title: {
        show: true,
        text: "QPS",
        x: "center",
        y: "top",
    },
    xAxis: {
        type: "category",
        data: []
    },
    yAxis: {
        name: "次/s",
        type: "value"
    },
    tooltip: {
        trigger: "axis"
    },
    legend: {
        data: [getDayFormat(-1), getDayFormat()],
        y: "bottom",
    },
    series: [
        {
            name: getDayFormat(-1),
            type: "line",
            data: []
        },
        {
            name: getDayFormat(),
            type: "line",
            data: []
        },
    ]
});

// 获取QPS统计信息
function getDetailedRequestQPSLog() {
    api.getDetailedRequestQPSLog({
        ...props.condition,
        appId: props.condition!.appId ? Number(props.condition!.appId) : null
    }).then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "QPS统计信息获取失败：" + msg });
            return;
        }
        // 数据赋值
        option.xAxis.data = data.xAxis;
        if (props.condition!.timeInterval == "today" || props.condition!.timeInterval == "yesterday") {
            let dayFront = getDayFormat(props.condition!.timeInterval == "today" ? 0 : -1);
            let datLast = getDayFormat(props.condition!.timeInterval == "today" ? -1 : -2);
            option.legend.data = [datLast, dayFront];
            option.series = [
                {
                    name: datLast,
                    type: "line",
                    data: data.lastSeries
                },
                {
                    name: dayFront,
                    type: "line",
                    data: data.series
                },
            ];
        } else {
            option.legend.data = [];
            option.series = [
                {
                    type: "line",
                    data: data.series
                }
            ];
        }
        refreshKey.value++;
    });
}

onMounted(() => {
    getDetailedRequestQPSLog();
});

</script>

<style lang="scss" scoped>
.qps {
    height: 300px;
    padding: 15px 0;
    
    .polyline {
        height: 100%;
    }
}
</style>