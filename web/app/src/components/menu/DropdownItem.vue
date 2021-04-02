<template>
  <component
    :is="to.length > 0 ? 'router-link' : 'div'"
    :to="to"
    class="block px-4 py-2.5 text-sm leading-4 font-medium hover:bg-secondary-700"
  >
    <slot></slot>
  </component>
</template>

<script>
import { defineComponent } from "vue";
export default defineComponent({
  name: "dropdown-item",
  props: {
    to: {
      type: String,
      default: "",
    },
  },
});
</script>