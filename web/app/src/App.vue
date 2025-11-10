<template>
  <div class="flex min-h-screen min-w-full grow flex-col font-sans">
    <nav
      id="nav"
      class="fixed top-0 z-30 m-0 h-nav w-full bg-primary p-0 leading-nav text-neutral-50 shadow"
    >
      <div class="container flex h-nav flex-row justify-between leading-nav">
        <div class="flex">
          <to-link
            :to="loggedIn ? '/home' : '/'"
            class="relative truncate py-2 pr-4 font-sans text-lg tracking-wide text-white no-underline"
          >
            Allocamelus
          </to-link>
        </div>
        <div class="flex">
          <div class="ml-1 flex items-center justify-start">
            <div class="cursor-pointer rounded-full p-1" @click="toggleTheme">
              <span class="sr-only">Toggle Theme</span>
              <component
                :is="theme != 'dark' ? 'MoonIcon' : 'SunIcon'"
                class="h-5.5 w-5.5"
              ></component>
            </div>

            <div v-if="loggedIn" class="relative ml-1.5">
              <div class="cursor-pointer rounded-full p-1" @click="clickAlerts">
                <span class="sr-only">Open User Alerts</span>
                <BellIcon class="h-5.5 w-5.5"></BellIcon>
              </div>
              <dropdown v-model="alerts.menu" class="w-80 max-w-sm">
                <div
                  class="overflow-x-hidden overflow-y-auto bg-neutral-100 dark:bg-neutral-800"
                >
                  <bar-loader :show="alerts.loading" />
                  <div
                    class="scrollbar h-48 max-h-48 px-3 py-2.5 text-black dark:text-white"
                  >
                    <div v-if="alerts.err.length != 0" class="text-base">
                      {{ alerts.err }}
                    </div>
                    <div v-else>
                      <div
                        v-for="(userId, index) in alerts.requests.requests"
                        :key="index"
                        class="flex shrink grow pb-3"
                      >
                        <div class="flex h-10 items-center">
                          <user-avatar
                            :user="alerts.requests.user(userId)"
                            :is-link="true"
                            class="h-8 w-8"
                          ></user-avatar>
                        </div>
                        <div class="flex grow flex-col justify-between">
                          <div class="ml-2 flex flex-col pb-2">
                            <div class="flex">
                              <div class="flex w-0 grow text-base leading-5">
                                <user-name
                                  :user="alerts.requests.user(userId)"
                                ></user-name>
                              </div>
                            </div>
                            <text-small class="font-medium">
                              Sent You a
                              {{
                                user.type === Private &&
                                alerts.requests.user(userId).type === Private
                                  ? "Friend"
                                  : "Follow"
                              }}
                              Request!
                            </text-small>
                          </div>
                          <div class="ml-2 flex items-center">
                            <div
                              class="btn-secondary cursor-pointer rounded px-2 py-1.5 text-sm leading-4 font-semibold"
                              @click="followRequest(userId, true)"
                            >
                              Accept
                            </div>
                            <div
                              class="link ml-1.5 cursor-pointer rounded p-1 text-sm leading-4 font-semibold"
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

            <div v-if="!loggedIn" class="mx-2 flex items-center justify-start">
              <basic-btn to="/signup" class="border border-white px-3 py-2">
                Sign Up
              </basic-btn>
              <basic-btn to="/login" class="ml-1.5 hidden py-2 pl-3 xs:block">
                Login
              </basic-btn>
            </div>
            <div v-else class="relative ml-1.5">
              <div
                class="flex cursor-pointer items-center p-1"
                @click="toggleUserMenu"
              >
                <span class="sr-only">Open user menu</span>
                <user-avatar :user="user" class="h-6 w-6"></user-avatar>
                <component
                  :is="userMenu ? 'ChevronUpIcon' : 'ChevronDownIcon'"
                  v-if="!user.avatar"
                  class="hidden h-4 w-4 md:block"
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
                    <UserCircleIcon class="mr-2 h-5 w-5"></UserCircleIcon>
                    <div>Profile</div>
                  </dropdown-item>
                  <dropdown-item class="hover:bg-secondary-700">
                    <CogIcon class="mr-2 h-5 w-5"></CogIcon>
                    <div>Settings (TODO)</div>
                  </dropdown-item>
                  <dropdown-item to="/logout" class="hover:bg-secondary-700">
                    <LogoutIcon class="mr-2 h-5 w-5"></LogoutIcon>
                    <div>Logout</div>
                  </dropdown-item>
                </div>
              </dropdown>
              <Overlay
                v-else
                v-model="userMenu"
                :block-scroll="true"
                :xs-full-height="true"
              >
                <Box
                  class="flex h-full w-full grow flex-col justify-between xs:mx-2 xs:rounded-lg"
                >
                  <div class="flex flex-col">
                    <div
                      class="flex w-full shrink-0 items-end border-b border-secondary-600 p-3"
                    >
                      <div class="flex flex-1 justify-end">
                        <basic-btn @click="userMenu = false">
                          <XMarkIcon class="xIcon h-5 w-5"></XMarkIcon>
                        </basic-btn>
                      </div>
                    </div>
                    <div
                      class="flex flex-col border-b border-secondary-600 py-2"
                    >
                      <dropdown-item :to="`/u/${user.userName}`">
                        <user-avatar
                          :user="user"
                          class="h-11 w-11"
                        ></user-avatar>
                        <div class="ml-3 flex grow flex-col justify-evenly">
                          <user-name :user="user" :is-link="false"></user-name>
                          <div class="link">View Profile</div>
                        </div>
                      </dropdown-item>
                      <dropdown-item to="/post/new">
                        <PlusIcon class="mr-2 h-5 w-5"></PlusIcon>
                        <div>New Post</div>
                      </dropdown-item>
                    </div>
                    <div class="flex flex-col py-2">
                      <dropdown-item>
                        <CogIcon class="mr-2 h-5 w-5"></CogIcon>
                        <div>Settings (TODO)</div>
                      </dropdown-item>
                      <dropdown-item to="/logout">
                        <LogoutIcon class="mr-2 h-5 w-5"></LogoutIcon>
                        <div class="flex items-center">Logout</div>
                      </dropdown-item>
                    </div>
                  </div>
                  <BottomLinks class="mb-3 justify-self-end"></BottomLinks>
                </Box>
              </Overlay>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <div class="mt-nav flex flex-1 grow flex-col">
      <router-view :key="viewKey" />
      <snackbar v-model="snackbar.show" :close-btn="true">
        {{ snackbar.msg }}
      </snackbar>
    </div>
    <!--TODO: Mobile Menu-->
    <bottom-footer v-show="footer"></bottom-footer>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, toRefs, reactive } from "vue";
import { useStateStore } from "@/store";
import { useSessionStore } from "@/store/session";

import { MinToSec, SecToMs, UnixTime } from "./pkg/time";

import {
  accept as userAccept,
  decline as userDecline,
  requests,
  API_Requests,
} from "./api/user/follow";
import { SomethingWentWrong } from "./components/form/errors";

import {
  SunIcon,
  MoonIcon,
  ChevronDownIcon,
  ChevronUpIcon,
  XMarkIcon,
  ArrowRightStartOnRectangleIcon as LogoutIcon,
} from "@heroicons/vue/20/solid";

import {
  BellIcon,
  UserCircleIcon,
  CogIcon,
  PlusIcon,
} from "@heroicons/vue/24/outline";

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
import { Private } from "./models/user";

function setTheme(theme = "dark") {
  if (theme == "dark") {
    document.documentElement.classList.add("dark");
  } else {
    document.documentElement.classList.remove("dark");
  }
}

export default defineComponent({
  setup() {
    const state = useStateStore(),
      session = useSessionStore();
    const theme = computed(() => state.theme);
    const data = reactive({
      sesKeepAliveInterval: setInterval(
        () => {
          return;
        },
        SecToMs(MinToSec(10))
      ),
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

    (async () => session.getStatus())();
    var keepAliveDelay = async () => {
      const interval = SecToMs(MinToSec(5));
      clearInterval(data.sesKeepAliveInterval);
      setTimeout(() => {
        data.sesKeepAliveInterval = setInterval(session.keepAlive, interval);
      }, interval);
    };
    keepAliveDelay();

    setTheme(theme.value);
    return {
      ...toRefs(data),
      loggedIn: computed(() => session.loggedIn),
      user: computed(() => session.user),
      theme,
      viewKey: computed(() => state.viewKey),
      toggleTheme: state.toggleTheme,
      Private,
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
    XMarkIcon,
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

<style src="./main.css" lang="css"></style>
