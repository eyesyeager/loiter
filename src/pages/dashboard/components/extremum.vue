<template>
    <div class="extremum">
        <ButtonCard class="buttonCard" 
            v-for="item in extremumData" 
            :value="item.value" :describe="item.describe"
            :key="item.describe + refreshKey"
        />
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import ButtonCard from "@/components/card/buttonCard.vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";

interface ExtremumDataInterface {
    key: string,
    value: number,
    describe: string
};

const props = defineProps({
    condition: Object,
});
const refreshKey = ref(0);
const extremumData = reactive<ExtremumDataInterface[]>([
    {
        key: "requestNum",
        value: 0,
        describe: "总请求数"
    },
    {
        key: "runTimeMin",
        value: 0,
        describe: "最小响应时间(ms)"
    },
    {
        key: "runTimeMax",
        value: 0,
        describe: "最大响应时间(ms)"
    },
    {
        key: "runTimeAvg",
        value: 0,
        describe: "平均响应时间(ms)"
    },
    {
        key: "qpsAvg",
        value: 0,
        describe: "平均QPS"
    },
    {
        key: "requestReject",
        value: 0,
        describe: "请求拒绝数"
    },
]);

// 获取极值数据
function getDetailedRequestExtremumLog() {
    api.getDetailedRequestExtremumLog(props.condition).then(({code, msg, data}) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "极值信息获取失败：" + msg });
            return;
        }
        for (let index in extremumData) {
            extremumData[index].value = data[extremumData[index].key];
        }
        refreshKey.value++;
    })
}

onMounted(() => {
    getDetailedRequestExtremumLog();
});

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";

.extremum {
    margin: 15px auto;
    display: flex;
    justify-content: space-between;

    .buttonCard {
        width: 15%;
        height: 80px;
        background: $baseBackground;
    }
}
</style>