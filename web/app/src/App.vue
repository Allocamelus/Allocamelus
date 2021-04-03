<template>
  <div>
    <nav
      id="nav"
      class="bg-primary-600 text-gray-50 shadow z-30 m-0 p-0 fixed top-0 w-full h-nav leading-nav"
    >
      <div class="container flex flex-row justify-between h-nav leading-nav">
        <div class="flex">
          <router-link
            to="/"
            class="pr-4 py-2 text-white text-lg font-sans truncate no-underline tracking-wide relative"
            >Allocamelus</router-link
          >
        </div>
        <div class="flex">
          <div class="flex justify-start items-center ml-1">
            <div class="p-1 rounded-full cursor-pointer" @click="toggleTheme">
              <span class="sr-only">Toggle Theme</span>
              <component
                :is="this.theme != 'dark' ? 'MoonIcon' : 'SunIcon'"
                class="w-5.5 h-5.5"
              ></component>
            </div>

            <div v-if="!loggedIn" class="flex justify-start items-center mx-2">
              <basic-btn to="/signup" class="border border-white py-2 px-3">
                Sign Up
              </basic-btn>
              <basic-btn to="/login" class="ml-1.5 py-2 pl-3">
                Login
              </basic-btn>
            </div>
            <div v-else class="ml-1.5 relative">
              <div
                class="p-1 cursor-pointer flex items-center"
                @click="toggleUserMenu"
              >
                <span class="sr-only">Open user menu</span>
                <!--TODO:User Image-->
                <UserCircleIcon class="w-5.5 h-5.5"></UserCircleIcon>
                <component
                  :is="userMenu ? 'ChevronUpIcon' : 'ChevronDownIcon'"
                  class="hidden md:block w-4 h-4"
                ></component>
              </div>
              <dropdown v-model="userMenu" class="w-44">
                <dropdown-item :to="`/u/${userName}`">Profile</dropdown-item>
                <dropdown-item>Settings (TODO)</dropdown-item>
                <dropdown-item to="/logout">Logout</dropdown-item>
              </dropdown>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <div id="bodyContent" class="mt-nav">
      <router-view />
    </div>
    <!--TODO: Mobile Menu-->
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
  </div>
</template>

<script>
import { defineComponent, computed, toRefs, reactive } from "vue";
import { useStore } from "vuex";

import SunIcon from "@heroicons/vue/solid/SunIcon";
import MoonIcon from "@heroicons/vue/solid/MoonIcon";
import UserCircleIcon from "@heroicons/vue/outline/UserCircleIcon";
import ChevronDownIcon from "@heroicons/vue/solid/ChevronDownIcon";
import ChevronUpIcon from "@heroicons/vue/solid/ChevronUpIcon";
import Dropdown from "./components/menu/Dropdown.vue";
import DropdownItem from "./components/menu/DropdownItem.vue";
import BasicBtn from "./components/button/BasicBtn.vue";

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
      userName = computed(() => store.getters.userName),
      theme = computed(() => store.getters.theme),
      toggleTheme = () => store.commit("toggleTheme"),
      sessionCheck = () => store.dispatch("sessionCheck"),
      sessionKeepAlive = () => store.dispatch("sessionKeepAlive");
    const data = reactive({
      sesKeepAliveInterval: null,
      userMenu: false,
      userMenuMobile: false,
    });

    sessionCheck();
    var keepAliveDelay = async () => {
      const interval = SecToMs(MinToSec(5));
      setTimeout(
        (data.sesKeepAliveInterval = setInterval(sessionKeepAlive, interval)),
        interval
      );
    };
    keepAliveDelay();

    setTheme(theme.value);
    return {
      ...toRefs(data),
      loggedIn,
      userName,
      theme,
      toggleTheme,
      sessionCheck,
    };
  },
  watch: {
    theme(newTheme, old) {
      setTheme(newTheme);
    },
    $route(to, from) {
      this.userMenu = false;
      this.userMenuMobile = false;
    },
  },
  methods: {
    toggleUserMenu() {
      if (screen.width >= 768) {
        this.userMenu = !this.userMenu;
        this.userMenuMobile = false;
      } else {
        this.userMenuMobile = !this.userMenuMobile;
        this.userMenu = false;
      }
    },
  },
  components: {
    SunIcon,
    MoonIcon,
    UserCircleIcon,
    Dropdown,
    DropdownItem,
    ChevronDownIcon,
    ChevronUpIcon,
    BasicBtn,
  },
});
</script>

<style src="./scss/index.scss" lang="scss">
</style>

<style src="./scss/App.scss" lang="scss">
</style>