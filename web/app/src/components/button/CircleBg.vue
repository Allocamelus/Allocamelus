<template>
  <div
    class="group flex-shrink-0 cursor-pointer rounded-full p-1 hover:bg-opacity-30"
  >
    <slot></slot>
  </div>
</template>

<script>
import { defineComponent } from "vue";
export default defineComponent({
  name: "CircleBg",
});
</script>
