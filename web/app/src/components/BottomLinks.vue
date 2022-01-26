<template>
  <div class="mt-3">
    <div
      class="px-4 text-sm flex justify-center text-gray-800 dark:text-gray-500"
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
    <div class="flex justify-center mt-3">
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

<script lang="ts">
import { defineComponent } from "vue";

import Box from "./box/Box.vue";
import BasicBtn from "./button/BasicBtn.vue";
import Allocamelus from "./icons/Allocamelus.vue";
import RadixGithub from "./icons/RadixGithub.vue";
import ToLink from "./ToLink.vue";

export default defineComponent({
  components: { Box, BasicBtn, ToLink, RadixGithub, Allocamelus },
});
</script>
