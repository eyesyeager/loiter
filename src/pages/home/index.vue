<template>
    <div class="home">
        <Header class="header"/>
        <div class="flex main-container">
            <Sider class="sider"/>
            <router-view v-slot="{ Component }">
                <template v-if="Component">
                    <component :is="Component" class="content"/>
                </template>
            </router-view>
        </div>
    </div>
</template>

<script setup lang="ts">
import Header from "./header.vue";
import Sider from "./sider.vue";
</script>

<style lang="scss" scoped>
.home {
    min-width: 1000px;
    .header {
        position: fixed;
        width: 100%;
        min-width: 1000px;
        height: 59px;
    }
    .sider {
        position: fixed;
        top: 60px;
        width: 140px;
        height: 100%;
    }
    .content {
        position: fixed;
        top: 60px;
        left: 140px;
        width: calc(100% - 140px);
        min-width: 860px;
        height: calc(100% - 60px);
        overflow-y: scroll;
    }
}
</style>