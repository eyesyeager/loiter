import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
    state: () => ({
        uid: 0,
        username: "", 
        weight: 0
    }),
    actions: {
        init(uid: number, username: string, weight: number) {
            this.$state.uid = uid;
            this.$state.username = username;
            this.$state.weight = weight;
        },
    },
})