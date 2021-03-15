<template>
  <div id="nav" class="txtC">
    <router-link to="/" class="link">Home</router-link>
    |
    <router-link to="/about" class="link">About</router-link>
    |
    <router-link to="/post/10" class="link">Post 10</router-link>
  </div>
  <router-view class="mt-5" />
</template>

<script>
import { defineComponent } from "vue";
export default defineComponent({
  setup() {
    return {};
  },
});
</script>

<style src="./scss/App.scss" lang="scss">
</style>