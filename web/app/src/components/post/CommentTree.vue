<template>
  <div class="">
    <article class="flex flex-col flex-grow flex-shrink">
      <div class="flex flex-grow flex-shrink">
        <div class="flex flex-col">
          <user-avatar
            :user="user"
            :isLink="true"
            class="w-[30px] h-[30px]"
          ></user-avatar>
          <div class="mt-1 pt-1.5 flex flex-grow justify-center group">
            <div
              class="w-0 border-[1px] border-gray-400 dark:border-gray-700"
              :class="[
                /*TODO Click to hide group-hover:border-gray-700 dark:group-hover:border-gray-400*/
              ]"
            ></div>
          </div>
        </div>
        <div class="flex flex-col flex-grow">
          <div class="ml-2">
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
              <div class="flex h-[30px]">
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
            <div class="py-1.5 leading-5">{{ comment.content }}</div>
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
        </div>
      </div>
    </article>
  </div>
</template>

<script>
import { defineComponent } from "vue";

import { API_Comment } from "../../api/post/comment";
import { user_list } from "../../models/ordered_list";

import PencilAltIcon from "@heroicons/vue/solid/PencilAltIcon";
import RadixEyeNone from "../icons/RadixEyeNone.vue";

import UserName from "../user/Name.vue";
import FmtTime, { Fmt_Short_Time } from "../FmtTime.vue";
import CircleBg from "../button/CircleBg.vue";
import UserAvatar from "../user/Avatar.vue";
import ToLink from "../ToLink.vue";
import ImageBox from "../box/ImageBox.vue";

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
    return {
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
  },
  components: {
    FmtTime,
    UserName,
    CircleBg,
    UserAvatar,
    PencilAltIcon,
    RadixEyeNone,
    ToLink,
    ImageBox,
  },
});
</script>