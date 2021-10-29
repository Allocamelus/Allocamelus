<template>
  <div class="hidden ml-6 w-64 items-stretch md:flex flex-col flex-shrink-0">
    <box v-if="loggedIn" class="py-3 px-4 rounded-xl mb-3">
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
    <div
      class="
        px-4
        mb-3
        text-sm
        flex
        justify-center
        text-gray-800
        dark:text-gray-500
      "
    >
      <div class="self-center text-gray-700 dark:text-gray-400">
        &copy; {{ new Date().getFullYear() }}
      </div>
      <div class="dash-before">
        <to-link to="/tos" class="link">Terms</to-link>
      </div>
      <div class="dash-before">
        <to-link to="/privacy" class="link">Privacy</to-link>
      </div>
      <div class="dash-before">
        <to-link to="/about" class="link">About</to-link>
      </div>
    </div>
    <div class="flex justify-center">
      <a
        rel="noopener noreferrer"
        target="_blank"
        href="https://github.com/Allocamelus/Allocamelus"
      >
        <radix-github class="h-5.5 w-5.5" />
      </a>
      <a href="https://www.allocamelus.com/u/Allocamelus" class="ml-3 link">
        <allocamelus class="h-5.5 w-5.5" />
      </a>
    </div>
  </div>
</template>

<script>
import { defineComponent, computed } from "vue";
import { useStore } from "vuex";

import Box from "./box/Box.vue";
import BasicBtn from "./button/BasicBtn.vue";
import Allocamelus from "./icons/Allocamelus.vue";
import RadixGithub from "./icons/RadixGithub.vue";
import ToLink from "./ToLink.vue";

export default defineComponent({
  setup() {
    const store = useStore();
    return {
      loggedIn: computed(() => store.getters.loggedIn),
    };
  },
  components: { Box, BasicBtn, ToLink, RadixGithub, Allocamelus },
});
</script>
