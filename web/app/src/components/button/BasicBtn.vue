<template>
  <component
    :is="notNull(to) ? 'to-link' : 'button'"
    :to="to"
    class="cursor-pointer rounded text-base leading-4 font-medium disabled:cursor-auto disabled:opacity-50"
  >
    <slot></slot>
  </component>
</template>

<script>
import { defineComponent } from "vue";

import { notNull } from "../../models/responses";

import ToLink from "../ToLink.vue";

export default defineComponent({
  name: "BasicButton",
  props: ["to"],
  setup() {
    return {
      notNull,
    };
  },
  components: { ToLink },
});
</script>
