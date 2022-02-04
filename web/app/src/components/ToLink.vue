<template>
  <router-link :to="to" @click="userEvent">
    <slot></slot>
  </router-link>
</template>

<script>
import { defineComponent } from "vue";
import { useStateStore } from "@/store";

export default defineComponent({
  props: ["to"],
  setup() {
    const state = useStateStore();

    return {
      updateViewKey: () => state.updateViewKey(),
    };
  },
  methods: {
    userEvent() {
      if (
        this.to == this.$route.path ||
        (this.to.path != undefined && this.to.path == this.$route.path)
      ) {
        this.updateViewKey();
      }
    },
  },
});
</script>
