<template>
    <div class="runtime">
        <Polyline class="polyline" :option="option" :key="refreshKey" v-if="show" />
        <div class="empty" v-else>
            <div class="title">{{ option.title.text }}</div>
            <div class="content">暂无数据</div>
        </div>
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
const show = ref(false);
const option = reactive({
    title: {
        show: true,
        text: "响应时间(RT)",
        x: "center",
        y: "top",
    },
    xAxis: {
        type: "category",
        data: []
    },
    yAxis: {
        name: "ms",
        type: "value"
    },
    tooltip: {
        trigger: "axis"
    },
    series: [{
        type: "line",
        data: []
    }]
});

// 获取响应时间统计信息
function getDetailedRequestRuntimeLog() {
    api.getDetailedRequestRuntimeLog(props.condition).then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "响应时间信息获取失败：" + msg });
            return;
        }
        // 数据赋值
        let emptyFlag = true;
        for (let index in data.series) {
            if (data.series[index] == -1) {
                data.series[index] = "-";
            } else {
                emptyFlag = false;
            }
        }
        if (emptyFlag) {
            return;
        }
        option.xAxis.data = data.xAxis;
        option.series = [
            {
                type: "line",
                data: data.series
            }
        ];
        refreshKey.value++;
        show.value = true;
    });
}

onMounted(() => {
    getDetailedRequestRuntimeLog();
});

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";

.runtime {
    height: 300px;
    padding: 15px 0;

    .polyline {
        width: 100%;
        height: 300px;
    }

    .empty {
        .title {
            font-size: 18px;
            font-weight: bold;
            color: $echartsTitleText;
            text-align: center;
            height: 28px;
            line-height: 28px;
        }

        .content {
            text-align: center;
            color: $regularText;
            height: 272px;
            line-height: 272px;
        }
    }
}
</style>