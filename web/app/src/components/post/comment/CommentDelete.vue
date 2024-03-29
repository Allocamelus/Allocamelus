<template>
  <overlay v-model="visible" :xs-full-height="false" :xs-self-end="true">
    <snackbar
      v-model="err.show"
      :close-btn="true"
      class="text-gray-800 dark:text-gray-200"
    >
      {{ err.msg }}
    </snackbar>
    <box
      class="flex w-full flex-col self-end overflow-hidden rounded-t-lg shadow-lg focus:outline-none xs:m-3 xs:rounded-md"
    >
      <div class="py-4 px-5 text-gray-800 dark:text-gray-200">
        <h3 class="text-xl">Delete Comment</h3>
        <h4 class="text-base text-rose-800 dark:text-red-600">
          This action is irreversible
        </h4>
      </div>
      <div class="flex justify-end bg-gray-200 py-2 px-3 dark:bg-gray-800">
        <basic-btn class="mr-1.5 py-2 px-1.5" @click="visible = false">
          Cancel
        </basic-btn>
        <basic-btn
          class="py-2 px-1.5 text-rose-800 dark:text-red-600"
          :disabled="submitted"
          @click="deleteComment"
        >
          Delete
        </basic-btn>
      </div>
    </box>
  </overlay>
</template>

<script>
import { defineComponent, reactive, toRefs } from "vue";

import remove from "../../../api/post/comment/remove";
import { notNull, RespToError } from "../../../models/responses";
import { SomethingWentWrong } from "../../form/errors";

import Overlay from "../../overlay/Overlay.vue";
import Box from "../../box/Box.vue";
import BasicBtn from "../../button/BasicBtn.vue";
import Snackbar from "../../box/Snackbar.vue";

export default defineComponent({
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    postId: {
      type: Number,
      required: true,
    },
    commentId: {
      type: Number,
      required: true,
    },
  },
  emits: ["close", "deleted"],
  setup(props) {
    // visible is used for overlay instead of show to keep parent and overlay in sync
    const data = reactive({
      visible: props.show,
      submitted: false,
      err: {
        msg: "",
        show: false,
      },
    });

    return {
      ...toRefs(data),
    };
  },
  watch: {
    show(newValue) {
      this.visible = newValue;
    },
    visible(newValue) {
      if (!newValue) {
        this.close();
      }
    },
  },
  methods: {
    deleteComment() {
      this.submitted = true;

      remove(this.postId, this.commentId)
        .then((r) => {
          this.handleResp(r.error);
        })
        .catch((e) => {
          this.handleResp(e);
        });
    },
    close() {
      this.$emit("close");
    },
    deleted() {
      this.$emit("deleted");
    },
    handleResp(e) {
      this.submitted = false;
      if (notNull(e)) {
        let errText = RespToError(e);
        if (errText.length > 0) {
          this.err.msg = errText;
        } else {
          this.err.msg = SomethingWentWrong;
        }
        this.err.show = true;
      } else {
        this.visible = false;
        this.deleted();
      }
    },
  },
  components: {
    Overlay,
    Box,
    BasicBtn,
    Snackbar,
  },
});
</script>
