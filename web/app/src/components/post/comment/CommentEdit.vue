<template>
  <div>
    <text-input
      v-model="comment.content"
      name="comment"
      :watch-model="true"
      :check="true"
      :required="true"
      :min-len="2"
      :max-len="4096"
      :placeholder="`Editing...`"
      :regex="/^[^<>\[\]]*$/"
      :regex-msg="InvalidCharacters"
      @error="err.comment = $event"
    >
      <div class="mr-1.5 flex items-center">
        <basic-btn
          class="mr-1.5 p-1 text-neutral-700 dark:text-neutral-300"
          title="Close"
          @click="close()"
        >
          Cancel
        </basic-btn>
        <basic-btn
          class="link p-1"
          title="Update Comment"
          :disabled="commentDisabled"
          @click="updateComment()"
        >
          Update
        </basic-btn>
      </div>
    </text-input>
    <input-label for="comment" class="flex" :err="err.comment"> </input-label>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, reactive, toRefs } from "vue";
import { useCommentStore } from "@/store/comments";
import { useSessionStore } from "@/store/session";

import { API_Comment } from "@/api/post/comment";
import UpdateComment from "@/api/post/comment/update";
import { InvalidCharacters, SomethingWentWrong } from "../../form/errors";
import { notNull, RespToError } from "@/models/responses";
import { UnixTime } from "@/pkg/time";

import InputLabel from "../../form/InputLabel.vue";
import TextInput from "../../form/TextInput.vue";
import BasicBtn from "../../button/BasicBtn.vue";

export default defineComponent({
  name: "CommentInput",
  props: {
    postId: {
      type: Number,
      required: true,
    },
    commentId: {
      type: Number,
      required: true,
    },
  },
  emits: ["edited", "close"],
  setup(props) {
    const commentStore = useCommentStore(props.postId);
    const session = useSessionStore();

    const data = reactive({
      // Don't use computed value so it can be edited
      comment: commentStore.comment(props.commentId) || new API_Comment(),
      submitted: false,
      err: {
        comment: "",
      },
    });

    return {
      ...toRefs(data),
      loggedIn: computed(() => session.loggedIn),
      storeUser: computed(() => session.user),
      InvalidCharacters,
    };
  },
  computed: {
    commentErr() {
      return this.err.comment.length != 0 || this.comment.content.length < 2;
    },
    commentDisabled() {
      return this.commentErr || this.submitted;
    },
  },
  methods: {
    close() {
      this.$emit("close");
    },
    updateComment() {
      if (!this.commentErr && this.loggedIn) {
        this.submitted = true;

        // Start time of query
        let start = UnixTime();
        UpdateComment(this.postId, this.commentId, this.comment.content)
          .then((r) => {
            // End time of query + processing
            let end = UnixTime();
            if (r.success) {
              this.comment.updated = (start + end) / 2; // Guess update time with query times
              this.$emit("edited", new API_Comment(this.comment));
              this.close();
            }
            // Handle error (if any)
            this.onPostErr(r.error);
          })
          // Handle error
          .catch((e) => {
            this.onPostErr(String(e));
          });
      }
    },
    onPostErr(e: string | undefined) {
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
  },
});
</script>
