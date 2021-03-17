<template>
  <div>
    <p>{{ id }}</p>
    <div v-html="post.content"></div>
    <router-link to="/post/10">10</router-link>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { useStore } from "vuex";
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
    const store = useStore();
    const usedSession = () => store.dispatch("usedSession");
    const data = reactive({
      post: new Post(),
    });

    getPost(props.id).then((r) => {
      data.post = r;
      usedSession();
    });
    return {
      ...toRefs(data),
      usedSession,
    };
  },
  async beforeRouteUpdate(to, from) {
    this.post = new Post();

    getPost(to.params.id).then((r) => {
      this.post = r;
      this.usedSession();
    });
  },
});
</script>

<style>
</style>