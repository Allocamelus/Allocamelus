<template>
  <article
    class="flex flex-shrink flex-grow flex-col"
    :class="isLink ? 'cursor-pointer' : ''"
  >
    <div class="flex flex-shrink flex-grow py-3 px-3.5" @click.self="toPost">
      <user-avatar :user="user" :is-link="true" class="h-11 w-11"></user-avatar>
      <div
        class="ml-3 flex flex-grow flex-col"
        :class="post.content?.length == 0 ? 'justify-center' : ''"
        @click.self="toPost"
      >
        <div
          class="flex items-center justify-between text-gray-700 dark:text-gray-300"
          @click.self="toPost"
        >
          <div class="flex">
            <user-name :user="user"></user-name>
            <div
              v-if="published"
              class="dot-before flex items-center whitespace-nowrap"
            >
              <to-link :to="link" class="group no-underline">
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
          <dots-dropdown class="ml-3">
            <div class="bg-secondary-800">WIP</div>
          </dots-dropdown>
        </div>
        <div
          :class="[
            isLink ? 'cursor-pointer' : '',
            dynamicContent ? ['text-lg', 'sm:text-xl'] : '',
          ]"
          @click="toPost"
          v-html="purifiedContent /* skipcq: JS-0693 */"
        ></div>
        <div
          v-if="post.media"
          class="mt-2 flex flex-wrap overflow-hidden rounded-lg"
          @click="toPost"
        >
          <image-box
            v-for="(media, key) in post.mediaList"
            :key="key"
            :index="key"
            :url="media.url"
            :alt="media.meta.alt"
            :width="media.meta.width"
            :height="media.meta.height"
            :total-number="post.mediaList.length"
            :rounded="false"
            loading="lazy"
          >
          </image-box>
        </div>
      </div>
    </div>
  </article>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { sanitize } from "@/pkg/sanitize";

import { local } from "@/pkg/url";

import PencilAltIcon from "@heroicons/vue/solid/PencilAltIcon";
import RadixEyeNone from "../icons/RadixEyeNone.vue";

import UserName from "../user/Name.vue";
import FmtTime, { Fmt_Short_Time } from "../FmtTime.vue";
import UserAvatar from "../user/Avatar.vue";
import ToLink from "../ToLink.vue";
import ImageBox from "../box/ImageBox.vue";
import DotsDropdown from "../menu/DotsDropdown.vue";

import { User } from "@/models/user";
import { Post } from "@/models/post";

export default defineComponent({
  name: "PostBox",
  props: {
    post: {
      type: Post,
      required: true,
    },
    user: {
      type: User,
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
    purifiedContent() {
      return sanitize(this.post.content);
    },
  },
  methods: {
    toPost(e: PointerEvent) {
      // Open L
      if (e.target !== null) {
        let el = e.target as HTMLElement;
        if (el instanceof HTMLAnchorElement) {
          if (el.href.length > 0) {
            if (!local(el.href)) {
              window.open(el.href, "_blank");
              e.preventDefault();
              return;
            }
          }
        }
      }
      if (this.isLink) {
        this.$router.push(this.link);
      }
    },
  },
  components: {
    FmtTime,
    UserName,
    UserAvatar,
    PencilAltIcon,
    RadixEyeNone,
    ToLink,
    ImageBox,
    DotsDropdown,
  },
});
</script>
