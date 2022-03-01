<template>
  <div
    class="w-full h-[3px]"
    :class="show ? 'bg-rose-100 dark:bg-primary' : 'bg-transparent'"
  >
    <div
      v-if="show"
      class="w-1/4 h-[3px] bg-rose-900 dark:bg-rose-600 animate-slide"
    ></div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
export default defineComponent({
  props: {
    show: {
      type: Boolean,
      required: true,
    },
  },
});
</script>
