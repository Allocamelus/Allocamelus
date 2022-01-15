<template>
  <div class="font-sans">
    <nav
      id="nav"
      class="bg-primary-600 text-gray-50 shadow z-30 m-0 p-0 fixed top-0 w-full h-nav leading-nav"
    >
      <div class="container flex flex-row justify-between h-nav leading-nav">
        <div class="flex">
          <to-link
            :to="loggedIn ? '/home' : '/'"
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
                <div
                  class="dark:bg-gray-800 bg-gray-100 overflow-x-hidden overflow-y-auto"
                >
                  <bar-loader :show="alerts.loading" />
                  <div
                    class="dark:text-white text-black max-h-48 h-48 scrollbar px-3 py-2.5"
                  >
                    <div v-if="alerts.err.length != 0">{{ alerts.err }}</div>
                    <div v-else>
                      <text-small
                        class="pb-1 text-sm font-medium text-gray-700 dark:text-gray-300"
                      >
                        Follow/Friend Request:
                      </text-small>
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
                        <div
                          class="flex flex-grow items-center justify-between"
                        >
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
                <user-avatar :user="user" class="w-6 h-6"></user-avatar>
                <component
                  v-if="!user.avatar"
                  :is="userMenu ? 'ChevronUpIcon' : 'ChevronDownIcon'"
                  class="hidden md:block w-4 h-4"
                ></component>
              </div>
              <dropdown
                v-if="!userMobile"
                v-model="userMenu"
                class="w-min whitespace-nowrap"
              >
                <div class="bg-secondary-800">
                  <dropdown-item
                    :to="`/u/${user.userName}`"
                    class="hover:bg-secondary-700"
                  >
                    <UserCircleIcon class="w-5 h-5 mr-2"></UserCircleIcon>
                    <div>Profile</div>
                  </dropdown-item>
                  <dropdown-item class="hover:bg-secondary-700">
                    <CogIcon class="w-5 h-5 mr-2"></CogIcon>
                    <div>Settings (TODO)</div>
                  </dropdown-item>
                  <dropdown-item to="/logout" class="hover:bg-secondary-700">
                    <LogoutIcon class="w-5 h-5 mr-2"></LogoutIcon>
                    <div>Logout</div>
                  </dropdown-item>
                </div>
              </dropdown>
              <Overlay
                v-else
                v-model="userMenu"
                :blockScroll="true"
                :xsFullHeight="true"
              >
                <Box
                  class="h-full w-full flex flex-col flex-grow justify-between xs:mx-2 xs:rounded-lg"
                >
                  <div class="flex flex-col">
                    <div
                      class="w-full p-3 border-b border-secondary-600 flex items-end flex-shrink-0"
                    >
                      <div class="flex-1 flex justify-end">
                        <basic-btn @click="userMenu = false">
                          <XIcon
                            class="w-5 h-5 text-black dark:text-gray-100 hover:text-gray-600 dark:hover:text-gray-300"
                          ></XIcon>
                        </basic-btn>
                      </div>
                    </div>
                    <div
                      class="flex flex-col border-b border-secondary-600 py-2"
                    >
                      <dropdown-item :to="`/u/${user.userName}`">
                        <user-avatar
                          :user="user"
                          class="w-11 h-11"
                        ></user-avatar>
                        <div
                          class="ml-3 flex flex-col flex-grow justify-evenly"
                        >
                          <user-name :user="user" :isLink="false"></user-name>
                          <div class="link">View Profile</div>
                        </div>
                      </dropdown-item>
                      <dropdown-item to="/post/new">
                        <PlusIcon class="w-5 h-5 mr-2"></PlusIcon>
                        <div>New Post</div>
                      </dropdown-item>
                    </div>
                    <div class="flex flex-col py-2">
                      <dropdown-item>
                        <CogIcon class="w-5 h-5 mr-2"></CogIcon>
                        <div>Settings (TODO)</div>
                      </dropdown-item>
                      <dropdown-item to="/logout">
                        <LogoutIcon class="w-5 h-5 mr-2"></LogoutIcon>
                        <div class="flex items-center">Logout</div>
                      </dropdown-item>
                    </div>
                  </div>
                  <BottomLinks class="justify-self-end mb-3"></BottomLinks>
                </Box>
              </Overlay>
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
    <bottom-footer v-show="footer"></bottom-footer>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, toRefs, reactive } from "vue";
import { useStore } from "./store";

import { MinToSec, SecToMs, UnixTime } from "./pkg/time";

import {
  accept as userAccept,
  decline as userDecline,
  requests,
  API_Requests,
} from "./api/user/follow";
import { SomethingWentWrong } from "./components/form/errors";

import SunIcon from "@heroicons/vue/solid/SunIcon";
import MoonIcon from "@heroicons/vue/solid/MoonIcon";
import ChevronDownIcon from "@heroicons/vue/solid/ChevronDownIcon";
import ChevronUpIcon from "@heroicons/vue/solid/ChevronUpIcon";
import BellIcon from "@heroicons/vue/outline/BellIcon";
import XIcon from "@heroicons/vue/solid/XIcon";
import UserCircleIcon from "@heroicons/vue/outline/UserCircleIcon";
import CogIcon from "@heroicons/vue/outline/CogIcon";
import LogoutIcon from "@heroicons/vue/outline/LogoutIcon";
import PlusIcon from "@heroicons/vue/outline/PlusIcon";

import Dropdown from "./components/menu/Dropdown.vue";
import DropdownItem from "./components/menu/DropdownItem.vue";
import BasicBtn from "./components/button/BasicBtn.vue";
import UserAvatar from "./components/user/Avatar.vue";
import UserName from "./components/user/Name.vue";
import ToLink from "./components/ToLink.vue";
import BarLoader from "./components/overlay/BarLoader.vue";
import Snackbar from "./components/box/Snackbar.vue";
import TextSmall from "./components/text/Small.vue";
import BottomFooter from "./components/BottomFooter.vue";
import Overlay from "./components/overlay/Overlay.vue";
import Box from "./components/box/Box.vue";
import BottomLinks from "./components/BottomLinks.vue";

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
      sesKeepAliveInterval: setInterval(() => {}, SecToMs(MinToSec(10))),
      userMenu: false,
      footer: false,
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

    (async () => sessionCheck())();
    var keepAliveDelay = async () => {
      const interval = SecToMs(MinToSec(5));
      clearInterval(data.sesKeepAliveInterval);
      setTimeout(() => {
        data.sesKeepAliveInterval = setInterval(sessionKeepAlive, interval);
      }, interval);
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
    theme(newTheme) {
      setTheme(newTheme);
    },
    $route() {
      this.onNavigate();
    },
    viewKey() {
      this.onNavigate();
    },
  },
  methods: {
    toggleUserMenu() {
      this.checkMenu();
      this.userMenu = !this.userMenu;
    },
    async clickAlerts() {
      this.alerts.menu = !this.alerts.menu;

      // limit alerts fetch to every 1 seconds
      if (this.alerts.menu && this.alerts.lastFetched < UnixTime(-1)) {
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
            this.alerts.lastFetched = UnixTime();
          })
          .catch(() => {
            this.alerts.err = SomethingWentWrong;
          })
          .finally(() => {
            this.alerts.loading = false;
          });
      }
    },
    followRequest(userId: number, accept: boolean) {
      (() => {
        var uN = this.alerts.requests.user(userId).userName;
        if (accept) {
          return userAccept(uN);
        }
        return userDecline(uN);
      })()
        .then((r) => {
          if (!r.success) {
            this.snackbarMsg(SomethingWentWrong);
            return;
          }
          let requests = this.alerts.requests.requests;
          let del = Object.keys(requests).find(
            (k) => requests[Number(k)] === userId
          );
          if (del != undefined) {
            delete requests[Number(del)];
          }
        })
        .catch(() => {
          this.snackbarMsg(SomethingWentWrong);
        });
    },
    snackbarMsg(msg: string) {
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
      if (typeof this.$route.meta.footer === "boolean") {
        this.footer = this.$route.meta.footer;
      }
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
    XIcon,
    UserCircleIcon,
    CogIcon,
    LogoutIcon,
    PlusIcon,
    BasicBtn,
    UserAvatar,
    UserName,
    ToLink,
    BarLoader,
    Snackbar,
    TextSmall,
    BottomFooter,
    Overlay,
    Box,
    BottomLinks,
  },
});
</script>

<style src="./scss/App.scss" lang="scss"></style>
