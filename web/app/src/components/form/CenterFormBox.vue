<template>
  <div class="form-box-wrapper">
    <Box class="form-box" :class="classes">
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
  setup(props) {
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
@layer components {
  .form-box {
    @apply flex flex-col mx-auto flex-grow;
    @apply xs-max:bg-transparent xs-max:shadow-none xs-max:dark:bg-transparent;
    // TODO
    max-width: 400px;
  }
  .form-box-wrapper {
    @apply w-full flex justify-center items-center;
    min-height: calc(100vh - var(--nav-height));
  }
}
</style>