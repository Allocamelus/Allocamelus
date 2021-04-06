<template>
  <div class="container py-5">
    <box class="py-3 px-4 rounded-xl max-w-prose mx-auto" v-html="errorHtml"></box>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import Box from "../../components/box/Box.vue";
import { Html404Func } from "../../components/htmlErrors";

export default defineComponent({
  computed: {
    errorHtml() {
      return Html404Func(this.$route.fullPath);
    },
  },
  components: { Box },
});
</script>
