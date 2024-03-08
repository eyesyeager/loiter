<template>
    <div class="footer">
        <div>{{ `Copyright © ${site.establishment}-至今. All Rights Reserved.` }}</div>
        <div>{{ site.name + " - " + site.author }}</div>
    </div>
</template>
  
<script lang="ts" setup>
import { site } from "@/config";
</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";

.footer {
    text-align: center;
    color: $secondaryText;
    font-size: 12px;
    height: 60px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    div {
        height: 20px;
        line-height: 20px;
    }
}
</style>