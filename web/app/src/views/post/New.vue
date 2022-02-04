<template>
  <div class="container py-5">
    <box class="py-3 px-3.5 max-w-prose mx-auto rounded-xl">
      <div class="flex">
        <div class="flex-shrink-0">
          <user-avatar :user="user" class="w-11 h-11"></user-avatar>
        </div>
        <div class="ml-3 flex-grow flex flex-col">
          <h2 class="text-xl font-medium">New Post</h2>
          <form @submit.prevent ref="form" class="mt-2">
            <post-form></post-form>
          </form>
        </div>
      </div>
    </box>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import { useSessionStore } from "@/store/session";

import Box from "@/components/box/Box.vue";
import PostForm from "@/components/post/PostForm.vue";
import UserAvatar from "@/components/user/Avatar.vue";

export default defineComponent({
  setup() {
    const session = useSessionStore();
    document.title = `New Post - ${import.meta.env.VITE_SITE_NAME}`;

    return {
      user: computed(() => session.user),
    };
  },
  components: { Box, UserAvatar, PostForm },
});
</script>

<style></style>
