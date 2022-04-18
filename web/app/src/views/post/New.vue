<template>
  <div class="container px-1 py-5 xs:px-2">
    <box class="mx-auto max-w-prose rounded-xl px-2 py-3 xs:px-3.5">
      <div class="flex">
        <div class="hidden flex-shrink-0 xs:block">
          <user-avatar :user="user" class="h-11 w-11"></user-avatar>
        </div>
        <div class="flex flex-grow flex-col xs:ml-3">
          <h2 class="px-1 text-xl font-medium">New Post</h2>
          <form ref="form" class="mt-2" @submit.prevent>
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
