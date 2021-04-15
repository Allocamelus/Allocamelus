<template>
  <div
    class="group relative flex"
    :class="[
      totalNumber == 1 ? 'w-full' : '',
      totalNumber == 2 ? [index == 0 || index == 1 ? 'w-1/2' : ''] : '',
      totalNumber == 3 ? [index == 0 ? 'w-full' : 'w-1/2'] : '',
      totalNumber == 4 ? 'w-1/2' : '',
    ]"
  >
    <slot></slot>
    <img :src="url" class="w-full h-full object-cover" />
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    url: String,
    index: Number, // Index
    totalNumber: Number, // How many other images are there
    loading: {
      type: String,
      default: "auto",
    },
  },
});
</script>