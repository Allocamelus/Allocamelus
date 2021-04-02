<template>
  <component
    :is="to.length > 0 ? 'router-link' : 'button'"
    :to="to"
    class="font-medium text-base leading-4 rounded cursor-pointer disabled:opacity-50"
  >
    <slot></slot>
  </component>
</template>

<script>
import { defineComponent } from "vue";
export default defineComponent({
  name: "basic-button",
  props: {
    to: {
      type: String,
      default: "",
    },
  },
});
</script>
