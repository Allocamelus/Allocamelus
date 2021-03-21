<template>
  <router-link :to="'/' + user.uniqueName">
    <div class="name-container">
      <div class="name">
        {{ user.name }}
      </div>
      <div class="unique-name">@{{ user.uniqueName }}</div>
    </div>
  </router-link>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

export default defineComponent({
  props: {
    user: {
      type: Object,
      required: true,
    },
  },
});
</script>

<style lang="scss" scoped>
@import "@/scss/vars";

.name-container {
  white-space: nowrap;
  display: flex;
  color: $light-title-color;
  text-decoration: none;
  padding: 4px 0;
  margin-bottom: 4px;
  &:hover .name {
    text-decoration: underline;
  }
}

.name {
  overflow: hidden;
  text-overflow: ellipsis;
  font-weight: 600;
}

.unique-name {
  margin-left: 4px;
  color: hsl(0, 0%, 30%);
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