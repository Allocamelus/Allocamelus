<template>
  <div class="container py-5">
    <box class="py-3 px-4 max-w-prose mx-auto rounded-xl">
      <div class="flex">
        <div class="flex-shrink-0">
          <user-avatar
            :user="user"
            class="w-11 h-11 xs:w-14 xs:h-14 sm:w-16 sm:h-16"
          ></user-avatar>
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

<script>
import { defineComponent, toRefs, reactive, computed } from "vue";
import { useStore } from "vuex";

import Box from "../../components/box/Box.vue";
import PostForm from "../../components/post/PostForm.vue";
import UserAvatar from "../../components/user/Avatar.vue";

export default defineComponent({
  setup() {
    const store = useStore(),
      loggedIn = computed(() => store.getters.loggedIn),
      user = computed(() => store.getters.user);
    const data = reactive({});
    document.title = `New Post - ${import.meta.env.VITE_SITE_NAME}`;
    return {
      ...toRefs(data),
      loggedIn,
      user,
    };
  },
  components: { Box, UserAvatar, PostForm },
});
</script>

<style></style>
