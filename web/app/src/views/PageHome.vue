<template>
  <div class="container flex py-5">
    <feed>
      <error v-if="err.error?.length > 0" class="mb-3" :error="err"></error>
      <new-post-text-input></new-post-text-input>
      <box v-if="list.total() == 0" class="rounded-xl px-4 py-3">
        Follow someone to see their post here
      </box>
      <post-feed :list="list"></post-feed>
    </feed>
    <sidebar></sidebar>
  </div>
</template>

<script lang="ts">
import { defineComponent, toRefs, reactive } from "vue";

import { get as getPosts } from "../api/posts/get";
import { API_Posts } from "../models/api_post";
import PostFeed from "../components/post/Feed.vue";
import Feed from "../components/Feed.vue";
import Sidebar from "../components/Sidebar.vue";
import NewPostTextInput from "../components/post/NewPostTextInput.vue";
import Box from "../components/box/Box.vue";
import Error from "@/components/box/Error.vue";
import { API_Error } from "@/models/api_error";

export default defineComponent({
  setup() {
    const data = reactive({
      list: new API_Posts(),
      page: 1,
      // TODO Better Errors
      err: new API_Error(),
    });
    getPosts(data.page)
      .then((r) => {
        if ("error" in r) {
          data.err = r;
        } else {
          data.list = r;
        }
      })
      .catch((e) => {
        data.err = new API_Error({ error: String(e) });
      });

    document.title = `Home - ${import.meta.env.VITE_SITE_NAME}`;

    return {
      ...toRefs(data),
    };
  },
  async beforeRouteUpdate() {
    this.list = new API_Posts();

    getPosts(this.page)
      .then((r) => {
        if ("error" in r) {
          this.err = r;
        } else {
          this.list = r;
        }
      })
      .catch((e) => {
        this.err = new API_Error({ error: String(e) });
      });
  },
  components: {
    PostFeed,
    Feed,
    Sidebar,
    NewPostTextInput,
    Box,
    Error,
  },
});
</script>
