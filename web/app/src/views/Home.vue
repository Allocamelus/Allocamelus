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
    setInterval(() => {
      this.$store.dispatch("increment");
    }, 1000);
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