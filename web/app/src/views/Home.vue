<template>
  <div class="container flex py-5">
    <feed>
      <div v-if="err.length > 0" v-html="err"></div>
      <new-post-text-input v-if="loggedIn"></new-post-text-input>
      <box v-if="list.total() == 0" class="rounded-xl py-3 px-4">
        Follow someone to see their post here
      </box>
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
import PostFeed from "../components/post/Feed.vue";
import Feed from "../components/Feed.vue";
import Sidebar from "../components/Sidebar.vue";
import NewPostTextInput from "../components/post/NewPostTextInput.vue";
import Box from "../components/box/Box.vue";

export default defineComponent({
  setup() {
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
  async beforeRouteUpdate(to, _from) {
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
    PostFeed,
    Feed,
    Sidebar,
    NewPostTextInput,
    Box,
  },
});
</script>

<style lang="scss" scoped>
</style>