<template>
  <div>
    <box class="mt-5 mb-3 rounded-xl px-4 py-3">
      <div v-if="loggedIn">
        <comment-input
          :post-id="postId"
          :reply-to="0"
          @commented="newComment($event)"
        ></comment-input>
      </div>
      <div v-else>
        <text-input
          placeholder="Post a Comment"
          :readonly="true"
          @click="
            $router.push({ path: '/login', query: { r: `/post/${postId}` } })
          "
        >
        </text-input>
      </div>
    </box>
    <feed class="flex-col-reverse">
      <div v-for="(commentId, index) in list.order" :key="index" class="py-2">
        <comment-tree :comment-id="commentId" :post-id="postId"></comment-tree>
      </div>
    </feed>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent } from "vue";
import { useCommentStore } from "@/store/comments";
import { useSessionStore } from "@/store/session";

import { API_Comments } from "@/api/post/comments/get";
import { API_Comment } from "@/api/post/comment";
import { User } from "@/models/user";

import Feed from "../Feed.vue";
import CommentTree from "./CommentTree.vue";
import CommentInput from "./comment/CommentInput.vue";
import Box from "../box/Box.vue";
import TextInput from "../form/TextInput.vue";

export default defineComponent({
  name: "CommentFeed",
  props: {
    list: {
      type: API_Comments,
      default: new API_Comments(),
    },
    postId: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const commentStore = useCommentStore(props.postId);
    const session = useSessionStore();

    return {
      loggedIn: computed(() => session.loggedIn),
      storeUser: computed(() => session.user),
      addComment: (c: API_Comment, isNew: boolean) =>
        commentStore.addComment(c, isNew),
      addUser: (u: User) => commentStore.addUser(u),
      populate: (nv: Partial<API_Comments>) => {
        commentStore.populate(nv);
      },
      disposeComment: () => {
        commentStore.$dispose();
      },
    };
  },
  watch: {
    list(newValue) {
      this.populate(newValue);
    },
  },
  beforeUnmount() {
    this.disposeComment();
  },
  methods: {
    newComment(c: API_Comment) {
      this.addComment(c, true);
      this.addUser(this.storeUser);
    },
  },
  components: {
    Feed,
    CommentTree,
    CommentInput,
    Box,
    TextInput,
  },
});
</script>
