<template>
  <div>
    <text-input
      v-model="comment.content"
      name="comment"
      :watchModel="true"
      :check="true"
      :required="true"
      :minLen="2"
      :maxLen="4096"
      :placeholder="`Editing...`"
      :regex="/^[^<>\[\]]*$/"
      :regexMsg="InvalidCharacters"
      @error="err.comment = $event"
    >
      <div class="flex items-center mr-1.5">
        <basic-btn
          class="text-gray-700 dark:text-gray-300 p-1 mr-1.5"
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

<script>
import { computed, defineComponent, reactive, toRefs } from "vue";
import { useStore } from "vuex";

import { API_Comment } from "../../../api/post/comment";
import UpdateComment from "../../../api/post/comment/update";
import { InvalidCharacters, SomethingWentWrong } from "../../form/errors";
import { notNull, RespToError } from "../../../models/responses";
import { UnixTime } from "../../../pkg/time";

import InputLabel from "../../form/InputLabel.vue";
import TextInput from "../../form/TextInput.vue";
import BasicBtn from "../../button/BasicBtn.vue";
import UserName, { OneLineLink, NoName } from "../../user/Name.vue";

export default defineComponent({
  name: "comment-input",
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
    const store = useStore();
    const storeName = `p${props.postId}-comments`;
    const loggedIn = computed(() => store.getters.loggedIn),
      storeUser = computed(() => store.getters.user);
    const data = reactive({
      // Don't use computed value so it can be edited
      comment: store.getters[`${storeName}/comment`](props.commentId),
      submitted: false,
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