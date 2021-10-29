<template>
  <div class="container flex py-5">Landing</div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

export default defineComponent({
  setup() {
    const data = reactive({});

    document.title = import.meta.env.VITE_SITE_NAME;

    return {
      ...toRefs(data),
    };
  },
});
</script>
