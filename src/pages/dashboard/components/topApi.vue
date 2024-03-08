<template>
    <div class="topApi">
        <div class="title">Top10 API</div>
        <div class="table" >
            <div class="header">
                <div class="api">接口地址</div>
                <div class="num">调用次数</div>
                <div class="rate">占比</div>
            </div>
            <div class="data" v-for="item in apiData">
                <div class="api">{{ item.api }}</div>
                <div class="num">{{ item.num }}</div>
                <div class="rate">{{ item.rate }}</div>
            </div>
        </div>
        <div class="empty" v-show="show">暂无数据</div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";

interface TopApiInterface {
    api: string,
    num: number,
    rate: string
}

const props = defineProps({
    condition: Object
});
const show = ref(true);
const apiData = ref<TopApiInterface[]>([]);

// 获取请求统计信息
function getDetailedRequestTopApiLog() {
    api.getDetailedRequestTopApiLog(props.condition).then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "请求统计信息获取失败：" + msg });
            return;
        }
        if (data != null && data.length != 0) {
            apiData.value = data;
            show.value = false;
        }
    });
}

onMounted(() => {
    getDetailedRequestTopApiLog();
});

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";

.topApi {
    height: 300px;
    padding: 15px 0;

    .title {
        font-size: 18px;
        font-weight: bold;
        color: $echartsTitleText;
        text-align: center;
        height: 25px;
        line-height: 25px;
    }

    .header,
    .data {
        display: flex;
        color: $regularText;
        padding-left: 15px;
        height: 25px;
        line-height: 27.5px;
    }

    .header {
        font-size: 14px;
    }

    .data {
        font-size: 12px;
    }

    .api {
        flex: 1;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }

    .num,
    .rate {
        width: 100px;
        text-align: center;
    }

    .empty {
        text-align: center;
        color: $regularText;
        height: 272px;
        line-height: 272px;
    }
}
</style>