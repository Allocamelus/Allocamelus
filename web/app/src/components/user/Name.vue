<template>
  <div class="name-container">
    <div class="name">
      {{ user.name }}
    </div>
    <div class="unique-name">@{{ user.uniqueName }}</div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { User } from "../../models/user_gen";
export default defineComponent({
  props: {
    user: {
      type: User,
      required: true,
    },
  },
});
</script>

<style lang="scss" scoped>
.name-container {
  white-space: nowrap;
  display: flex;
}
.name {
  overflow: hidden;
  text-overflow: ellipsis;
}
.unique-name {
  margin-left: 4px;
  color: hsl(0, 0%, 30%);
}
.dark-theme {
  .unique-name {
    color: hsl(0, 0%, 60%);
  }
}
</style>