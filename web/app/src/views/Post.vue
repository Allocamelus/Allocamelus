<template>
  <div>
    <p>{{ id }}</p>
    <div v-html="post.content"></div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { useStore } from "vuex";
import { get as getPost} from "../api/post/get";
import { Post } from "@/models/post_gen";

export default defineComponent({
  props: {
    id: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const store = useStore();
    const data = reactive({
      post: new Post(),
    });

    getPost(props.id).then((r) => {
      data.post = r;
    });
    return {
      ...toRefs(data),
    };
  },
  async beforeRouteUpdate(to, from) {
    this.post = new Post();

    getPost(to.params.id).then((r) => {
      this.post = r;
    });
  },
});
</script>

<style>
</style>