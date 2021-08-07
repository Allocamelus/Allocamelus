<template>
  <div class="">
    <article class="flex flex-col flex-grow flex-shrink">
      <div class="flex flex-grow flex-shrink">
        <div class="flex flex-col flex-grow">
          <div
            class="
              text-gray-700
              dark:text-gray-300
              flex
              items-center
              justify-between
            "
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
              class="mt-1 pt-1.5 flex xs:w-[30px] mr-2 flex-grow-0 group"
              :class="
                comment.depth == 0
                  ? 'w-6 justify-center'
                  : 'w-2 justify-start xs:justify-center'
              "
            >
              <div
                class="w-0 border-[1px] border-gray-400 dark:border-gray-700"
                :class="[
                  /*TODO Click to hide group-hover:border-gray-700 dark:group-hover:border-gray-400*/
                ]"
              ></div>
            </div>
            <div class="flex flex-col flex-grow">
              <div class="py-1.5 leading-5">{{ comment.content }}</div>
              <div
                class="
                  flex flex-row-reverse
                  xs:flex-row
                  text-sm
                  font-medium
                  mt-2
                  text-gray-600
                  dark:text-gray-400
                "
              >
                <small-btn class="flex pr-0.5 mr-2">
                  <AnnotationIcon class="h-4 w-4"></AnnotationIcon>
                  <div class="pl-1">Reply</div>
                </small-btn>
                <small-btn class="flex pr-0.5 mr-2">
                  <div class="px-0.5">Share</div>
                </small-btn>
              </div>
              <div v-if="comment.hasChildren()">
                <div
                  v-for="(child, index) in comment.children"
                  :key="index"
                  class="pt-3"
                >
                  <comment-tree
                    :comment="API_Comment.createFrom(child)"
                    :userList="userList"
                  ></comment-tree>
                </div>
              </div>
              <div class="mt-2" v-if="missingReplies > 0">
                <a class="link text-sm font-semibold">
                  {{ missingReplies }}
                  {{ missingReplies > 1 ? "Replies" : "Reply" }}
                </a>
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
import { useStore } from "vuex";

import { API_Comment } from "../../api/post/comment";
import { user_list } from "../../models/ordered_list";

import PencilAltIcon from "@heroicons/vue/solid/PencilAltIcon";
import AnnotationIcon from "@heroicons/vue/solid/AnnotationIcon";

import UserName from "../user/Name.vue";
import FmtTime, { Fmt_Short_Time } from "../FmtTime.vue";
import CircleBg from "../button/CircleBg.vue";
import UserAvatar from "../user/Avatar.vue";
import ToLink from "../ToLink.vue";
import ImageBox from "../box/ImageBox.vue";
import SmallBtn from "../button/SmallBtn.vue";

export default defineComponent({
  name: "comment-tree",
  props: {
    comment: {
      type: API_Comment,
      default: new API_Comment(),
    },
    userList: {
      type: user_list,
      default: new user_list(),
    },
  },
  setup() {
    const store = useStore();
    const loggedIn = computed(() => store.getters.loggedIn),
      storeUser = computed(() => store.getters.user);

    return {
      loggedIn,
      storeUser,
      Fmt_Short_Time,
      API_Comment,
    };
  },
  computed: {
    edited() {
      if (this.comment.updated > this.comment.created + 60) {
        return true;
      }
      return false;
    },
    user() {
      return this.userList.user(this.comment.userId);
    },
    missingReplies() {
      return this.comment.numNotHad();
    },
  },
  components: {
    FmtTime,
    UserName,
    CircleBg,
    UserAvatar,
    PencilAltIcon,
    AnnotationIcon,
    ToLink,
    ImageBox,
    SmallBtn,
  },
});
</script>