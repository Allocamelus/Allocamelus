<template>
  <div class="container py-5">
    <error-box :error="err.user" class="p-3.5 mb-3">
      <div class="flex flex-col xs:flex-row justify-between">
        <div>
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
            <text-small
              v-if="this.user.type === UNVERIFIED_USER"
              class="font-normal flex-none ml-1"
            >
              [[Unverified]]
            </text-small>
          </div>
          <div class="mt-3 text-lg">
            {{ canEdit ? storeUser.bio : user.bio }}
          </div>
        </div>
        <div
          class="mt-3 xs:mt-0 xs:ml-3 flex-shrink-0 flex justify-end items-start"
        >
          <basic-btn
            class="px-3 py-2 border whitespace-nowrap"
            :class="[
              'border-secondary-700 text-secondary-700 dark:text-rose-600 ',
              'hover:bg-secondary-700 hover:text-white dark:hover:text-white',
            ]"
            @click="clickFollowEdit"
          >
            {{ followEditBtnTxt }}
          </basic-btn>
        </div>
        <edit-overlay
          v-if="canEdit"
          :show="overlay"
          :user="storeUser"
          @close="overlay = false"
        ></edit-overlay>
        <sign-up-overlay
          v-if="!loggedIn"
          :show="overlay"
          :redirect="`/u/${user.name}`"
          @close="overlay = false"
        >
          <div>Sign Up or Login to Follow {{ user.name }}</div>
          <div class="pl-1 font-normal text-gray-700 dark:text-gray-400">
            @{{ user.userName }}
          </div>
        </sign-up-overlay>
      </div>
    </error-box>
    <div class="flex">
      <feed>
        <new-post-text-input v-if="canEdit"></new-post-text-input>
        <error-box :error="err.posts">
          <box v-if="postsList.total() == 0" class="rounded-xl py-3 px-4">
            No Post Here
          </box>
        </error-box>
        <post-feed :list="postsList"></post-feed>
        <snackbar v-model="err.snackbar.show" :closeBtn="true">
          {{ err.snackbar.msg }}
        </snackbar>
      </feed>
      <sidebar></sidebar>
    </div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive, computed } from "vue";
import { useStore } from "vuex";

import { get as getUser } from "../api/user/get";
import { posts as getPosts } from "../api/user/posts";
import { post as userFollow, remove as userUnfollow } from "../api/user/follow";
import { API_Error } from "../models/api_error";
import { API_Posts } from "../models/api_post";
import {
  Public as PUBLIC_USER,
  Unverified as UNVERIFIED_USER,
} from "../models/user_types";
import {
  InvalidCharacters,
  SomethingWentWrong,
} from "../components/form/errors";

import { GEN_User } from "../models/go_structs_gen";

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
import NewPostTextInput from "../components/post/NewPostTextInput.vue";
import Snackbar from "../components/box/Snackbar.vue";
import TextSmall from "../components/text/Small.vue";

export default defineComponent({
  props: {
    userName: {
      type: Array,
      required: true,
    },
  },
  setup(props) {
    const store = useStore();
    const loggedIn = computed(() => store.getters.loggedIn),
      storeUser = computed(() => store.getters.user);
    const data = reactive({
      user: new GEN_User(),
      postsList: new API_Posts(),
      overlay: false,
      page: 1,
      err: {
        user: new API_Error(),
        posts: new API_Error(),
        snackbar: {
          show: false,
          msg: "",
        },
      },
    });

    getUser(props.userName[0])
      .then((r) => {
        data.user = r;
      })
      .catch((e) => {
        data.err.user = e;
      });

    getPosts(props.userName[0], data.page)
      .then((r) => {
        data.postsList = r;
        if (Object.keys(r.posts).length == 0) {
          data.err.posts = "No posts here";
        }
      })
      .catch((e) => {
        data.err.posts = e;
      });

    return {
      ...toRefs(data),
      TwoLine,
      loggedIn,
      storeUser,
      InvalidCharacters,
      UNVERIFIED_USER,
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
    followEditBtnTxt() {
      if (this.canEdit) {
        return "Edit Profile";
      }
      if (this.user.type !== PUBLIC_USER) {
        if (this.user.selfFollow?.following) {
          return "Unfriend";
        } else if (this.user.selfFollow?.requested) {
          return "Requested";
        }
        return "Friend";
      }
      if (this.user.selfFollow?.following) {
        return "Unfollow";
      }
      return "Follow";
    },
  },
  watch: {
    user(newUser) {
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
      } else {
        (() => {
          if (
            this.user.selfFollow.following ||
            this.user.selfFollow.requested
          ) {
            return userUnfollow(this.user.userName).then((r) => {
              if (r.success) {
                this.user.selfFollow.requested =
                  this.user.selfFollow.following = false;
              }
              return r;
            });
          }
          return userFollow(this.user.userName).then((r) => {
            if (r.success) {
              if (this.user.type !== PUBLIC_USER) {
                this.user.selfFollow.requested = true;
              } else {
                this.user.selfFollow.following = true;
              }
            }
            return r;
          });
        })()
          .then((r) => {
            if (!r.success) {
              this.snackbarErr(SomethingWentWrong);
            } else if (this.user.type === UNVERIFIED_USER) {
              this.snackbarErr("This user is unverified");
            }
          })
          .catch(() => {
            this.snackbarErr(SomethingWentWrong);
          });
      }
    },
    snackbarErr(err) {
      this.err.snackbar.msg = "";
      if (err.length > 0) {
        this.err.snackbar.msg = err;
        this.err.snackbar.show = true;
      }
    },
  },
  async beforeRouteUpdate(to) {
    this.user = new GEN_User();
    this.postsList = new API_Posts();
    this.page = 1;

    getUser(to.params.userName[0])
      .then((r) => {
        this.user = r;
      })
      .catch((e) => {
        this.err.user = e;
      });

    getPosts(to.params.userName[0], this.page)
      .then((r) => {
        this.postsList = r;
        if (Object.keys(r.posts).length == 0) {
          this.err.posts = "No posts here";
        }
      })
      .catch((e) => {
        this.err.posts = e;
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
    NewPostTextInput,
    Snackbar,
    TextSmall,
  },
});
</script>
