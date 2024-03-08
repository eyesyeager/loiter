<template>
    <div class="dashboard">
        <Statistic />
        <Analysis />
        <Footer />
    </div>
</template>

<script setup lang="ts">
import Statistic from "./statistic.vue";
import Analysis from "./analysis.vue";
import Footer from "@/components/common/footer.vue";
</script>