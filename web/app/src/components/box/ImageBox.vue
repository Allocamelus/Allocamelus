<template>
  <div
    class="group relative flex"
    :class="[
      totalNumber == 1 ? 'w-full' : '',
      totalNumber == 2 ? (index == 0 || index == 1 ? 'w-1/2' : '') : '',
      totalNumber == 3 ? (index == 0 ? 'w-full' : 'w-1/2') : '',
      totalNumber == 4 ? 'w-1/2' : '',
    ]"
  >
    <slot></slot>
    <img
      :src="url"
      :alt="alt"
      :width="width"
      :height="height"
      class="w-full h-full object-cover"
      :loading="loading"
    />
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    url: {
      type: String,
      default: "",
    },
    index: {
      type: Number,
      default: 0,
    }, // Index
    totalNumber: {
      type: Number,
      default: 1,
    }, // How many other images are there
    loading: {
      type: String,
      default: "auto",
    },
    alt: {
      type: String,
      default: "Image",
    },
    width: {
      type: Number,
      default: 0,
    },
    height: {
      type: Number,
      default: 0,
    },
  },
});
</script>