<template>
  <article
    class="flex flex-col flex-grow flex-shrink"
    :class="isLink ? 'cursor-pointer' : ''"
  >
    <div class="flex flex-grow flex-shrink py-3 px-4" @click.self="toPost">
      <user-avatar :user="user" :isLink="true" class="w-11 h-11"></user-avatar>
      <div
        class="ml-3 flex flex-col flex-grow"
        :class="post.content?.length == 0 ? 'justify-center' : ''"
      >
        <div
          class="text-gray-700 dark:text-gray-300 flex items-center justify-between"
          @click.self="toPost"
        >
          <div class="flex">
            <user-name :user="user"></user-name>
            <div
              v-if="published"
              class="dot-before flex items-center whitespace-nowrap"
            >
              <to-link :to="link" class="no-underline group">
                <fmt-time
                  :time="post.published"
                  :type="Fmt_Short_Time"
                  class="group-hover:underline"
                ></fmt-time>
              </to-link>
            </div>
            <div v-else class="dot-before flex items-center whitespace-nowrap">
              <div title="Not Published">
                <radix-eye-none class="h-4 w-4 dark:text-gray-400" />
              </div>
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
          <circle-bg class="ml-3 hover:bg-rose-800">
            <DotsVerticalIcon
              class="h-4.5 w-4.5 text-gray-800 dark:text-gray-200 group-hover:text-rose-700"
            ></DotsVerticalIcon>
          </circle-bg>
        </div>
        <div
          @click="toPost"
          :class="[
            isLink ? 'cursor-pointer' : '',
            dynamicContent ? ['text-lg', 'sm:text-xl'] : '',
          ]"
          v-html="post.content"
        ></div>
      </div>
    </div>
    <div v-if="post.media" class="flex flex-wrap rounded-b-xl overflow-hidden">
      <image-box
        v-for="(media, key) in post.mediaList"
        :key="key"
        :index="key"
        :url="media.url"
        :alt="media.alt"
        :width="media.width"
        :height="media.height"
        :totalNumber="post.mediaList.length"
        :rounded="false"
        :loading="lazy"
      >
      </image-box>
    </div>
  </article>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

import DotsVerticalIcon from "@heroicons/vue/outline/DotsVerticalIcon";
import PencilAltIcon from "@heroicons/vue/solid/PencilAltIcon";
import RadixEyeNone from "../icons/RadixEyeNone.vue";

import UserName from "../user/Name.vue";
import FmtTime, { Fmt_Short_Time } from "../FmtTime.vue";
import CircleBg from "../button/CircleBg.vue";
import UserAvatar from "../user/Avatar.vue";
import ToLink from "../ToLink.vue";
import ImageBox from "../box/ImageBox.vue";

import { GEN_User, GEN_Post } from "../../models/go_structs_gen";

export default defineComponent({
  name: "post-box",
  props: {
    post: {
      type: GEN_Post,
      required: true,
    },
    user: {
      type: GEN_User,
      required: true,
    },
    isLink: {
      type: Boolean,
      default: false,
    },
    dynamicContent: {
      type: Boolean,
      default: false,
    },
  },
  setup() {
    return {
      Fmt_Short_Time,
    };
  },
  computed: {
    link() {
      return `/post/${this.post.id}`;
    },
    published() {
      if (this.post.published != 0) {
        return true;
      }
      return false;
    },
    edited() {
      if (this.published) {
        if (this.post.updated > this.post.published + 60) {
          return true;
        }
      }
      return false;
    },
  },
  methods: {
    toPost() {
      if (this.isLink) {
        this.$router.push(this.link);
      }
    },
  },
  components: {
    FmtTime,
    UserName,
    DotsVerticalIcon,
    CircleBg,
    UserAvatar,
    PencilAltIcon,
    RadixEyeNone,
    ToLink,
    ImageBox,
  },
});
</script>
