<template>
  <button
    type="submit"
    class="submit"
    :title="title"
    @click="$emit('submit', true)"
  >
    <slot></slot>
  </button>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    title: {
      type: String,
      default: "Submit",
    },
  },
  emits: ["submit"],
});
</script>

<style lang="scss" scoped>
@import "@/scss/vars";

.submit {
  background-color: $secondary-color;
  color: #fff;
  font-weight: 600;
  font-size: 15px;
  padding: 8px 12px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  -webkit-transition: 0.15s;
  transition: 0.15s;
}
.submit:hover {
  background-color: $secondary-hover-color;
}
</style>