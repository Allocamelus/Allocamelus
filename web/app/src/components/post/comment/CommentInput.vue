<template>
  <div>
    <input-label for="comment" class="flex" :err="err.comment">
      {{ theType }}ing as
      <user-name :user="storeUser" :isLink="true" :noName="true"></user-name>
    </input-label>
    <text-input
      v-model="comment"
      name="comment"
      :watchModel="true"
      :check="true"
      :required="true"
      :minLen="2"
      :maxLen="4096"
      :placeholder="`Post a ${theType}`"
      :regex="/^[^<>\[\]]*$/"
      :regexMsg="InvalidCharacters"
      @error="err.comment = $event"
    >
      <div class="flex items-center mr-1.5">
        <basic-btn
          class="link p-1"
          :title="`Submit ${theType}`"
          :disabled="commentDisabled"
          @click="submitComment()"
        >
          {{ theType }}
        </basic-btn>
      </div>
    </text-input>
  </div>
</template>

<script>
import { computed, defineComponent, reactive, toRefs } from "vue";
import { useStore } from "../../../store";

import { API_Comment } from "../../../api/post/comment";
import CreateComment from "../../../api/post/comment/create";
import { InvalidCharacters, SomethingWentWrong } from "../../form/errors";
import { UnixTime } from "../../../pkg/time";
import { notNull, RespToError } from "../../../models/responses";

import InputLabel from "../../form/InputLabel.vue";
import TextInput from "../../form/TextInput.vue";
import BasicBtn from "../../button/BasicBtn.vue";
import UserName from "../../user/Name.vue";

export default defineComponent({
  name: "comment-input",
  props: {
    postId: {
      type: String,
      required: true,
    },
    replyTo: {
      type: Number,
      default: 0,
    },
  },
  emits: ["commented"],
  setup() {
    const store = useStore();
    const loggedIn = computed(() => store.getters.loggedIn),
      storeUser = computed(() => store.getters.user);
    const data = reactive({
      comment: "",
      submitted: false,
      err: {
        comment: "",
      },
    });

    return {
      ...toRefs(data),
      loggedIn,
      storeUser,
      InvalidCharacters,
    };
  },
  computed: {
    commentErr() {
      return this.err.comment.length != 0 || this.comment.length < 2;
    },
    commentDisabled() {
      return this.commentErr || this.submitted;
    },
    theType() {
      return this.replyTo > 0 ? "Reply" : "Comment";
    },
  },
  methods: {
    submitComment() {
      if (!this.commentErr && this.loggedIn) {
        this.submitted = true;
        // Start time of query
        let start = UnixTime();
        CreateComment(this.postId, this.replyTo, this.comment)
          .then((r) => {
            if (r.success) {
              // End time of query + processing
              let end = UnixTime();
              this.$emit(
                "commented",
                new API_Comment({
                  id: r.id,
                  userId: this.storeUser.id,
                  postId: Number(this.postId).valueOf(),
                  parentId: this.replyTo,
                  created: (start + end) / 2, // Guess creation time with query times
                  updated: 0,
                  content: this.comment,
                  replies: 0,
                  depth: 0,
                  children: [],
                })
              );
              // Clear comment text
              this.comment = "";
            }
            // Handle error (if any)
            this.onPostErr(r.error);
          })
          // Handle error
          .catch((e) => {
            this.onPostErr(e);
          });
      }
    },
    onPostErr(e) {
      this.submitted = false;
      // Check if error actually exist
      if (notNull(e)) {
        let errText = RespToError(e);
        if (errText.length > 0) {
          this.err.comment = errText;
        } else {
          this.err.comment = SomethingWentWrong;
        }
      }
    },
  },
  components: {
    InputLabel,
    TextInput,
    BasicBtn,
    UserName,
  },
});
</script>
