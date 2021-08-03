<template>
  <article class="flex flex-col flex-grow flex-shrink">
    <div class="flex flex-grow flex-shrink py-2 px-3">
      <user-avatar :user="user" :isLink="true" class="w-8 h-8"></user-avatar>
      <div class="ml-3 flex flex-col flex-grow">
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
          <div class="flex">
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
          <!-- TODO: Real options -->
          <dots-dropdown class="ml-3">
            <div class="bg-secondary-800">WIP</div>
          </dots-dropdown>
        </div>
        <div>{{ comment.content }}</div>
      </div>
    </div>
  </article>
</template>

<script>
import { defineComponent } from "vue";

import { GEN_User, GEN_Comment } from "../../models/go_structs_gen";

import PencilAltIcon from "@heroicons/vue/solid/PencilAltIcon";
import RadixEyeNone from "../icons/RadixEyeNone.vue";

import UserName from "../user/Name.vue";
import FmtTime, { Fmt_Short_Time } from "../FmtTime.vue";
import CircleBg from "../button/CircleBg.vue";
import UserAvatar from "../user/Avatar.vue";
import ToLink from "../ToLink.vue";
import ImageBox from "../box/ImageBox.vue";
import DotsDropdown from "../menu/DotsDropdown.vue";

export default defineComponent({
  name: "comment-box",
  props: {
    comment: {
      type: GEN_Comment,
      required: true,
    },
    user: {
      type: GEN_User,
      required: true,
    },
  },
  setup() {
    return {
      Fmt_Short_Time,
    };
  },
  computed: {
    edited() {
      if (this.comment.updated > this.comment.created + 60) {
        return true;
      }
    },
  },
  methods: {},
  components: {
    FmtTime,
    UserName,
    CircleBg,
    UserAvatar,
    PencilAltIcon,
    RadixEyeNone,
    ToLink,
    ImageBox,
    DotsDropdown,
  },
});
</script>
