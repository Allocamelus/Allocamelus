<template>
  <div class="mt-3">
    <div
      class="flex justify-center px-4 text-sm text-neutral-800 dark:text-neutral-500"
    >
      <div class="self-center text-neutral-700 dark:text-neutral-400">
        &copy; {{ new Date().getFullYear() }}
      </div>
      <div class="dash-before">
        <to-link to="/terms" class="link">Terms</to-link>
      </div>
      <div class="dash-before">
        <to-link to="/privacy" class="link">Privacy</to-link>
      </div>
      <div class="dash-before">
        <to-link to="/about" class="link">About</to-link>
      </div>
    </div>
    <div class="mt-3 flex justify-center">
      <to-link to="https://github.com/Allocamelus/Allocamelus">
        <radix-github class="h-5.5 w-5.5" />
      </to-link>
      <to-link to="https://www.allocamelus.com/u/Allocamelus" class="link ml-3">
        <allocamelus class="h-5.5 w-5.5" />
      </to-link>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

import Allocamelus from "./icons/Allocamelus.vue";
import RadixGithub from "./icons/RadixGithub.vue";
import ToLink from "./ToLink.vue";

export default defineComponent({
  components: { ToLink, RadixGithub, Allocamelus },
});
</script>
