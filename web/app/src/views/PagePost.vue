<template>
  <div class="container py-5">
    <div class="max-w-full sm:mx-2 md:mx-4 lg:mx-8 xl:mx-12">
      <error-box :error="err">
        <post-box :post="apiPost.post" :user="apiPost.user" :dynamic-content="true"></post-box>
      </error-box>
      <div></div>
      <comment-feed :list="comments" :post-id="id"></comment-feed>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, toRefs, reactive } from "vue";
import { get as getPost } from "@/api/post/get";
import { get as getComments, API_Comments } from "@/api/post/comments/get";
import { API_Post } from "@/models/api_post";
import { API_Error } from "@/models/api_error";
import ErrorBox from "@/components/box/Error.vue";
import { textContent } from "@/pkg/sanitize";
import PostBox from "@/components/post/Box.vue";
import CommentFeed from "@/components/post/CommentFeed.vue";

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
      async getPost(id: string) {
        getPost(id)
          .then(async (r) => {
            this.apiPost = r;
            // Get Comments once post is fetched
            let [comments, err] = await getComments(id)
            if (err != null) {
              this.err = err
              return
            }
            this.comments = comments
          })
          .catch((e) => {
            this.err = e;
          });
      }
    });

    data.getPost(props.id);

    return {
      ...toRefs(data),
    };
  },
  watch: {
    apiPost(newPost) {
      let text = textContent(newPost.post.content).trim(),
        truncatedContent = text.substring(0, 256).trim();

      if (text.length > truncatedContent.length) {
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

    const id = Array.isArray(to.params.id)
      ? to.params.id[0]
      : to.params.id;

    await this.getPost(id);
  },
  components: {
    ErrorBox,
    PostBox,
    CommentFeed,
  },
});
</script>

<style></style>
