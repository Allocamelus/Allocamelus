<template>
  <div>
    <feed v-for="(commentId, index) in list.order" :key="index" class="flex-c">
      <div
        v-if="list.comment(commentId).replyToId == replyToId"
        class="pl-3 py-2"
      >
        <comment-box
          :comment="list.comment(commentId)"
          :user="list.user(list.comment(commentId).userId)"
        ></comment-box>
        <div class="flex flex-grow">
          <div class="pl-3 w-8 flex align-items-center">
            <div class="w-[1px] bg-primary-50"></div>
          </div>
          <div class="">
            <comment-feed
              v-if="list.comment(commentId).replies > 0"
              :list="list"
              :depth="depth + 1"
              :replyToId="commentId"
            ></comment-feed>
          </div>
        </div>
      </div>
    </feed>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { API_Comments } from "../../api/post/comments/get";

import Box from "../box/Box.vue";
import Feed from "../Feed.vue";
import CommentBox from "./CommentBox.vue";

export default defineComponent({
  name: "comment-feed",
  props: {
    list: {
      type: API_Comments,
      default: new API_Comments(),
    },
    depth: {
      type: Number,
      default: 0,
    },
    replyToId: {
      type: Number,
      default: 0,
    },
  },
  components: { Box, CommentBox, Feed },
});
</script>