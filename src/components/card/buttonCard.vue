<template>
    <div class="buttonCard">
        <div class="value">{{ props.value }}</div>
        <div class="describe">{{ props.describe }}</div>
    </div>
</template>

<script setup lang="ts">
let props = defineProps({
    value: { type: Number },
    describe: { type: String },
});
</script>

<style lang="scss" scoped>
@import "@/assets/css/color.scss";
@import "@/assets/css/size.scss";

.buttonCard {
    border-radius: $cardBorder;
    text-align: center;
    display: flex;
    flex-direction: column;
    justify-content: center;
    .value {
        color: $primaryText;
        font-size: 24px;
        margin-bottom: 10px;
    }
    .describe {
        color: $secondaryText;
        font-size: 12px;
    }
}
</style>