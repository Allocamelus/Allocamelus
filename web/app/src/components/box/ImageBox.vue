<template>
  <div
    class="group relative mx-auto flex"
    :class="[
      totalNumber == 1 ? 'w-fit' : '',
      totalNumber == 2
        ? index == 0 || index == 1
          ? 'w-1/2 max-w-fit'
          : ''
        : '',
      totalNumber == 3 ? (index == 0 ? 'w-fit' : 'w-1/2 max-w-fit') : '',
      totalNumber == 4 ? 'w-1/2 max-w-fit' : '',
    ]"
  >
    <slot></slot>
    <img
      :src="fullUrl"
      :alt="alt"
      :width="width"
      :height="height"
      class="max-w-full object-cover"
      :loading="loading"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { FullURL } from "@/pkg/url";

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
      default: undefined,
    },
    height: {
      type: Number,
      default: undefined,
    },
  },
  computed: {
    fullUrl() {
      return FullURL(this.url, import.meta.env.BASE_URL);
    },
  },
});
</script>
