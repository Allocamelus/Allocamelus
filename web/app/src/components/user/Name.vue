<template>
  <component
    :is="isLink ? 'router-link' : 'div'"
    class="name-container"
    :class="displayType"
    :to="'/u/' + user.userName"
  >
    <div class="name">{{ user.name }}</div>
    <div class="user-name">@{{ user.userName }}</div>
  </component>
</template>

<script>
import { defineComponent } from "vue";
import { User } from "../../models/user_gen";

export const OneLineLink = "one-line";
export const TwoLine = "two-line";

export default defineComponent({
  props: {
    user: {
      type: User,
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

<style lang="scss" scoped>
@import "@/scss/vars";

@layer components {
  .one-line {
    &.name-container {
      @apply pb-1;
    }
    @apply min-w-0 truncate;
    .user-name {
      @apply ml-1;
    }
    a {
      @apply no-underline;
    }
    &:hover .name {
      @apply underline;
    }
  }
  .two-line {
    &.name-container {
      @apply flex flex-col items-start;
      .name {
        @apply text-xl mb-0.5;
      }
    }
  }
  .name-container {
    @apply min-w-0 items-center text-gray-800 dark:text-gray-200 truncate;
    /*
      overflow: hidden;
      overflow-wrap: break-word;
      */
  }

  .name {
    @apply inline whitespace-nowrap align-middle font-semibold;
    // white-space: nowrap;
  }

  .user-name {
    @apply inline whitespace-nowrap align-middle text-gray-700 dark:text-gray-400 text-sm font-normal;
    /* white-space: nowrap;
    overflow: hidden;
    overflow-wrap: break-word;*/
  }
}
</style>