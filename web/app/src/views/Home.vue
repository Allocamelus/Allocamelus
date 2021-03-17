<template>
  <div>
    <div>Home</div>
    <p>{{ count }}</p>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { onBeforeRouteUpdate } from "vue-router";

export default defineComponent({
  setup() {
    return {};
  },
  created() {
  },
  async beforeRouteUpdate(to, from, next) {
    console.log(this.$store);
    next();
  },
  computed: {
    count() {
      return this.$store.state.session.userId;
    },
  },
});
</script>

<style>
</style>