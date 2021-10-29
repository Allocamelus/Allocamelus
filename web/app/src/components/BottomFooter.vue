<template>
  <footer id="footer">
    <div class="footer container">
      <div class="flex">
        <div class="font-semibold self-center">
          &copy; {{ new Date().getFullYear() }}
        </div>
        <div class="dash-before">
          <to-link to="/about">About</to-link>
        </div>
      </div>
      <div class="hidden xs:flex">
        <a
          rel="nofollow"
          target="_blank"
          href="https://github.com/Allocamelus/Allocamelus"
        >
          <radix-github class="h-5 w-5" />
        </a>
        <a href="https://www.allocamelus.com/u/Allocamelus" class="ml-3">
          <allocamelus class="h-5 w-5" />
        </a>
      </div>
      <div class="flex">
        <!-- TODO -->
        <to-link to="/tos">Terms</to-link>
        <div class="dash-before">
          <to-link to="/privacy">Privacy</to-link>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
import { defineComponent } from "vue";

import Allocamelus from "./icons/Allocamelus.vue";
import RadixGithub from "./icons/RadixGithub.vue";

import ToLink from "./ToLink.vue";

export default defineComponent({
  components: { ToLink, RadixGithub, Allocamelus },
});
</script>

<style lang="scss" scoped>
@tailwind components;

@layer components {
  #footer {
    @apply bg-footer text-white xs:h-footer;
  }
  .footer {
    @apply flex justify-between items-center py-3;
    > div {
      @apply items-center;
    }
    &,
    a {
      @apply text-sm text-white;
    }
    a {
      @apply font-semibold no-underline cursor-pointer;
    }
  }
}
</style>
