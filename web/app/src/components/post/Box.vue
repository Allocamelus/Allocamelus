<template>
  <div>
    <div
      class="text-gray-700 dark:text-gray-300 flex items-center truncate justify-between"
    >
      <div class="flex min-w-0">
        <user-name :user="user"></user-name>
        <div class="dot-before flex items-center">
          <router-link :to="link" class="no-underline group">
            <fmt-time
              :time="post.published"
              :type="Fmt_Short_Time"
              class="group-hover:underline"
            ></fmt-time>
          </router-link>
        </div>
      </div>
      <!-- TODO: license https://github.com/tailwindlabs/heroicons/blob/master/LICENSE -->
      <!-- TODO: Real options -->
      <circle-bg class="ml-3">
        <DotsVerticalIcon
          class="h-4.5 w-4.5 text-gray-800 dark:text-gray-200 group-hover:text-rose-700"
        ></DotsVerticalIcon>
      </circle-bg>
    </div>
    <div
      @click="textClick"
      :class="[
        isLink ? 'cursor-pointer' : '',
        dynamicContent ? ['text-lg', 'sm:text-xl'] : '',
      ]"
      v-html="post.content"
    ></div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

import UserName from "../user/Name.vue";
import FmtTime, { Fmt_Short_Time } from "../FmtTime.vue";
import DotsVerticalIcon from "@heroicons/vue/outline/DotsVerticalIcon";
import CircleBg from "../button/CircleBg.vue";

export default defineComponent({
  name: "post-box",
  props: {
    post: {
      type: Object,
    },
    user: {
      type: Object,
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
  },
  methods: {
    textClick() {
      if (this.isLink) {
        this.$router.push(this.link);
      }
    },
  },
  components: { FmtTime, UserName, DotsVerticalIcon, CircleBg },
});
</script>