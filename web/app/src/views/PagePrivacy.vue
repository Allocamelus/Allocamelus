<template>
  <div class="container flex flex-col py-5">
    <Box class="flex flex-col gap-6 rounded-xl p-5 text-base">
      <h1 :class="textStyle.h2">Privacy Policy</h1>
    </Box>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

import Box from "@/components/box/Box.vue";

export default defineComponent({
  setup() {
    const data = reactive({});
    document.title = `Privacy Policy - ${import.meta.env.VITE_SITE_NAME}`;
    return {
      ...toRefs(data),
    };
  },
  components: { Box },
});
</script>

<style
  src="@/scss/modules/text.modules.scss"
  lang="scss"
  module="textStyle"
  scoped
></style>
