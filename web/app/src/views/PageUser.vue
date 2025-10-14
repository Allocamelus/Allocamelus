<template>
  <div class="container py-5">
    <error-box :error="err.user" class="mb-3 p-3.5">
      <div class="flex flex-col justify-between xs:flex-row">
        <div>
          <div class="flex min-w-0 items-center">
            <component
              :is="canEdit ? 'change-avatar' : 'user-avatar'"
              :user="canEdit ? storeUser : user"
              class="h-14 w-14 flex-shrink-0 xs:h-16 xs:w-16 md:h-20 md:w-20"
            ></component>
            <user-name
              class="ml-3"
              :user="canEdit ? storeUser : user"
              :two-line="true"
              :is-link="false"
            ></user-name>
            <text-small
              v-if="user.type === UNVERIFIED_USER"
              class="ml-1 flex-none font-normal"
            >
              [[Unverified]]
            </text-small>
          </div>
          <div class="mt-3 text-lg">
            {{ canEdit ? storeUser.bio : user.bio }}
          </div>
        </div>
        <div
          class="mt-3 flex flex-shrink-0 items-start justify-end xs:mt-0 xs:ml-3"
        >
          <basic-btn
            class="whitespace-nowrap px-3 py-2"
            :class="buttonStyle.secondaryBorderInvert"
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
          <box v-if="postsList.total() == 0" class="rounded-xl px-4 py-3">
            No Post Here
          </box>
        </error-box>
        <post-feed :list="postsList"></post-feed>
        <snackbar v-model="err.snackbar.show" :close-btn="true">
          {{ err.snackbar.msg }}
        </snackbar>
      </feed>
      <sidebar></sidebar>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, toRefs, reactive, computed } from "vue";
import { useSessionStore } from "@/store/session";

import { get as getUser } from "@/api/user/get";
import { posts as getPosts } from "@/api/user/posts";
import { post as userFollow, remove as userUnfollow } from "@/api/user/follow";
import { API_Error, API_Success_Error } from "@/models/api_error";
import { API_Posts } from "@/models/api_post";
import {
  Public as PUBLIC_USER,
  Unverified as UNVERIFIED_USER,
  User,
} from "@/models/user";
import {
  InvalidCharacters,
  SomethingWentWrong,
} from "@/components/form/errors";

import UserName from "@/components/user/Name.vue";
import ErrorBox from "@/components/box/Error.vue";
import PostFeed from "@/components/post/Feed.vue";
import Feed from "@/components/Feed.vue";
import Sidebar from "@/components/Sidebar.vue";
import Box from "@/components/box/Box.vue";
import UserAvatar from "@/components/user/Avatar.vue";
import BasicBtn from "@/components/button/BasicBtn.vue";
import Overlay from "@/components/overlay/Overlay.vue";
import ChangeAvatar from "@/components/user/ChangeAvatar.vue";
import EditOverlay from "@/components/user/EditOverlay.vue";
import SignUpOverlay from "@/components/overlay/SignUpOverlay.vue";
import NewPostTextInput from "@/components/post/NewPostTextInput.vue";
import Snackbar from "@/components/box/Snackbar.vue";
import TextSmall from "@/components/text/Small.vue";

export default defineComponent({
  props: {
    userName: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const session = useSessionStore();
    const data = reactive({
      user: new User(),
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

    getUser(props.userName)
      .then((r) => {
        data.user = r;
      })
      .catch((e) => {
        data.err.user = e;
      });

    getPosts(props.userName, data.page)
      .then((r) => {
        data.postsList = r;
        if (Object.keys(r.posts).length == 0) {
          data.err.posts.error = "No posts here";
        }
      })
      .catch((e) => {
        data.err.posts = e;
      });

    return {
      ...toRefs(data),
      loggedIn: computed(() => session.loggedIn),
      storeUser: computed(() => session.user),
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
        (async () => {
          let r: API_Success_Error;
          if (
            this.user.selfFollow.following ||
            this.user.selfFollow.requested
          ) {
            r = await userUnfollow(this.user.userName);
            if (r.success) {
              this.user.selfFollow.requested =
                this.user.selfFollow.following = false;
            }
          } else {
            r = await userFollow(this.user.userName);
            if (r.success) {
              if (this.user.type !== PUBLIC_USER) {
                this.user.selfFollow.requested = true;
              } else {
                this.user.selfFollow.following = true;
              }
            }
          }
          return r;
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
    snackbarErr(err: string) {
      this.err.snackbar.msg = "";
      if (err.length > 0) {
        this.err.snackbar.msg = err;
        this.err.snackbar.show = true;
      }
    },
  },
  async beforeRouteUpdate(to) {
    this.user = new User();
    this.postsList = new API_Posts();
    this.page = 1;

    const userName = Array.isArray(to.params.userName)
      ? to.params.userName[0]
      : to.params.userName;

    getUser(userName)
      .then((r) => {
        this.user = r;
      })
      .catch((e) => {
        this.err.user = e;
      });

    getPosts(userName, this.page)
      .then((r) => {
        this.postsList = r;
        if (Object.keys(r.posts).length == 0) {
          this.err.posts.error = "No posts here";
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
    ChangeAvatar,
    EditOverlay,
    SignUpOverlay,
    NewPostTextInput,
    Snackbar,
    TextSmall,
  },
});
</script>

<style
  src="@/scss/modules/button.modules.scss"
  lang="scss"
  module="buttonStyle"
  scoped
></style>
