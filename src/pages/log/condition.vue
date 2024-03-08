<template>
    <div class="condition">
        <div class="line">
            <div class="inputGroup">
                <span class="label">操作人:</span>
                <el-input class="input" v-model="inputValue.operatorName" clearable />
            </div>
            <div class="inputGroup">
                <span class="label">标题:</span>
                <el-input class="input" v-model="inputValue.title" clearable />
            </div>
            <div class="inputGroup">
                <span class="label">内容:</span>
                <el-input class="input" v-model="inputValue.content" clearable />
            </div>
        </div>
        <div class="line">
            <div>
                <span class="label">时间:</span>
                <el-date-picker class="timeRange" v-model="timeRange" type="daterange" range-separator="至"
                    start-placeholder="开始" end-placeholder="结束" />
            </div>
        </div>
        <div class="buttonGroup">
            <el-button class="search" type="primary" @click="search">查询</el-button>
            <el-button class="reset" @click="reset">重置</el-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";

const emit = defineEmits([ "search" ]);
const inputValue = reactive({
    operatorName: "",
    title: "",
    content: "",
});
const timeRange = ref<Date[]>([]);

// 重置表单
function reset() {
    inputValue.operatorName = "";
    inputValue.title = "";
    inputValue.content = "";
    timeRange.value = [];
}

// 通知父组件执行查询
function search() {
    emit("search", inputValue, timeRange.value);
}

</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";
@import "@/assets/css/size.scss";

.condition {
    width: calc(100% - 30px - 30px);
    margin: 15px auto;
    padding: 15px;
    background: $baseBackground;
    border-radius: $cardBorder;

    .line {
        display: flex;
        height: 50px;
        line-height: 50px;

        .inputGroup {
            margin-right: 10px;
        }

        .label {
            width: 60px;
            margin-right: 5px;
        }

        .input {
            width: 250px;
        }
    }

    .buttonGroup {
        text-align: right;
    }
}

@media screen and (max-width: 1250px) {
    .input {
        width: 180px !important;
    }
}
</style>