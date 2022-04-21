<template>
  <div class="container flex flex-col py-5">Terms</div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
export default defineComponent({
  setup() {
    const data = reactive({});

    document.title = `Terms - ${import.meta.env.VITE_SITE_NAME}`;

    return {
      ...toRefs(data),
    };
  },
});
</script>

<style></style>
