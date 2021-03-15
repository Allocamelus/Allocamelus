<template>
  <div>
    <p>{{ id }}</p>
    <div v-html="post.content"></div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { onBeforeRouteUpdate } from "vue-router";
import { getPost } from "@/api/getPost";
import { Post } from "@/models/post_gen";
export default defineComponent({
  props: {
    id: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const data = reactive({
      post: new Post(),
    });
    onBeforeRouteUpdate(async (to, from, next) => {
      data.post = await getPost(to.params.id);
      next();
    });
    getPost(props.id).then((r) => {
      data.post = r;
    });
    return {
      ...toRefs(data),
    };
  },
  /*
  async beforeRouteUpdate(to, from) {
    this.post = Post;
    try {
      this.post = await getPost(to.params.id);
    } catch (error) {
      console.log(error);
    }
  },
  */
});
</script>

<style>
</style>