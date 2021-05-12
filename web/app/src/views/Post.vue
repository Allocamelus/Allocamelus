<template>
  <div class="container py-5">
    <error-box
      :error="err"
      class="max-w-full sm:mx-2 md:mx-4 lg:mx-8 xl:mx-12"
    >
      <post-box
        :post="apiPost.post"
        :user="apiPost.user"
        :dynamicContent="true"
      ></post-box>
    </error-box>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { get as getPost } from "../api/post/get";
import { API_Post } from "../models/api_post";
import { API_Error } from "../models/api_error";
import ErrorBox from "../components/box/Error.vue";
import sanitize from "../pkg/sanitize";
import PostBox from "../components/post/Box.vue";

export default defineComponent({
  props: {
    id: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const data = reactive({
      apiPost: new API_Post(),
      err: new API_Error(),
    });

    getPost(props.id)
      .then((r) => {
        data.apiPost = r;
      })
      .catch((e) => {
        data.err = e;
      });

    return {
      ...toRefs(data),
    };
  },
  watch: {
    apiPost(newPost, _old) {
      var sanitizedContent = sanitize(newPost.post.content).trim(),
        truncatedContent = sanitizedContent.substring(0, 256).trim();

      if (sanitizedContent.length > truncatedContent.length) {
        truncatedContent += "...";
      }

      document.title =
        `${newPost.user.name} (@${newPost.user.userName}) - ` +
        `"${truncatedContent}" - ` +
        `${import.meta.env.VITE_SITE_NAME}`;
    },
  },
  async beforeRouteUpdate(to, _from) {
    this.apiPost = new API_Post();

    getPost(to.params.id)
      .then((r) => {
        this.apiPost = r;
      })
      .catch((e) => {
        this.err = e;
      });
  },
  components: {
    ErrorBox,
    PostBox,
  },
});
</script>

<style>
</style>