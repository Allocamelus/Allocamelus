<template>
  <basic-btn
    type="submit"
    class="flex-shrink-0 px-3.5 py-2.5 transition duration-75"
    :class="buttonStyle.secondary"
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

<style
  src="@/scss/modules/button.modules.scss"
  lang="scss"
  module="buttonStyle"
  scoped
></style>
