<template>
  <div id="nav">
    <ul>
      <li>
        <router-link to="/" class="link">Home</router-link>
      </li>
      <li>
        <router-link to="/about" class="link">About</router-link>
      </li>
      <li>
        <router-link to="/login" class="link">Login</router-link>
      </li>
      <li>
        <router-link to="/posts" class="link">Posts</router-link>
      </li>
    </ul>
  </div>
  <router-view class="pt-5" />
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