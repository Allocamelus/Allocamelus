<template>
  <label :for="label">
    <slot></slot>
    <div v-if="err != undefined" class="pl-2 inline error">{{ err }}</div>
  </label>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
export default defineComponent({
  name: "input-label",
  props: {
    label: String,
    err: String
  },
});
</script>