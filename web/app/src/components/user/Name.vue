<template>
  <component
    :is="isLink ? 'to-link' : 'div'"
    class="text-gray-800 dark:text-gray-200 truncate flex"
    :class="isLink ? 'group no-underline items-center' : 'flex-col'"
    :to="'/u/' + user.userName"
  >
    <div
      class="truncate flex-shrink font-semibold"
      :class="isLink ? 'group-hover:underline' : 'text-xl mb-0.5'"
    >
      {{ user.name }}
    </div>
    <div
      class="text-gray-700 dark:text-gray-400 text-sm font-normal flex-none"
      :class="isLink ? 'ml-1' : ''"
    >
      @{{ user.userName }}
    </div>
  </component>
</template>

<script>
import { defineComponent } from "vue";
import { GEN_User } from "../../models/go_structs_gen";

import ToLink from "../ToLink.vue";

export const OneLineLink = "one-line";
export const TwoLine = "two-line";

export default defineComponent({
  props: {
    user: {
      type: GEN_User,
      required: true,
    },
    displayType: {
      type: String,
      default: OneLineLink,
    },
  },
  computed: {
    isLink() {
      if (this.displayType == OneLineLink) {
        return true;
      }
      return false;
    },
  },
  components: {ToLink},
});
</script>
