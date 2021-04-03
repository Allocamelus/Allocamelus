<template>
  <div class="container py-5">
    <error-box :error="err.user" class="py-3 px-4 mb-3">
      <user-name :user="user" :displayType="TwoLine"></user-name>
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
import { defineComponent, toRefs, reactive } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";

import { get as getUser } from "../api/user/get";
import { posts as getPosts } from "../api/user/posts";
import { API_Error } from "../models/api_error";
import { API_Posts } from "../models/api_posts";
import { Html404Func, HtmlSomthingWentWrong } from "../components/htmlErrors";

import { User } from "../models/user_gen";
import ApiResp from "../models/responses";

import UserName, { TwoLine } from "../components/user/Name.vue";
import ErrorBox from "../components/box/Error.vue";
import PostFeed from "../components/post/Feed.vue";
import Feed from "../components/Feed.vue";
import Sidebar from "../components/Sidebar.vue";
import Box from "../components/box/Box.vue";

function userErrors(api_error, path) {
  if (api_error instanceof API_Error) {
    switch (api_error.error) {
      case ApiResp.Shared.NotFound:
        return Html404Func(path);
    }
  }
  return HtmlSomthingWentWrong;
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
    const data = reactive({
      user: new User(),
      postsList: new API_Posts(),
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
      })
      .catch((e) => {
        data.err.posts = userErrors(e, route.currentRoute.value.fullPath);
      });

    return {
      ...toRefs(data),
      TwoLine,
    };
  },
  watch: {
    user(newUser, old) {
      document.title = `${newUser.name} (@${newUser.userName}) - ${
        import.meta.env.VITE_SITE_NAME
      }`;
    },
  },
  async beforeRouteUpdate(to, from) {
    this.user = new User();
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
  },
});
</script>
