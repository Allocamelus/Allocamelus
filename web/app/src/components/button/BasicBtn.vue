<template>
  <component
    :is="to.length > 0 ? 'to-link' : 'button'"
    :to="to"
    class="font-medium text-base leading-4 rounded cursor-pointer disabled:opacity-50"
  >
    <slot></slot>
  </component>
</template>

<script>
import { defineComponent } from "vue";
import ToLink from "../ToLink.vue";
export default defineComponent({
  components: { ToLink },
  name: "basic-button",
  props: {
    to: {
      type: String,
      default: "",
    },
  },
});
</script>
