<template>
  <div class="flex w-full grow items-center justify-center">
    <Box
      class="mx-auto flex max-w-[400px] grow flex-col rounded-xl bg-transparent shadow-none xs:bg-neutral-50 xs:shadow dark:bg-transparent xs:dark:bg-neutral-900"
      :class="classes"
    >
      <slot></slot>
    </Box>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import Box from "../box/Box.vue";

export default defineComponent({
  name: "CenterFormBox",
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
