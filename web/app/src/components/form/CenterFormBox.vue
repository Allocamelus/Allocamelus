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
  @apply flex flex-col mx-auto flex-grow;
  @apply bg-transparent shadow-none dark:bg-transparent;
  @apply xs:bg-gray-50 xs:dark:bg-gray-900 xs:shadow;
  // TODO
  max-width: 400px;
}
.form-box-wrapper {
  @apply w-full flex justify-center items-center;
  min-height: calc(100vh - calc(var(--nav-height) * 1.5));
  @screen xs {
    min-height: calc(100vh - var(--nav-height));
  }
}
</style>
