<template>
  <button
    type="submit"
    class="submit"
    :title="title"
    @click="$emit('submit', true)"
  >
    <slot></slot>
  </button>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    title: {
      type: String,
      default: "Submit",
    },
  },
  emits: ["submit"],
});
</script>

<style lang="scss" scoped>
@layer components {
  .submit {
    @apply text-white font-medium text-base leading-4 py-2.5 px-3.5 rounded cursor-pointer;
    @apply bg-secondary-600 transition duration-75 hover:bg-secondary-800 disabled:opacity-50;
  }
}
</style>