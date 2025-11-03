<template>
  <div
    class="group hover:bg-opacity-30 shrink-0 cursor-pointer rounded-full p-1"
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
