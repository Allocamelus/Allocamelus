<template>
  <div>
    <nav
      id="nav"
      class="bg-primary-600 text-gray-50 shadow z-30 m-0 p-0 fixed top-0 w-full h-nav leading-nav"
    >
      <div class="container flex flex-row justify-between h-nav leading-nav">
        <div class="flex">
          <to-link
            to="/"
            class="pr-4 py-2 text-white text-lg font-sans truncate no-underline tracking-wide relative"
          >
            Allocamelus
          </to-link>
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

            <div v-if="loggedIn" class="ml-1.5 relative">
              <div class="p-1 rounded-full cursor-pointer" @click="clickAlerts">
                <span class="sr-only">Open User Alerts</span>
                <BellIcon class="w-5.5 h-5.5"></BellIcon>
              </div>
              <dropdown v-model="alerts.menu" class="max-w-sm w-80">
                <bar-loader v-if="alerts.loading" />
                <div
                  class="dark:bg-gray-800 dark:text-white bg-gray-100 text-black max-h-48 h-48 overflow-x-hidden overflow-y-auto scrollbar px-3 py-2.5"
                >
                  <div v-if="alerts.err.length != 0">{{ alerts.err }}</div>
                  <div v-else>
                    <div
                      class="pb-1 text-sm font-medium text-gray-700 dark:text-gray-300"
                    >
                      Follow/Friend Request:
                    </div>
                    <div
                      v-for="(userId, index) in alerts.requests.requests"
                      :key="index"
                      class="pb-3 flex flex-grow flex-shrink items-center"
                    >
                      <user-avatar
                        :user="alerts.requests.user(userId)"
                        :isLink="true"
                        class="w-8 h-8"
                      ></user-avatar>
                      <div class="flex flex-grow items-center justify-between">
                        <div class="ml-2 flex">
                          <user-name
                            :user="alerts.requests.user(userId)"
                          ></user-name>
                        </div>
                        <div class="ml-2 flex items-center">
                          <div
                            class="text-sm font-semibold leading-4 rounded cursor-pointer px-2 py-1.5 text-white bg-secondary-700 hover:bg-secondary-800"
                            @click="followRequest(userId, true)"
                          >
                            Accept
                          </div>
                          <div
                            class="text-sm font-semibold leading-4 rounded cursor-pointer ml-1.5 p-1 link"
                            @click="followRequest(userId, false)"
                          >
                            Decline
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </dropdown>
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
                <!--TODO:User Mobile Menu-->
                <user-avatar
                  :user="user"
                  class="w-6 h-6"
                  :isLink="userMobile"
                ></user-avatar>
                <component
                  v-if="!user.avatar"
                  :is="userMenu ? 'ChevronUpIcon' : 'ChevronDownIcon'"
                  class="hidden md:block w-4 h-4"
                ></component>
              </div>
              <dropdown v-model="userMenu" class="w-44">
                <div class="bg-secondary-800">
                  <dropdown-item
                    :to="`/u/${user.userName}`"
                    class="hover:bg-secondary-700"
                  >
                    Profile
                  </dropdown-item>
                  <dropdown-item class="hover:bg-secondary-700"
                    >Settings (TODO)</dropdown-item
                  >
                  <dropdown-item to="/logout" class="hover:bg-secondary-700"
                    >Logout</dropdown-item
                  >
                </div>
              </dropdown>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <div id="bodyContent" class="mt-nav">
      <router-view :key="viewKey" />
      <snackbar v-model="snackbar.show" :closeBtn="true">
        {{ snackbar.msg }}
      </snackbar>
    </div>
    <!--TODO: Mobile Menu-->
    <footer id="footer">
      <div>
        <div class="copyright">&copy; {{ new Date().getFullYear() }}</div>
        <to-link to="/about" class="dash">About</to-link>
      </div>
      <div></div>
      <div>
        <!-- TODO -->
        <to-link to="/tos">Terms</to-link>
        <to-link to="/privacy" class="dash">Privacy</to-link>
      </div>
    </footer>
  </div>
</template>

<script>
import { defineComponent, computed, toRefs, reactive } from "vue";
import { useStore } from "vuex";

import { MinToSec, SecToMs } from "./pkg/time";

import {
  post as userFollow,
  remove as userUnfollow,
  requests,
  API_Requests,
} from "./api/user/follow";
import { SomethingWentWrong } from "./components/form/errors";

import SunIcon from "@heroicons/vue/solid/SunIcon";
import MoonIcon from "@heroicons/vue/solid/MoonIcon";
import ChevronDownIcon from "@heroicons/vue/solid/ChevronDownIcon";
import ChevronUpIcon from "@heroicons/vue/solid/ChevronUpIcon";
import BellIcon from "@heroicons/vue/outline/BellIcon";
import Dropdown from "./components/menu/Dropdown.vue";
import DropdownItem from "./components/menu/DropdownItem.vue";
import BasicBtn from "./components/button/BasicBtn.vue";
import UserAvatar from "./components/user/Avatar.vue";
import UserName from "./components/user/Name.vue";
import ToLink from "./components/ToLink.vue";
import BarLoader from "./components/overlay/BarLoader.vue";
import Snackbar from "./components/box/Snackbar.vue";

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
      user = computed(() => store.getters.user),
      theme = computed(() => store.getters.theme),
      viewKey = computed(() => store.getters.viewKey),
      toggleTheme = () => store.commit("toggleTheme"),
      sessionCheck = () => store.dispatch("sessionCheck"),
      sessionKeepAlive = () => store.dispatch("sessionKeepAlive");
    const data = reactive({
      sesKeepAliveInterval: null,
      userMenu: false,
      alerts: {
        err: "",
        menu: false,
        loading: false,
        lastFetched: 0,
        requests: new API_Requests(),
      },
      snackbar: {
        show: false,
        msg: "",
      },
      userMobile: window.screen.width < 768,
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
      user,
      theme,
      viewKey,
      toggleTheme,
      sessionCheck,
    };
  },
  watch: {
    theme(newTheme, _old) {
      setTheme(newTheme);
    },
    $route(_to, _from) {
      this.onNavigate();
    },
    viewKey(_newKey, _old) {
      this.onNavigate();
    },
  },
  methods: {
    toggleUserMenu() {
      this.checkMenu();
      if (!this.userMobile) {
        this.userMenu = !this.userMenu;
      } else {
        this.userMenu = false;
      }
    },
    async clickAlerts() {
      this.alerts.menu = !this.alerts.menu;
      this.alerts.loading = true;
      this.alerts.err = "Loading...";
      requests()
        .then((r) => {
          this.alerts.err = "";
          if (Object.keys(r.requests).length != 0) {
            this.alerts.requests = r;
          } else {
            this.alerts.err = "No Notifications";
          }
        })
        .catch((_e) => {
          this.alerts.err = SomethingWentWrong;
        })
        .finally(() => {
          this.alerts.loading = false;
        });
    },
    followRequest(userId, accept) {
      (() => {
        var uN = this.alerts.requests.user(userId).userName;
        if (accept) {
          return userFollow(uN);
        }
        return userUnfollow(uN);
      })()
        .then((r) => {
          if (!r.success) {
            this.snackbarMsg(SomethingWentWrong);
            return;
          }
          var requests = this.alerts.requests.requests;
          delete requests[
            Object.keys(requests).find((k) => requests[k] === userId)
          ];
        })
        .catch((_e) => {
          this.snackbarMsg(SomethingWentWrong);
        });
    },
    snackbarMsg(msg) {
      this.snackbar.msg = "";
      if (msg.length > 0) {
        this.snackbar.msg = msg;
        this.snackbar.show = true;
      }
    },
    checkMenu() {
      this.userMobile = screen.width < 768;
    },
    onNavigate() {
      this.checkMenu();
      this.userMenu = false;
      this.alerts.menu = false;
    },
  },
  components: {
    SunIcon,
    MoonIcon,
    Dropdown,
    DropdownItem,
    BellIcon,
    ChevronDownIcon,
    ChevronUpIcon,
    BasicBtn,
    UserAvatar,
    UserName,
    ToLink,
    BarLoader,
    Snackbar,
  },
});
</script>

<style src="./scss/index.scss" lang="scss">
</style>

<style src="./scss/App.scss" lang="scss">
</style>