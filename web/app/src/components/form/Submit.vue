<template>
  <basic-btn
    type="submit"
    class="btn-secondary shrink-0 px-3.5 py-2.5 transition duration-75"
    :title="title"
    @click="$emit('submit', true)"
  >
    <slot></slot>
  </basic-btn>
</template>

<script>
import { defineComponent } from "vue";
import BasicBtn from "../button/BasicBtn.vue";

export default defineComponent({
  props: {
    title: {
      type: String,
      default: "Submit",
    },
  },
  emits: ["submit"],
  components: { BasicBtn },
});
</script>
