<template>
    <div class="condition">
        <div class="line">
            <div class="inputGroup">
                <span class="label">应用:</span>
                <el-select class="input" v-model="inputValue.appId" filterable clearable placeholder="">
                    <el-option v-for="item in appOptions" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </div>
            <div class="inputGroup">
                <span class="label">ip:</span>
                <el-input class="input" v-model="inputValue.ip" clearable />
            </div>
            <div class="inputGroup">
                <span class="label">备注:</span>
                <el-input class="input" v-model="inputValue.remarks" clearable />
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
import { ref, onMounted, reactive } from "vue";
import api from "@/apis/api";
import { ElMessage } from "element-plus";
import { responseCode } from "@/config";
import { OptionsInterface } from "@/d.ts/common";

const emit = defineEmits(["search"]);
const appOptions = reactive<OptionsInterface[]>([]);
const timeRange = ref<Date[]>([]);
const inputValue = reactive({
    appId: "",
    ip: "",
    remarks: ""
});

// 获取所有应用信息
function getAppDictionary() {
    api.getAppDictionary().then(({ code, msg, data }) => {
        if (code != responseCode.success) {
            ElMessage({ type: "error", message: "应用信息获取失败：" + msg });
            return;
        }
        data.forEach((item: any, index: number) => {
            appOptions[index] = {
                "label": item.label,
                "value": item.value
            };
        });
    });
}

// 重置表单
function reset() {
    inputValue.appId = "";
    inputValue.ip = "";
    inputValue.remarks = "";
    timeRange.value = [];
}

// 通知父组件进行查询
function search() {
    emit("search", inputValue, timeRange.value);
}

onMounted(() => {
    getAppDictionary();
});
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
            width: 200px;
        }
    }

    .buttonGroup {
        text-align: right;
    }
}
</style>