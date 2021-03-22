<template>
  <div class="name-wrapper" :class="displayType">
    <router-link
      v-if="isLink"
      class="name-container"
      :to="'/u/' + user.uniqueName"
    >
      <div class="name">{{ user.name }} Long Name For Testing</div>
      <div class="unique-name">
        @{{ user.uniqueName }}_Long_Name_For_Testing
      </div>
    </router-link>
    <div class="name-container" v-else>
      <div class="name">{{ user.name }} Long Name For Testing</div>
      <div class="unique-name">
        @{{ user.uniqueName }}_Long_Name_For_Testing
      </div>
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
    display: flex;
    padding: 3px 0;
    color: $light-text-3;

    overflow: hidden;
    overflow-wrap: break-word;
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
      margin-bottom: 4px;
      .name {
        font-size: 22px;
        margin-bottom: 4px;
      }
    }
  }
}

.name {
  white-space: nowrap;
  font-weight: 600;
}

.unique-name {
  white-space: nowrap;
  overflow: hidden;
  overflow-wrap: break-word;
  color: $light-text-12;
  font-size: 15px;
  font-weight: 400;
}

.dark-theme {
  .name-container {
    color: $dark-text-4;
  }
  .unique-name {
    color: $dark-text-26;
  }
}
</style>