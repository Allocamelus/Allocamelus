<template>
  <div class="container flex py-5">
    <feed>
      <div v-if="err.length > 0" v-html="err"></div>
      <new-post-text-input v-if="loggedIn"></new-post-text-input>
      <post-feed :list="list"></post-feed>
    </feed>
    <sidebar></sidebar>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive, computed } from "vue";
import { useStore } from "vuex";
import { get as getPosts } from "../api/posts/get";
import { API_Posts } from "../models/api_posts";
import Box from "../components/box/Box.vue";
import PostFeed from "../components/post/Feed.vue";
import Feed from "../components/Feed.vue";
import Sidebar from "../components/Sidebar.vue";
import NewPostTextInput from "../components/post/NewPostTextInput.vue";

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
    this.list = new API_Posts();

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
    PostFeed,
    Feed,
    Sidebar,
    NewPostTextInput,
  },
});
</script>

<style lang="scss" scoped>
</style>