<template>
  <div class="form-box-wrapper">
    <Box class="form-box rounded-xl" :class="classes">
      <slot></slot>
    </Box>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import Box from "../box/Box.vue";

export default defineComponent({
  name: "center-form-box",
  props: {
    classes: {
      type: String,
      default: "py-4 px-5",
    },
  },
  setup() {
    const data = reactive({});
    return {
      ...toRefs(data),
    };
  },
  components: {
    Box,
  },
});
</script>

<style lang="scss" scoped>
.form-box {
  @apply mx-auto flex flex-grow flex-col;
  @apply bg-transparent shadow-none dark:bg-transparent;
  @apply xs:bg-gray-50 xs:shadow xs:dark:bg-gray-900;
  // TODO
  max-width: 400px;
}
.form-box-wrapper {
  @apply flex w-full items-center justify-center;
  min-height: calc(100vh - calc(var(--nav-height) * 1.5));
  @screen xs {
    min-height: calc(100vh - var(--nav-height));
  }
}
</style>
