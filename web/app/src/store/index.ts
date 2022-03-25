import { useStorageAsync } from "@vueuse/core";
import { defineStore } from "pinia";
import { IDBStore } from "@/pkg/idbWrapper";
import { A9sDatabase } from "@/pkg/idbWrapper/allocamelus";

const storeName = "state";
const idbStore = new IDBStore(A9sDatabase, storeName);

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
