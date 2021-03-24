<template>
  <!-- TODO: license https://github.com/tailwindlabs/heroicons/blob/master/LICENSE -->
  <svg
    class="fill-current"
    :class="sizeClass"
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 20 20"
    fill="currentColor"
  >
    <path
      d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z"
    />
  </svg>
</template>

<script>
import { defineComponent } from "vue";
export default defineComponent({
  props: {
    sizeClass: {
      type: String,
      default: "w-5 h-5",
    },
  },
});
</script>