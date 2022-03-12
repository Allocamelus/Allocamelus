<template>
  <component
    :is="isLink ? 'to-link' : 'div'"
    :to="'/u/' + user.userName"
    class="flex-shrink-0 block"
  >
    <img
      :src="hasAvatar ? fullAvatarURL : gray5x5Url"
      loading="auto"
      :alt="`@${user.userName}'s Profile Image`"
      :width="500"
      :height="500"
      class="object-cover max-w-full mr-2 border-none rounded-full"
      :class="
        hasAvatar
          ? 'bg-gray-50 dark:bg-gray-900'
          : 'bg-gray-200 dark:bg-gray-800'
      "
    />
  </component>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { User } from "../../models/user";
import gray5x5Url from "../../assets/gray5x5.jpg";
import FullURL from "../../pkg/fullUrl";

import ToLink from "../ToLink.vue";

export default defineComponent({
  props: {
    user: {
      type: User,
      required: true,
    },
    isLink: {
      type: Boolean,
      default: false,
    },
  },
  setup() {
    return {
      gray5x5Url,
    };
  },
  computed: {
    fullAvatarURL() {
      return FullURL(this.user.avatar, import.meta.env.BASE_URL);
    },
    hasAvatar() {
      return this.user.avatar.length > 0;
    },
  },
  components: { ToLink },
});
</script>
