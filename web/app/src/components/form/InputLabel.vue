<template>
  <label class="text-sm font-medium" :for="label">
    <slot></slot>
    <div v-if="err?.length != 0" class="pl-2 inline text-red-600">{{ err }}</div>
  </label>
</template>

<script>
import { defineComponent } from "vue";
export default defineComponent({
  name: "input-label",
  props: {
    label: String,
    err: String
  },
});
</script>