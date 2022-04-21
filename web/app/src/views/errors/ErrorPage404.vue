<template>
  <div class="container py-5">
    <box class="mx-auto max-w-prose rounded-xl py-3 px-4">
      <html-errors :error="404"></html-errors>
    </box>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Box from "../../components/box/Box.vue";
import HtmlErrors from "@/components/HtmlErrors.vue";

export default defineComponent({
  setup() {
    document.title = `Error 404 (Not Found) - ${
      import.meta.env.VITE_SITE_NAME
    }`;
  },
  components: { Box, HtmlErrors },
});
</script>
