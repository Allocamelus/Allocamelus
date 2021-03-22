<template>
  <div class="name-wrapper" :class="displayType">
    <router-link
      v-if="isLink"
      class="name-container"
      :to="'/u/' + user.uniqueName"
    >
      <div class="name">{{ user.name }}</div>
      <div class="unique-name">@{{ user.uniqueName }}</div>
    </router-link>
    <div class="name-container" v-else>
      <div class="name">{{ user.name }}</div>
      <div class="unique-name">@{{ user.uniqueName }}</div>
    </div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

export const OneLineLink = "one-line";
export const TwoLine = "two-line";

export default defineComponent({
  props: {
    user: {
      type: Object,
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

.name-wrapper {
  .name-container {
    white-space: nowrap;
    display: flex;
    padding: 4px 0;
    margin-bottom: 4px;
    color: $light-title-color;
  }
  &.one-line {
    .unique-name {
      margin-left: 4px;
    }
    a {
      text-decoration: none;
    }
    &:hover .name {
      text-decoration: underline;
    }
  }
  &.two-line {
    .name-container {
      flex-direction: column;
      .name {
        font-size: 22px;
        margin-bottom: 4px;
      }
    }
  }
}

.name {
  overflow: hidden;
  text-overflow: ellipsis;
  font-weight: 600;
}

.unique-name {
  color: hsl(0, 0%, 30%);
  font-weight: 400;
}

.dark-theme {
  .name-container {
    color: $dark-title-color;
  }
  .unique-name {
    color: hsl(0, 0%, 60%);
  }
}
</style>