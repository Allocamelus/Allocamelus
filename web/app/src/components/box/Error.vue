<template>
  <box class="rounded-xl">
    <html-errors
      v-if="error.error?.length > 0"
      class="p-3.5"
      :error="htmlErr"
    ></html-errors>
    <slot v-else></slot>
  </box>
</template>

<script>
import { defineComponent } from "vue";

import Box from "./Box.vue";
import { API_Error } from "../../models/api_error";
import { RespToHtml } from "../../models/responses";
import HtmlErrors from "../HtmlErrors.vue";

export default defineComponent({
  props: {
    error: {
      type: API_Error,
      default: new API_Error(),
    },
  },
  computed: {
    htmlErr() {
      return RespToHtml(this.error.error);
    },
  },
  components: {
    Box,
    HtmlErrors,
  },
});
</script>
