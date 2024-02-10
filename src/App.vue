<template>
    <router-view v-slot="{ Component }">
        <template v-if="Component">
            <component :is="Component" />
        </template>
    </router-view>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { publicPath } from '@/router/path';

const router = useRouter();
router.isReady().then(async () => {
    router.replace(publicPath.common.login);
})
</script>