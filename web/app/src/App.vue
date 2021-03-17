<template>
  <div id="nav">
    <div>
      <router-link to="/" class="px-4">Allocamelus</router-link>
    </div>
    <div class="flex">
      <div class="fa-adjust-wrapper">
        <i class="fas fa-adjust" :class="toggleButtonClass" @click="toggleTheme"></i>
      </div>
      <ul>
        <router-link to="/about">About</router-link>
        <router-link to="/login">Login</router-link>
        <router-link to="/posts">Posts</router-link>
      </ul>
    </div>
  </div>
  <router-view class="pt-5" />
</template>

<script>
import { defineComponent, computed } from "vue";
import { useStore } from "vuex";

function setTheme(theme = "dark") {
  if (theme == "dark") {
    document.body.classList.add("dark-theme");
  } else {
    document.body.classList.remove("dark-theme");
  }
}

export default defineComponent({
  setup() {
    const store = useStore(),
      loggedIn = computed(() => store.getters.loggedIn),
      theme = computed(() => store.getters.theme),
      toggleTheme = () => store.commit("toggleTheme");
    setTheme(theme.value);
    return {
      loggedIn,
      theme,
      toggleTheme,
    };
  },
  watch: {
    theme(newTheme, old) {
      setTheme(this.theme)
    },
  },
  computed: {
    toggleButtonClass() {
      if (this.theme != 'dark') {
        return "fa-flip-horizontal"
      }
      return ""
    }
  }
});
</script>

<style src="./scss/App.scss" lang="scss">
</style>