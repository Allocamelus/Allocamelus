<template>
  <div>About</div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
export default defineComponent({
  setup() {
    const data = reactive({});
    return {
      ...toRefs(data),
    };
  },
});
</script>

<style></style>
