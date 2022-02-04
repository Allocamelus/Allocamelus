import { useStorageAsync } from "@vueuse/core";
import { defineStore } from "pinia";
import { newStore } from "@/pkg/idbWrapper";

const storeName = "state";
const idbStore = newStore(storeName);

export interface State {
  theme: string;
  viewKey: number;
}

export const useStateStore = defineStore(storeName, {
  state: () => {
    return {
      theme: useStorageAsync("theme", "dark", idbStore),
      // TODO: https://github.com/vuejs/vue-router/issues/974
      viewKey: 0,
    };
  },
  actions: {
    toggleTheme() {
      this.theme = this.theme === "dark" ? "light" : "dark";
    },
    updateViewKey() {
      this.viewKey++;
    },
  },
});
