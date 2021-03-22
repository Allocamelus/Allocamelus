<template>
  <div class="container container--flex">
    <div class="container__feed">
      <div v-if="err.length > 0" v-html="err"></div>
      <box
        v-for="post in list.posts"
        :key="post.id"
        class="box--link box--auto-mb"
      >
        <user-name :user="list.user(post.userId)"></user-name>
        <div
          @click="$router.push(`/post/${post.id}`)"
          data-pointer
          v-html="post.content"
        ></div>
      </box>
    </div>
    <div class="container__sidebar">
      <box class="pa-4 box--auto-mb">
        <div v-if="loggedIn">
          <router-link class="link" to="/post/new">New Post</router-link>
        </div>
        <div v-else>
          <router-link class="link" to="/login">Login</router-link>
        </div>
      </box>
    </div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive, computed } from "vue";
import { useStore } from "vuex";
import { get as getPosts } from "../api/posts/get";
import { API_Posts } from "../models/api_posts";
import Box from "../components/Box.vue";
import UserName from "../components/user/Name.vue";

export default defineComponent({
  setup(props) {
    const store = useStore();
    const data = reactive({
      list: new API_Posts(),
      page: 1,
      // TODO Better Errors
      err: "",
    });
    getPosts(data.page)
      .then((r) => {
        data.list = r;
      })
      .catch((e) => {
        data.err = String(e);
      });
    document.title = `Home - ${import.meta.env.VITE_SITE_NAME}`;
    return {
      ...toRefs(data),
      loggedIn: computed(() => store.getters.loggedIn),
    };
  },
  async beforeRouteUpdate(to, from) {
    this.list = new List();

    getPosts(this.page)
      .then((r) => {
        this.list = r;
      })
      .catch((e) => {
        this.err = String(e);
      });
  },
  components: {
    Box,
    UserName,
  },
});
</script>

<style lang="scss" scoped>
</style>