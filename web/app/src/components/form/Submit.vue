<template>
  <basic-btn
    type="submit"
    class="text-white bg-secondary-600 py-2.5 px-3.5 transition duration-75 hover:bg-secondary-800 flex-shrink-0"
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
