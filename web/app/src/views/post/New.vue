<template>
  <div class="container">New Post</div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
export default defineComponent({
  setup() {
    const data = reactive({});
    document.title = `New Post - ${import.meta.env.VITE_SITE_NAME}`;
    return {
      ...toRefs(data),
    };
  },
});
</script>

<style>
</style>