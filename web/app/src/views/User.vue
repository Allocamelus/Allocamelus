<template>
  <div class="container py-5">
    <error-box :error="err.user" class="p-3.5 mb-3">
      <div class="flex flex-col xs:flex-row justify-between">
        <div class="flex items-center">
          <component
            :is="canEdit ? 'change-avatar' : 'user-avatar'"
            :user="user"
            class="flex-shrink-0 w-16 h-16 xs:w-20 xs:h-20"
          ></component>
          <user-name
            class="ml-2"
            :user="user"
            :displayType="TwoLine"
          ></user-name>
        </div>
        <div class="mt-3 xs:mt-0 xs:ml-2">
          <basic-btn
            class="px-3 py-2 border whitespace-nowrap"
            :class="[
              'border-secondary-700 text-secondary-700 dark:text-rose-600 ',
              'hover:bg-secondary-700 hover:text-white dark:hover:text-white',
            ]"
            @click="clickFollowEdit"
          >
            {{ canEdit ? "Edit Profile" : "Follow" }}
          </basic-btn>
          <overlay v-model="followEditOverlay">
            <box
              class="w-full xs-max:h-full xs:m-3 rounded-none xs:rounded-md shadow-lg bg-secondary-800 focus:outline-none overflow-hidden flex flex-col"
            >
              <div
                class="w-full p-3 border-b border-secondary-600 flex items-end"
              >
                <div
                  class="flex-1 flex"
                  :class="canEdit ? 'justify-start' : 'justify-end'"
                >
                  <basic-btn @click="followEditOverlay = false">
                    <XIcon class="w-5 h-5"></XIcon>
                  </basic-btn>
                </div>
                <div v-if="canEdit" class="flex-1 flex justify-center">
                  <div class="font-medium text-base leading-4">
                    Edit Profile
                  </div>
                </div>
                <div v-if="canEdit" class="flex-1 flex justify-end">
                  <basic-btn>Save</basic-btn>
                </div>
              </div>
              <div
                class="flex-grow flex"
                :class="!canEdit ? 'items-center justify-center' : ''"
              >
                <div
                  v-if="!canEdit"
                  class="text-center flex flex-col py-8 px-6 xs:px-8"
                >
                  <div class="text-xl font-medium flex">
                    Sign Up or Login to Follow {{ user.name }}
                    <div
                      class="pl-1 font-normal text-gray-700 dark:text-gray-400"
                    >
                      @{{ user.userName }}
                    </div>
                  </div>
                  <div class="flex flex-col items-center mt-5">
                    <basic-btn
                      to="/signup"
                      class="w-full text-white bg-secondary-700 hover:bg-secondary-800 py-2 px-3.5 mb-4"
                    >
                      Sign Up
                    </basic-btn>
                    <basic-btn
                      to="/login"
                      class="w-full py-2 px-3 link border border-rose-800 dark:border-rose-500 hover:border-rose-900 dark:hover:border-rose-600"
                    >
                      Login
                    </basic-btn>
                  </div>
                </div>
                <div v-else class="flex flex-grow flex-col py-6 px-6 xs:px-8">
                  <div class="flex items-center mt-2">
                    <user-avatar
                      class="h-11 w-11"
                      :user="user"
                      :isLink="false"
                    ></user-avatar>
                    <change-avatar
                      class="w-full ml-2"
                      :user="user"
                      :blockScrool="false"
                    >
                      <basic-btn class="link">Change Avatar</basic-btn>
                    </change-avatar>
                  </div>
                </div>
              </div>
            </box>
          </overlay>
        </div>
      </div>
      <div></div>
      <div>{{ user.bio }}</div>
    </error-box>
    <div class="flex">
      <feed>
        <div v-if="err.posts.length > 0" v-html="err.posts"></div>
        <post-feed :list="postsList"></post-feed>
      </feed>
      <sidebar></sidebar>
    </div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive, computed } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";

import { get as getUser } from "../api/user/get";
import { posts as getPosts } from "../api/user/posts";
import { API_Error } from "../models/api_error";
import { API_Posts } from "../models/api_posts";
import { Html404Func, HtmlSomethingWentWrong } from "../components/htmlErrors";

import { GEN_User } from "../models/go_structs_gen";
import ApiResp from "../models/responses";

import XIcon from "@heroicons/vue/solid/XIcon";

import UserName, { TwoLine } from "../components/user/Name.vue";
import ErrorBox from "../components/box/Error.vue";
import PostFeed from "../components/post/Feed.vue";
import Feed from "../components/Feed.vue";
import Sidebar from "../components/Sidebar.vue";
import Box from "../components/box/Box.vue";
import UserAvatar from "../components/user/Avatar.vue";
import BasicBtn from "../components/button/BasicBtn.vue";
import Overlay from "../components/box/Overlay.vue";
import ChangeAvatar from "../components/user/ChangeAvatar.vue";

function userErrors(api_error, path) {
  if (api_error instanceof API_Error) {
    switch (api_error.error) {
      case ApiResp.Shared.NotFound:
        return Html404Func(path);
    }
  }
  return HtmlSomethingWentWrong;
}
export default defineComponent({
  props: {
    userName: {
      type: Array,
      required: true,
    },
  },
  setup(props) {
    const route = useRouter();
    const store = useStore();
    const loggedIn = computed(() => store.getters.loggedIn),
      storeUser = computed(() => store.getters.user);
    const data = reactive({
      user: new GEN_User(),
      postsList: new API_Posts(),
      followEditOverlay: false,
      page: 1,
      err: {
        user: "",
        posts: "",
      },
    });

    getUser(props.userName[0])
      .then((r) => {
        data.user = r;
      })
      .catch((e) => {
        data.err.user = userErrors(e, route.currentRoute.value.fullPath);
      });

    getPosts(props.userName[0], data.page)
      .then((r) => {
        data.postsList = r;
        if (Object.keys(r.posts).length == 0) {
          data.err.posts = "No posts here";
        }
      })
      .catch((e) => {
        data.err.posts = userErrors(e, route.currentRoute.value.fullPath);
      });

    return {
      ...toRefs(data),
      TwoLine,
      loggedIn,
      storeUser,
    };
  },
  computed: {
    canEdit() {
      if (this.loggedIn) {
        if (this.storeUser.id == this.user.id) {
          return true;
        }
      }
      return false;
    },
  },
  watch: {
    user(newUser, old) {
      document.title = `${newUser.name} (@${newUser.userName}) - ${
        import.meta.env.VITE_SITE_NAME
      }`;
    },
  },
  methods: {
    clickFollowEdit() {
      if (this.canEdit || !this.loggedIn) {
        this.followEditOverlay = !this.followEditOverlay;
      }
    },
  },
  async beforeRouteUpdate(to, from) {
    this.user = new GEN_User();
    this.postsList = new API_Posts();
    this.page = 1;

    getUser(to.params.userName[0])
      .then((r) => {
        this.user = r;
      })
      .catch((e) => {
        this.err.user = userErrors(e, this.$route.currentRoute.value.fullPath);
      });

    getPosts(to.params.userName[0], this.page)
      .then((r) => {
        this.postsList = r;
        if (Object.keys(r.posts).length == 0) {
          this.err.posts = "No posts here";
        }
      })
      .catch((e) => {
        this.err.posts = userErrors(e, this.$route.currentRoute.value.fullPath);
      });
  },
  components: {
    UserName,
    ErrorBox,
    PostFeed,
    Feed,
    Sidebar,
    Box,
    UserAvatar,
    BasicBtn,
    Overlay,
    XIcon,
    ChangeAvatar,
  },
});
</script>
