<template>
    <div class="statistic">
        <el-statistic class="item" title="今日接口调用量" :value="268500" />
        <el-statistic class="item" title="今日接口总数" :value="252" />
        <el-statistic class="item" title="今日访问总数" :value="859" />
        <el-statistic class="item" title="今日请求拒绝总数" :value="120" />
    </div>
</template>

<script setup lang="ts">
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