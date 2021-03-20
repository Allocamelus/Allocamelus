<template>
  <nav id="nav">
    <div class="flex">
      <router-link to="/" class="px-4">Allocamelus</router-link>
    </div>
    <div class="flex">
      <div class="fa-adjust-wrapper">
        <i
          class="fas fa-adjust"
          :class="toggleButtonClass"
          @click="toggleTheme"
        ></i>
      </div>
      <div class="flex">
        <router-link to="/login">Login</router-link>
      </div>
    </div>
  </nav>
  <div id="bodyContent">
    <router-view class="pt-5" />
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
import { MinToSec, SecToMs } from "./models/time";

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
  computed: {
    toggleButtonClass() {
      if (this.theme != "dark") {
        return "fa-flip-horizontal";
      }
      return "";
    },
  },
});
</script>

<style src="./scss/App.scss" lang="scss">
</style>