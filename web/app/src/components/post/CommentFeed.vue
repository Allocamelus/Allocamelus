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
    </box>
    <feed class="flex-col-reverse">
      <div v-for="(commentId, index) in list.order" :key="index" class="py-2">
        <comment-tree
          :commentId="commentId"
          :postId="postId"
        ></comment-tree>
      </div>
    </feed>
  </div>
</template>

<script>
import { computed, defineComponent, toRefs } from "vue";
import { useStore } from "vuex";
import CommentsStore from "../../store/module/comments";

import { API_Comments } from "../../api/post/comments/get";

import Feed from "../Feed.vue";
import CommentTree from "./CommentTree.vue";
import CommentInput from "./comment/CommentInput.vue";
import Box from "../box/Box.vue";

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
    const storeName = `p${props.postId}-comments`;
    const loggedIn = computed(() => store.getters.loggedIn),
      storeUser = computed(() => store.getters.user),
      updateComments = (c) => store.commit(`${storeName}/update`, c);

    store.registerModule(storeName, CommentsStore);

    return {
      storeName,
      loggedIn,
      storeUser,
      updateComments,
    };
  },
  watch: {
    list(newValue) {
      this.updateComments(newValue);
    },
  },
  unmounted() {
    console.log("unmounted");
    this.$store.unregisterModule(this.storeName);
  },
  methods: {
    newComment(c) {
      this.list.appendComment(c);
      this.list.appendUser(this.storeUser);
    },
  },
  components: {
    Feed,
    CommentTree,
    CommentInput,
    Box,
  },
});
</script>