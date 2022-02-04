<template>
  <div class="container py-5">
    <div class="max-w-full sm:mx-2 md:mx-4 lg:mx-8 xl:mx-12">
      <error-box :error="err">
        <post-box
          :post="apiPost.post"
          :user="apiPost.user"
          :dynamicContent="true"
        ></post-box>
      </error-box>
      <div></div>
      <comment-feed :list="comments" :postId="id"></comment-feed>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, toRefs, reactive } from "vue";
import { get as getPost } from "../api/post/get";
import { get as getComments, API_Comments } from "../api/post/comments/get";
import { API_Post } from "../models/api_post";
import { API_Error } from "../models/api_error";
import ErrorBox from "../components/box/Error.vue";
import sanitize from "../pkg/sanitize";
import PostBox from "../components/post/Box.vue";
import CommentFeed from "../components/post/CommentFeed.vue";

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
      comments: new API_Comments(),
      err: new API_Error(),
      commentErr: new API_Error(),
    });

    getPost(props.id)
      .then((r) => {
        data.apiPost = r;
        // Get Comments once post is fetched
        (async () => {
          getComments(props.id)
            .then((r) => {
              data.comments = r;
            })
            .catch((e) => {
              data.err = e;
            });
        })();
      })
      .catch((e) => {
        data.err = e;
      });

    return {
      ...toRefs(data),
    };
  },
  watch: {
    apiPost(newPost) {
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
  async beforeRouteUpdate(to) {
    this.apiPost = new API_Post();

    getPost(to.params.id)
      .then((r) => {
        this.apiPost = r;
        // Get Comments once post is fetched
        (async () => {
          getComments(to.params.id)
            .then((r) => {
              this.comments = r;
            })
            .catch((e) => {
              this.err = e;
            });
        })();
      })
      .catch((e) => {
        this.err = e;
      });
  },
  components: {
    ErrorBox,
    PostBox,
    CommentFeed,
  },
});
</script>

<style></style>
