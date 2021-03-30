<template>
  <nav id="nav">
    <div>
      <router-link to="/" class="px-4">Allocamelus</router-link>
    </div>
    <div>
      <div class="flex justify-start items-center px-1">
        <div class="p-2 rounded-3xl cursor-pointer">
          <component
            :is="this.theme != 'dark' ? 'MoonIcon' : 'SunIcon'"
            class="w-5 h-5"
            @click="toggleTheme"
          ></component>
        </div>
      </div>
      <div>
        <router-link to="/login">Login</router-link>
      </div>
    </div>
  </nav>
  <div id="bodyContent">
    <router-view />
  </div>
  <footer id="footer">
    <div>
      <div class="copyright">&copy; {{ new Date().getFullYear() }}</div>
      <router-link to="/about" class="dash">About</router-link>
    </div>
    <div></div>
    <div>
      <!-- TODO -->
      <router-link to="/tos">Terms</router-link>
      <router-link to="/privacy" class="dash">Privacy</router-link>
    </div>
  </footer>
</template>

<script>
import { defineComponent, computed, toRefs, reactive } from "vue";
import { useStore } from "vuex";

import SunIcon from '@heroicons/vue/solid/SunIcon';
import MoonIcon from '@heroicons/vue/solid/MoonIcon'

import { MinToSec, SecToMs } from "./pkg/time";

function setTheme(theme = "dark") {
  if (theme == "dark") {
    document.documentElement.classList.add("dark");
  } else {
    document.documentElement.classList.remove("dark");
  }
}

export default defineComponent({
  setup() {
    const store = useStore(),
      loggedIn = computed(() => store.getters.loggedIn),
      theme = computed(() => store.getters.theme),
      toggleTheme = () => store.commit("toggleTheme"),
      sessionCheck = () => store.dispatch("sessionCheck"),
      sessionKeepAlive = () => store.dispatch("sessionKeepAlive");
    const data = reactive({
      sesKeepAliveInterval: null,
    });

    sessionCheck();
    var keepAliveDelay = async () => {
      const interval = SecToMs(MinToSec(5));
      setTimeout(
        (data.sessionKAInter = setInterval(sessionKeepAlive, interval)),
        interval
      );
    };
    keepAliveDelay();

    setTheme(theme.value);
    return {
      ...toRefs(data),
      loggedIn,
      theme,
      toggleTheme,
      sessionCheck,
    };
  },
  watch: {
    theme(newTheme, old) {
      setTheme(this.theme);
    },
  },
  components: { SunIcon, MoonIcon },
});
</script>

<style src="./scss/index.scss" lang="scss">
</style>

<style src="./scss/App.scss" lang="scss">
</style>