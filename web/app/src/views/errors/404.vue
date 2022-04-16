<template>
  <div class="container py-5">
    <box
      class="mx-auto max-w-prose rounded-xl py-3 px-4"
      v-html="errorHtml"
    ></box>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Box from "../../components/box/Box.vue";
import { Html404Func } from "../../components/htmlErrors";

export default defineComponent({
  setup() {
    document.title = `Error 404 (Not Found) - ${
      import.meta.env.VITE_SITE_NAME
    }`;
  },
  computed: {
    errorHtml() {
      return Html404Func(this.$route.fullPath);
    },
  },
  components: { Box },
});
</script>
