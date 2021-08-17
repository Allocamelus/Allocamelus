<template>
  <div>
    <comment-input
      :postId="postId"
      :replyTo="0"
      @commented="newComment($event)"
    ></comment-input>
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
import { computed, defineComponent, toRefs } from "vue";
import { useStore } from "vuex";

import { API_Comments } from "../../api/post/comments/get";

import Feed from "../Feed.vue";
import CommentTree from "./CommentTree.vue";
import CommentInput from "./CommentInput.vue";

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
  setup() {
    const store = useStore();
    const loggedIn = computed(() => store.getters.loggedIn),
      storeUser = computed(() => store.getters.user);
    return {
      loggedIn,
      storeUser,
    };
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
  },
});
</script>