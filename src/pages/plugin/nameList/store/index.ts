import { defineStore } from "pinia";

export const useSearchStore = defineStore("search", {
    state: () => ({
        flag: false,
    }),
    actions: {
        switch() {
            this.$state.flag = !this.$state.flag;
        },
    },
})