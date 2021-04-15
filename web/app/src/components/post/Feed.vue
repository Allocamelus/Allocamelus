<template>
  <box
    v-for="(postId, index) in list.order"
    :key="index"
    class="mb-3 rounded-xl"
  >
    <post-box
      :post="list.post(postId)"
      :user="list.user(list.post(postId).userId)"
      :isLink="true"
    ></post-box>
  </box>
</template>

<script>
import { defineComponent } from "vue";

import { API_Posts } from "../../models/api_posts";

import Box from "../box/Box.vue";
import PostBox from "./Box.vue";

export default defineComponent({
  props: {
    list: API_Posts,
  },
  components: { Box, PostBox },
});
</script>