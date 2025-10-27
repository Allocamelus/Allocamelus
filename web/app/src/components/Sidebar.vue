<template>
  <div class="ml-6 hidden w-64 flex-shrink-0 flex-col items-stretch md:flex">
    <box v-if="loggedIn" class="rounded-xl px-4 py-3">
      <div class="mx-auto w-max">
        <basic-btn
          to="/post/new"
          class="px-3.5 py-2"
          :class="buttonStyle.secondary"
        >
          New Post
        </basic-btn>
      </div>
    </box>
    <box v-else class="rounded-xl px-3 py-2">
      <div class="mx-auto flex w-max">
        <basic-btn
          to="/signup"
          class="px-3.5 py-2"
          :class="buttonStyle.secondary"
        >
          Sign Up
        </basic-btn>
        <basic-btn to="/login" class="link ml-1.5 px-3 py-2"> Login </basic-btn>
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

<style
  src="@/scss/modules/button.modules.scss"
  lang="scss"
  module="buttonStyle"
  scoped
></style>
