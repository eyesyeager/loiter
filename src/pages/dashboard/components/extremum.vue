<template>
    <div class="extremum">
        <ButtonCard class="buttonCard" 
            v-for="item in extremumData" 
            :value="item.value" :describe="item.describe"
        />
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import ButtonCard from "@/components/card/buttonCard.vue";

const extremumData = ref([
    {
        value: "57829",
        describe: "总请求数"
    },
    {
        value: "5ms",
        describe: "最小响应时间"
    },
    {
        value: "1128ms",
        describe: "最大响应时间"
    },
    {
        value: "57ms",
        describe: "平均响应时间"
    },
    {
        value: "5",
        describe: "平均QPS"
    },
    {
        value: "2",
        describe: "请求拒绝数"
    },
]);

</script>

<style lang="scss" scoped>
.extremum {
    margin-top: 20px;
    margin-bottom: 10px;
    display: flex;
    justify-content: space-around;
    .buttonCard {
        margin: 0 5px;
    }
}
</style>