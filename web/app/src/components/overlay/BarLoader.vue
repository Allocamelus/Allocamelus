<template>
  <div
    class="h-[3px] w-full"
    :class="show ? 'bg-rose-100 dark:bg-primary' : 'bg-transparent'"
  >
    <div
      v-if="show"
      class="h-[3px] w-1/4 animate-slide bg-rose-900 dark:bg-rose-600"
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
