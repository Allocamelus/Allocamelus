<template>
  <div>
    <box class="mt-5 mb-3 px-4 py-3 rounded-xl">
      <div v-if="loggedIn">
        <comment-input
          :postId="postId"
          :replyTo="0"
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
        <comment-tree :commentId="commentId" :postId="postId"></comment-tree>
      </div>
    </feed>
  </div>
</template>

<script>
import { computed, defineComponent, toRefs } from "vue";
import { useStore } from "../../store";
import { useSessionStore } from "../../store2/session";
import CommentsStore from "../../store/module/comments";

import { API_Comments } from "../../api/post/comments/get";

import Feed from "../Feed.vue";
import CommentTree from "./CommentTree.vue";
import CommentInput from "./comment/CommentInput.vue";
import Box from "../box/Box.vue";
import TextInput from "../form/TextInput.vue";

export default defineComponent({
  name: "comment-feed",
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
    const store = useStore();
    const session = useSessionStore();
    const storeName = `p${props.postId}-comments`;
    const addComment = (c) => store.commit(`${storeName}/addComment`, c),
      addUser = (c) => store.commit(`${storeName}/addUser`, c),
      populateComments = (c) => store.commit(`${storeName}/populate`, c);

    store.registerModule(storeName, CommentsStore);

    return {
      storeName,
      loggedIn: computed(() => session.loggedIn),
      storeUser: computed(() => session.user),
      addComment,
      addUser,
      populateComments,
    };
  },
  watch: {
    list(newValue) {
      this.populateComments(newValue);
    },
  },
  beforeUnmount() {
    this.$store.unregisterModule(this.storeName);
  },
  methods: {
    newComment(c) {
      this.addComment({
        // AddCommentParams
        newComment: true,
        comment: c,
      });
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
