<template>
  <div class="p-1 rounded-full flex-shrink-0 group hover:bg-opacity-30 hover:bg-rose-800 cursor-pointer">
    <slot></slot>
  </div>
</template>

<script>
import { defineComponent } from "vue";
export default defineComponent({
  name: "circle-bg"
});
</script>