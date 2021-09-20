<template>
  <component
    :is="to.length > 0 ? 'to-link' : 'div'"
    :to="to"
    class="block px-4 py-2.5 text-sm leading-4 font-medium"
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
