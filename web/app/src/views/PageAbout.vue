<template>
  <div class="container flex py-5">About</div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
export default defineComponent({
  setup() {
    const data = reactive({});

    document.title = `About - ${import.meta.env.VITE_SITE_NAME}`;

    return {
      ...toRefs(data),
    };
  },
});
</script>

<style></style>
