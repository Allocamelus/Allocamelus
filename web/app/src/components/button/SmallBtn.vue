<template>
  <component
    :is="to.length > 0 ? 'to-link' : 'button'"
    :to="to"
    class="cursor-pointer rounded text-sm font-medium leading-[14px] disabled:opacity-50"
  >
    <slot></slot>
  </component>
</template>

<script>
import { defineComponent } from "vue";
import ToLink from "../ToLink.vue";
export default defineComponent({
  name: "SmallButton",
  props: {
    to: {
      type: String,
      default: "",
    },
  },
  components: { ToLink },
});
</script>
