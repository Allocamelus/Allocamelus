<template>
  <router-link :to="to" @click="userEvent">
    <slot></slot>
  </router-link>
</template>

<script>
import { defineComponent } from "vue";
export default defineComponent({
  props: ["to"],
  methods: {
    userEvent() {
      if (
        this.to == this.$route.path ||
        (this.to.path != undefined && this.to.path == this.$route.path)
      ) {
        this.$store.commit("updateViewKey");
      }
    },
  },
});
</script>