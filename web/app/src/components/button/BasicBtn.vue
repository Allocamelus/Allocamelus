<template>
  <component
    :is="notNull(to) ? 'to-link' : 'button'"
    :to="to"
    class="font-medium text-base leading-4 rounded cursor-pointer disabled:opacity-50 disabled:cursor-auto"
  >
    <slot></slot>
  </component>
</template>

<script>
import { defineComponent } from "vue";

import { notNull } from "../../models/responses";

import ToLink from "../ToLink.vue";

export default defineComponent({
  components: { ToLink },
  name: "basic-button",
  props: ["to"],
  setup() {
    return {
      notNull,
    };
  },
});
</script>
