<template>
  <component
    :is="to.length > 0 ? 'to-link' : 'div'"
    :to="to"
    class="flex items-center px-4 py-2.5 text-sm font-semibold leading-4"
  >
    <slot></slot>
  </component>
</template>

<script>
import { defineComponent } from "vue";
import ToLink from "../ToLink.vue";
export default defineComponent({
  components: { ToLink },
  name: "dropdown-item",
  props: {
    to: {
      type: String,
      default: "",
    },
  },
});
</script>
