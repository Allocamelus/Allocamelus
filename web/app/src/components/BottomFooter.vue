<template>
  <footer id="footer" class="bg-footer text-white xs:h-footer">
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
        <to-link to="https://github.com/Allocamelus/Allocamelus">
          <radix-github class="h-5 w-5" />
        </to-link>
        <to-link to="https://www.allocamelus.com/u/Allocamelus" class="ml-3">
          <allocamelus class="h-5 w-5" />
        </to-link>
      </div>
      <div class="flex">
        <!-- TODO -->
        <to-link to="/terms">Terms</to-link>
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
  @apply flex items-center justify-between py-3;
  > div {
    @apply items-center;
  }
  &,
  a {
    @apply text-sm text-white;
  }
  a {
    @apply cursor-pointer font-semibold no-underline;
  }
}
</style>
