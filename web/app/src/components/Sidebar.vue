<template>
  <div class="hidden ml-6 w-64 items-stretch md:flex flex-col flex-shrink-0">
    <box v-if="loggedIn" class="py-3 px-4 rounded-xl">
      <div class="w-max mx-auto">
        <basic-btn
          to="/post/new"
          class="text-white bg-secondary-700 hover:bg-secondary-800 py-2 px-3.5"
        >
          New Post
        </basic-btn>
      </div>
    </box>
    <box v-else class="py-2 px-3 rounded-xl">
      <div class="flex w-max mx-auto">
        <basic-btn
          to="/signup"
          class="text-white bg-secondary-700 hover:bg-secondary-800 py-2 px-3.5"
        >
          Sign Up
        </basic-btn>
        <basic-btn to="/login" class="ml-1.5 py-2 px-3 link"> Login </basic-btn>
      </div>
    </box>
    <slot></slot>
    <BottomLinks></BottomLinks>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import { useSessionStore } from "@/store/session";

import Box from "./box/Box.vue";
import BasicBtn from "./button/BasicBtn.vue";
import BottomLinks from "./BottomLinks.vue";

export default defineComponent({
  setup() {
    const session = useSessionStore();
    return {
      loggedIn: computed(() => session.loggedIn),
    };
  },
  components: { Box, BasicBtn, BottomLinks },
});
</script>
