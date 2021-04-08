<template>
  <component
    :is="isLink ? 'router-link' : 'div'"
    :to="'/u/' + user.userName"
    class="block"
  >
    <component
      :is="user.avatar ? 'img' : 'div'"
      :src="user.avatar ? user.avatarUrl : ''"
      loading="auto"
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
});
</script>