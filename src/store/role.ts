import { defineStore } from "pinia";
import { OptionsInterface } from "@/d.ts/common";
import { useUserStore } from "./user";

export const useRoleStore = defineStore("role", {
    state: () => (new Array<OptionsInterface>()),
    actions: {
        init(data: OptionsInterface[]) {
            this.$state = data;
        },
        checkAuth(threshold: string) {
            let userStore = useUserStore();
            for (let index in this.$state) {
                let option = this.$state[index];
                if (option.label == threshold) {
                    return userStore.$state.weight >= Number(option.value);
                }
            }
            return false;
        }
    },
})