<template>
  <div class="">
    <snackbar
      v-model="err.show"
      :closeBtn="true"
      class="text-gray-800 dark:text-gray-200"
    >
      {{ err.msg }}
    </snackbar>
    <article class="flex flex-col flex-grow flex-shrink">
      <div class="flex flex-grow flex-shrink">
        <div class="flex flex-col flex-grow">
          <div
            class="text-gray-700 dark:text-gray-300 flex items-center justify-between"
            @click.self="toPost"
          >
            <div class="flex text-sm xs:text-base items-center">
              <user-avatar
                :user="user"
                :isLink="true"
                class="w-6 h-6 xs:w-[30px] xs:h-[30px] mr-2"
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
              class="mt-1 pt-1.5 flex xs:w-[30px] mr-2 flex-grow-0 group cursor-pointer"
              :class="
                comment.depth == 0
                  ? 'w-6 justify-center'
                  : 'w-2 justify-start xs:justify-center'
              "
              @click="hidden = !hidden"
            >
              <div
                class="w-0 border-[1px] border-gray-400 dark:border-gray-700 group-hover:border-gray-700 dark:group-hover:border-gray-400"
              ></div>
            </div>
            <div v-if="hidden" class="flex flex-grow">
              <small-text class="pt-1.5">[[hidden]]</small-text>
            </div>
            <div v-else class="flex flex-col flex-grow">
              <div class="py-1.5">
                <div v-if="!showEdit" class="leading-5">
                  {{ comment.content }}
                </div>
                <div v-else>
                  <comment-edit
                    :postId="comment.postId"
                    :commentId="comment.id"
                    @edited="updated($event)"
                    @close="showEdit = false"
                  ></comment-edit>
                </div>
              </div>
              <div
                class="flex flex-row-reverse xs:flex-row text-sm font-medium mt-2 items-center text-gray-600 dark:text-gray-400"
              >
                <small-btn
                  class="flex items-center pr-0.5 mr-1.5"
                  @click="showReplyForm = !showReplyForm"
                >
                  <component
                    class="h-4 w-4"
                    :is="
                      showReplyForm
                        ? 'OutlineAnnotationIcon'
                        : 'SolidAnnotationIcon'
                    "
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
                    class="flex items-center pr-0.5 mr-1.5"
                    @click="showEdit = !showEdit"
                  >
                    <div class="px-0.5">Edit</div>
                  </small-btn>
                  <small-btn
                    class="flex items-center pr-0.5 mr-1.5"
                    @click="showDelete = !showDelete"
                  >
                    <div class="px-0.5">Delete</div>
                  </small-btn>
                  <comment-delete
                    :show="showDelete"
                    :postId="comment.postId"
                    :commentId="comment.id"
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
                  :postId="String(comment.postId)"
                  :replyTo="comment.id"
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
                    :commentId="child.id"
                    :postId="postId"
                  ></comment-tree>
                </div>
              </feed>
              <div class="mt-2" v-if="missingReplies > 0">
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

<script>
import { computed, defineComponent, reactive, toRefs } from "vue";
import { useStore } from "../../store";

import { API_Comment } from "../../api/post/comment";
import { replies } from "../../api/post/comment/replies";

import { RespToError } from "../../models/responses";
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
  name: "comment-tree",
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
    const store = useStore();
    const storeName = `p${props.postId}-comments`;
    const loggedIn = computed(() => store.getters.loggedIn),
      comment = computed(() =>
        store.getters[`${storeName}/comment`](props.commentId)
      ),
      missingReplies = computed(() =>
        store.getters[`${storeName}/missingReplies`](props.commentId)
      ),
      storeUser = computed(() => store.getters.user),
      commentUser = computed(() => store.getters[`${storeName}/user`]),
      removeComment = (id) => store.commit(`${storeName}/remove`, id),
      addComment = (c) => store.commit(`${storeName}/addComment`, c),
      addUser = (c) => store.commit(`${storeName}/addUser`, c),
      updateComment = (c) => store.commit(`${storeName}/updateComment`, c);

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
      loggedIn,
      comment,
      missingReplies,
      storeUser,
      addComment,
      addUser,
      commentUser,
      removeComment,
      updateComment,
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
    newReply(c) {
      this.showReplyForm = false;
      this.addUser(this.storeUser);
      this.addComment({
        // AddCommentParams
        newComment: true,
        comment: c,
      });
    },
    deleted() {
      this.removeComment(this.comment.id);
    },
    updated(c) {
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
            if (Object.hasOwnProperty.call(r.order, k)) {
              this.addComment({
                // AddCommentParams
                newComment: false,
                comment: r.comment(r.order[k]),
              });
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
