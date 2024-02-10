<template>
    <div class="polyline" ref="polyline"></div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import * as echarts from "echarts";

const polyline = ref();

let option = {
  xAxis: {
    type: 'category',
    data: ['A', 'B', 'C']
  },
  yAxis: {
    type: 'value'
  },
  series: [
    {
      data: [120, 200, 150],
      type: 'line'
    }
  ]
};

onMounted(() => {
  const currEcharts = echarts.init(polyline.value);
  currEcharts.setOption(option);
});

</script>

<style lang="scss" scoped>
</style>