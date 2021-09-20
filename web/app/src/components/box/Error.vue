<template>
  <box class="rounded-xl">
    <div v-if="error.error?.length > 0" class="p-3.5" v-html="htmlErr"></div>
    <slot v-else></slot>
  </box>
</template>

<script>
import { defineComponent } from "vue";

import Box from "./Box.vue";
import { API_Error } from "../../models/api_error";
import { RespToHtml } from "../../models/responses";

export default defineComponent({
  props: {
    error: {
      type: API_Error,
      default: new API_Error(),
    },
  },
  components: {
    Box,
  },
  computed: {
    htmlErr() {
      return RespToHtml(this.error.error);
    },
  },
});
</script>
