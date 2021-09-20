<template>
  <component
    :is="isLink ? 'to-link' : 'div'"
    :to="'/u/' + user.userName"
    class="block flex-shrink-0"
  >
    <img
      :src="user.avatar ? user.avatarUrl : gray5x5Url"
      loading="auto"
      :alt="`@${user.userName}'s Profile Image`"
      :width="500"
      :height="500"
      class="w-full h-full rounded-full mr-2 border-none"
      :class="
        user.avatar
          ? 'bg-gray-50 dark:bg-gray-900'
          : 'bg-gray-200 dark:bg-gray-800'
      "
    />
  </component>
</template>

<script>
import { defineComponent } from "vue";
import { GEN_User } from "../../models/go_structs_gen";
import gray5x5Url from "../../assets/gray5x5.jpg";

import ToLink from "../ToLink.vue";

export default defineComponent({
  props: {
    user: {
      type: GEN_User,
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
  components: { ToLink },
});
</script>
