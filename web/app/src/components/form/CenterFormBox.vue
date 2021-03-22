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
@import "@/scss/vars";

.form-box {
  max-width: 400px;
  margin: 0 auto;
  position: relative;
  top: calc(50% - #{$nav-height});
  transform: translateY(-50%);
  display: flex;
  flex-direction: column;
}
.form-box-wrapper {
  width: 100%;
  height: calc(100vh - #{$nav-height});
}
@media (max-width: 425px) {
  .form-box,
  .dark-theme .form-box {
    background-color: transparent;
    box-shadow: unset;
  }
}
</style>