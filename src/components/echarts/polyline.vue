<template>
  <div class="polyline" ref="polyline"></div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import * as echarts from "echarts";

const polyline = ref();
const props = defineProps({
  option: Object
});

onMounted(() => {
  const currEcharts = echarts.init(polyline.value);
  currEcharts.setOption(props.option!);
});

</script>

<style lang="scss" scoped></style>