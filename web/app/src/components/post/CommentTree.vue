<template>
  <div class="">
    <snackbar
      v-model="err.show"
      :close-btn="true"
      class="text-gray-800 dark:text-gray-200"
    >
      {{ err.msg }}
    </snackbar>
    <article class="flex flex-shrink flex-grow flex-col">
      <div class="flex flex-shrink flex-grow">
        <div class="flex flex-grow flex-col">
          <div
            class="flex items-center justify-between text-gray-700 dark:text-gray-300"
          >
            <div class="flex items-center text-sm xs:text-base">
              <user-avatar
                :user="user"
                :is-link="true"
                class="mr-2 h-6 w-6 xs:h-[30px] xs:w-[30px]"
              ></user-avatar>
              <user-name :user="user"></user-name>
              <div class="dot-before flex items-center whitespace-nowrap">
                <fmt-time
                  :time="comment.created"
                  :type="Fmt_Short_Time"
                  class="group-hover:underline"
                ></fmt-time>
              </div>
              <div
                v-if="edited"
                class="dot-before flex items-center whitespace-nowrap"
              >
                <div title="Edited">
                  <PencilAltIcon
                    class="h-4 w-4 dark:text-gray-400"
                  ></PencilAltIcon>
                </div>
              </div>
            </div>
          </div>
          <div class="flex">
            <div
              class="group mt-1 mr-2 flex flex-grow-0 cursor-pointer pt-1.5 xs:w-[30px]"
              :class="
                comment.depth == 0
                  ? 'w-6 justify-center'
                  : 'w-2 justify-start xs:justify-center'
              "
              @click="hidden = !hidden"
            >
              <div
                class="w-0 border border-gray-400 group-hover:border-gray-700 dark:border-gray-700 dark:group-hover:border-gray-400"
              ></div>
            </div>
            <div v-if="hidden" class="flex flex-grow">
              <small-text class="pt-1.5">[[hidden]]</small-text>
            </div>
            <div v-else class="flex flex-grow flex-col">
              <div class="py-1.5">
                <div v-if="!showEdit" class="leading-5">
                  {{ comment.content }}
                </div>
                <div v-else>
                  <comment-edit
                    :post-id="comment.postId"
                    :comment-id="comment.id"
                    @edited="updated($event)"
                    @close="showEdit = false"
                  ></comment-edit>
                </div>
              </div>
              <div
                class="mt-2 flex flex-row-reverse items-center text-sm font-medium text-gray-600 dark:text-gray-400 xs:flex-row"
              >
                <small-btn
                  class="mr-1.5 flex items-center pr-0.5"
                  @click="showReplyForm = !showReplyForm"
                >
                  <component
                    :is="
                      showReplyForm
                        ? 'OutlineAnnotationIcon'
                        : 'SolidAnnotationIcon'
                    "
                    class="h-4 w-4"
                  ></component>
                  <div class="pl-1">Reply</div>
                </small-btn>
                <!-- TODO: <small-btn class="pr-0.5 mr-1.5">
                  <div class="px-0.5">Share</div>
                </small-btn> -->
                <div
                  v-if="isCommenter"
                  class="flex flex-row-reverse xs:flex-row"
                >
                  <small-btn
                    class="mr-1.5 flex items-center pr-0.5"
                    @click="showEdit = !showEdit"
                  >
                    <div class="px-0.5">Edit</div>
                  </small-btn>
                  <small-btn
                    class="mr-1.5 flex items-center pr-0.5"
                    @click="showDelete = !showDelete"
                  >
                    <div class="px-0.5">Delete</div>
                  </small-btn>
                  <comment-delete
                    :show="showDelete"
                    :post-id="comment.postId"
                    :comment-id="comment.id"
                    @close="showDelete = false"
                    @deleted="deleted()"
                  ></comment-delete>
                </div>
                <!-- TODO: <small-btn v-else class="pr-0.5 mr-1.5">
                  <div class="px-0.5">Report</div>
                </small-btn>-->
              </div>
              <div v-if="loggedIn" class="pt-3">
                <comment-input
                  v-if="showReplyForm"
                  :post-id="String(comment.postId)"
                  :reply-to="comment.id"
                  @commented="newReply($event)"
                ></comment-input>
              </div>
              <sign-up-overlay
                v-else
                :show="showReplyForm"
                :redirect="`/post/${postId}`"
                @close="showReplyForm = false"
              >
                <div>Sign Up or Login to Reply</div>
              </sign-up-overlay>
              <feed v-if="comment.hasChildren()" class="flex-col-reverse">
                <div
                  v-for="(child, index) in comment.children"
                  :key="index"
                  class="pt-3"
                >
                  <comment-tree
                    :comment-id="child.id"
                    :post-id="postId"
                  ></comment-tree>
                </div>
              </feed>
              <div v-if="missingReplies > 0" class="mt-2">
                <div class="link text-sm font-semibold" @click="getReplies()">
                  {{ missingReplies }}
                  {{ missingReplies > 1 ? "Replies" : "Reply" }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </article>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, reactive, toRefs } from "vue";
import { useCommentStore } from "@/store/comments";
import { useSessionStore } from "@/store/session";

import { API_Comment } from "@/api/post/comment";
import { replies } from "@/api/post/comment/replies";
import { User } from "@/models/user";

import { RespToError } from "@/models/responses";
import { SomethingWentWrong } from "../form/errors";

import PencilAltIcon from "@heroicons/vue/solid/PencilAltIcon";
import SolidAnnotationIcon from "@heroicons/vue/solid/AnnotationIcon";
import OutlineAnnotationIcon from "@heroicons/vue/outline/AnnotationIcon";

import UserName from "../user/Name.vue";
import FmtTime, { Fmt_Short_Time } from "../FmtTime.vue";
import UserAvatar from "../user/Avatar.vue";
import SmallBtn from "../button/SmallBtn.vue";
import SmallText from "../text/Small.vue";
import CommentInput from "./comment/CommentInput.vue";
import Feed from "../Feed.vue";
import CommentDelete from "./comment/CommentDelete.vue";
import CommentEdit from "./comment/CommentEdit.vue";
import SignUpOverlay from "../overlay/SignUpOverlay.vue";
import Snackbar from "../box/Snackbar.vue";

export default defineComponent({
  name: "CommentTree",
  props: {
    commentId: {
      type: Number,
      required: true,
    },
    postId: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const session = useSessionStore();
    const commentStore = useCommentStore(props.postId);

    const data = reactive({
      hidden: false,
      showReplyForm: false,
      showEdit: false,
      showDelete: false,
      // Replies page
      page: 0,
      err: {
        msg: "",
        show: false,
      },
    });

    return {
      ...toRefs(data),
      loggedIn: computed(() => session.loggedIn),
      comment: computed(
        () => commentStore.comment(props.commentId) || new API_Comment()
      ),
      missingReplies: computed(() =>
        commentStore.missingReplies(props.commentId)
      ),
      storeUser: computed(() => session.user),
      addComment: (c: API_Comment, isNew: boolean) =>
        commentStore.addComment(c, isNew),
      addUser: (u: User) => commentStore.addUser(u),
      commentUser: computed(() => commentStore.user),
      removeComment: (id: number) => commentStore.remove(id),
      updateComment: (c: API_Comment) => commentStore.updateComment(c),
      Fmt_Short_Time,
      API_Comment,
    };
  },
  computed: {
    edited() {
      return this.comment.updated > this.comment.created + 60;
    },
    user() {
      return this.commentUser(this.comment.userId);
    },
    isCommenter() {
      return this.loggedIn && this.storeUser.id == this.comment.userId;
    },
  },
  methods: {
    newReply(c: API_Comment) {
      this.showReplyForm = false;
      this.addUser(this.storeUser);
      this.addComment(c, true);
    },
    deleted() {
      this.removeComment(this.comment.id);
    },
    updated(c: API_Comment) {
      this.updateComment(c);
    },
    getReplies() {
      replies(this.postId, this.commentId, this.page)
        .then((r) => {
          // Add users for replies
          for (const k in r.users) {
            if (Object.hasOwnProperty.call(r.users, k)) {
              this.addUser(r.users[k]);
            }
          }
          // Add replies
          for (const k in r.order) {
            const key = Number(k).valueOf();
            if (Object.hasOwnProperty.call(r.order, key)) {
              this.addComment(r.comment(r.order[key])!, false);
            }
          }
        })
        .catch((e) => {
          let errText = RespToError(e);
          if (errText.length > 0) {
            this.err.msg = errText;
          } else {
            this.err.msg = SomethingWentWrong;
          }
          this.err.show = true;
        });
    },
  },
  components: {
    FmtTime,
    UserName,
    UserAvatar,
    PencilAltIcon,
    SolidAnnotationIcon,
    OutlineAnnotationIcon,
    SmallBtn,
    SmallText,
    CommentInput,
    Feed,
    CommentDelete,
    CommentEdit,
    SignUpOverlay,
    Snackbar,
  },
});
</script>
