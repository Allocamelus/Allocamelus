<template>
  <div>
    <box class="mt-5 mb-3 px-4 py-3 rounded-xl">
      <div v-if="loggedIn">
        <input-label for="comment" class="flex" :err="err.comment">
          Commenting as
          <user-name :user="storeUser" :displayType="usernameType"></user-name>
        </input-label>
        <text-input
          v-model="comment"
          name="comment"
          :watchModel="true"
          :check="true"
          :required="true"
          :minLen="2"
          :maxLen="4096"
          placeholder="Post a Comment"
          :regex="/^[^<>\[\]]*$/"
          :regexMsg="InvalidCharacters"
          @error="err.comment = $event"
        >
          <div class="flex items-center mr-1.5">
            <basic-btn
              class="link p-1"
              title="Submit Comment"
              :disabled="commentErr"
              @click="submitComment()"
            >
              Comment
            </basic-btn>
          </div>
        </text-input>
      </div>
    </box>
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
import { computed, defineComponent, reactive, toRefs } from "vue";
import { useStore } from "vuex";

import { API_Comment } from "../../api/post/comment";
import { API_Comments } from "../../api/post/comments/get";
import CreateComment from "../../api/post/comment/create";
import { InvalidCharacters, SomethingWentWrong } from "../form/errors";

import { UnixTime } from "../../pkg/time";

import Feed from "../Feed.vue";
import CommentTree from "./CommentTree.vue";
import InputLabel from "../form/InputLabel.vue";
import TextInput from "../form/TextInput.vue";
import BasicBtn from "../button/BasicBtn.vue";
import UserName, { OneLineLink, NoName } from "../user/Name.vue";
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
  setup() {
    const store = useStore();
    const loggedIn = computed(() => store.getters.loggedIn),
      storeUser = computed(() => store.getters.user);
    const data = reactive({
      comment: "",
      err: {
        comment: "",
      },
    });
    const usernameType = OneLineLink | NoName;

    return {
      ...toRefs(data),
      loggedIn,
      storeUser,
      InvalidCharacters,
      usernameType,
    };
  },
  computed: {
    commentErr() {
      return this.err.comment.length != 0 || this.comment.length < 2;
    },
  },
  methods: {
    submitComment() {
      if (!this.commentErr) {
        var start = UnixTime();
        CreateComment(this.postId, 0, this.comment)
          .then((r) => {
            if (r.success) {
              var end = UnixTime();
              this.list.appendComment(
                new API_Comment({
                  id: r.id,
                  userId: this.storeUser.id,
                  postId: Number(this.postId).valueOf(),
                  parentId: 0,
                  created: (start + end) / 2, // Guess creation time with query times
                  updated: 0,
                  content: this.comment,
                  replies: 0,
                  depth: 0,
                  children: [],
                })
              );
              this.list.appendUser(this.storeUser);
              console.log(this.list);
              this.comment = "";
            }
          });
      }
    },
  },
  components: {
    Feed,
    CommentTree,
    InputLabel,
    TextInput,
    BasicBtn,
    UserName,
    Box,
  },
});
</script>