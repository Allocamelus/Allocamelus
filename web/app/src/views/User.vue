<template>
  <div class="container py-5">
    <error-box :error="err.user" class="p-3.5 mb-3">
      <div class="flex flex-col xs:flex-row justify-between">
        <div class="flex items-center min-w-0">
          <component
            :is="canEdit ? 'change-avatar' : 'user-avatar'"
            :user="canEdit ? storeUser : user"
            class="flex-shrink-0 w-16 h-16 xs:w-20 xs:h-20"
          ></component>
          <user-name
            class="ml-3"
            :user="canEdit ? storeUser : user"
            :displayType="TwoLine"
          ></user-name>
        </div>
        <div class="mt-3 xs:mt-0 xs:ml-3 flex-shrink-0">
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
          <edit-overlay
            v-if="canEdit"
            :show="overlay"
            :user="storeUser"
            @close="overlay = false"
          ></edit-overlay>
          <sign-up-overlay
            v-if="!loggedIn"
            :show="overlay"
            :user="user"
            @close="overlay = false"
          >
          </sign-up-overlay>
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
import { InvalidCharacters } from "../components/form/errors";

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
import Overlay from "../components/overlay/Overlay.vue";
import ChangeAvatar from "../components/user/ChangeAvatar.vue";
import EditOverlay from "../components/user/EditOverlay.vue";
import SignUpOverlay from "../components/overlay/SignUpOverlay.vue";

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
      overlay: false,
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
      InvalidCharacters,
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
      this.overlay = false;
      if (this.canEdit || !this.loggedIn) {
        this.overlay = true;
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
    EditOverlay,
    SignUpOverlay,
  },
});
</script>
