<template>
  <div class="container container--flex">
    <div class="container__feed">
      <div v-if="err.length > 0" v-html="err"></div>
      <box
        v-for="post in list.posts"
        :key="post.id"
        class="box--link box--auto-mb"
        data-pointer
        @click="$router.push(`/post/${post.id}`)"
      >
        <div>{{ post.id }}</div>
        <div v-html="post.content"></div>
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
import { getPosts } from "../api/getPosts";
import { List } from "../models/post_gen";
import Box from "../components/Box.vue";
export default defineComponent({
  props: {
    page: {
      type: Number,
      default: 1,
    },
  },
  setup(props) {
    const store = useStore();
    const usedSession = () => store.dispatch("usedSession");
    const data = reactive({
      list: new List(),
      // TODO Better Errors
      err: "",
    });
    getPosts(props.page)
      .then((r) => {
        data.list = r;
        usedSession();
      })
      .catch((e) => {
        data.err = String(e);
      });
    return {
      ...toRefs(data),
      usedSession,
      loggedIn: computed(() => store.getters.loggedIn),
    };
  },
  async beforeRouteUpdate(to, from) {
    this.list = new List();

    getPosts(to.params.page)
      .then((r) => {
        this.list = r;
      })
      .catch((e) => {
        this.err = String(e);
      });
    this.usedSession();
  },
  components: {
    Box,
  },
});
</script>

<style lang="scss" scoped>
</style>