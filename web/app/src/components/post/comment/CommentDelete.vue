<template>
  <overlay v-model="visible" :xsFullHeight="false" :xsSelfEnd="true">
    <snackbar
      v-model="err.show"
      :closeBtn="true"
      class="text-gray-800 dark:text-gray-200"
    >
      {{ err.msg }}
    </snackbar>
    <box
      class="w-full xs:m-3 rounded-t-lg xs:rounded-md shadow-lg focus:outline-none overflow-hidden flex flex-col self-end"
    >
      <div class="text-gray-800 dark:text-gray-200 py-4 px-5">
        <h3 class="text-xl">Delete Comment</h3>
        <h4 class="text-base text-rose-800 dark:text-red-600">
          This action is irreversible
        </h4>
      </div>
      <div class="bg-gray-200 dark:bg-gray-800 flex justify-end py-2 px-3">
        <basic-btn class="py-2 px-1.5 mr-1.5" @click="visible = false">
          Cancel
        </basic-btn>
        <basic-btn
          class="py-2 px-1.5 text-rose-800 dark:text-red-600"
          @click="deleteComment"
          :disabled="submitted"
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
