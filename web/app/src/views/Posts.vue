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
        <div v-if="$store.getters.loggedIn">
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
import { defineComponent, toRefs, reactive } from "vue";
import { onBeforeRouteUpdate } from "vue-router";
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
    const data = reactive({
      list: new List(),
      // TODO Better Errors
      err: "",
    });
    onBeforeRouteUpdate(async (to, from, next) => {
      await getPosts(to.params.page)
        .then((r) => {
          console.log(r);
          data.list = r;
        })
        .catch((e) => {
          data.err = String(e);
        });
      next();
    });
    getPosts(props.page)
      .then((r) => {
        console.log(r);
        data.list = r;
      })
      .catch((e) => {
        data.err = String(e);
      });
    return {
      ...toRefs(data),
    };
  },
  components: {
    Box,
  },
});
</script>

<style lang="scss" scoped>
</style>