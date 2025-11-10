<template>
  <div class="text-sm text-neutral-700 dark:text-neutral-400">
    <slot></slot>
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "TextSmall",
});
</script>
