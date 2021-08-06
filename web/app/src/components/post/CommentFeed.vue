<template>
  <div>
    <feed class="flex-col-reverse">
      <div v-for="(commentId, index) in list.order" :key="index" class="py-2">
        <comment-tree
          :comment="list.comment(commentId)"
          :userList="list"
        ></comment-tree>
      </div>
    </feed>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { API_Comments } from "../../api/post/comments/get";

import Feed from "../Feed.vue";
import CommentTree from "./CommentTree.vue";

export default defineComponent({
  name: "comment-feed",
  props: {
    list: {
      type: API_Comments,
      default: new API_Comments(),
    },
  },
  components: { Feed, CommentTree },
});
</script>