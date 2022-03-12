<template>
  <footer id="footer" class="text-white bg-footer xs:h-footer">
    <div class="container" :class="$style.footer">
      <div class="flex">
        <div class="self-center font-semibold">
          &copy; {{ new Date().getFullYear() }}
        </div>
        <div class="dash-before">
          <to-link to="/about">About</to-link>
        </div>
      </div>
      <div class="hidden xs:flex">
        <a
          rel="noopener noreferrer"
          target="_blank"
          href="https://github.com/Allocamelus/Allocamelus"
        >
          <radix-github class="w-5 h-5" />
        </a>
        <a href="https://www.allocamelus.com/u/Allocamelus" class="ml-3">
          <allocamelus class="w-5 h-5" />
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

<style lang="scss" module>
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
</style>
