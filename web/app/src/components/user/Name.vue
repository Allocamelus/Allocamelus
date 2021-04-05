<template>
  <component
    :is="isLink ? 'router-link' : 'div'"
    class="min-w-0 text-gray-800 dark:text-gray-200 truncate flex items-center"
    :class="
      isLink
        ? 'min-w-0 truncate group no-underline'
        : 'flex-col items-start'
    "
    :to="'/u/' + user.userName"
  >
    <div
      class="whitespace-nowrap align-middle font-semibold"
      :class="isLink ? 'group-hover:underline' : 'text-xl mb-0.5'"
    >
      {{ user.name }}
    </div>
    <div
      class="whitespace-nowrap align-middle text-gray-700 dark:text-gray-400 text-sm font-normal"
      :class="isLink ? 'ml-1' : ''"
    >
      @{{ user.userName }}
    </div>
  </component>
</template>

<script>
import { defineComponent } from "vue";
import { GEN_User } from "../../models/go_structs_gen";

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
});
</script>
