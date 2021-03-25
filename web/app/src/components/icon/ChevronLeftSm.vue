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
      fill-rule="evenodd"
      d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
      clip-rule="evenodd"
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