<template>
  <box class="rounded-xl">
    <div v-if="error.length > 0" v-html="error"></div>
    <slot v-else></slot>
  </box>
</template>

<script>
import { defineComponent } from "vue";
import Box from "./Box.vue";
export default defineComponent({
  props: {
    error: {
      type: String,
      default: "",
    },
  },
  components: {
    Box,
  },
});
</script>