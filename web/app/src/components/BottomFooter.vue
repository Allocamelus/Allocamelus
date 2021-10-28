<template>
  <footer id="footer">
    <div>
      <div class="copyright">&copy; {{ new Date().getFullYear() }}</div>
      <to-link to="/about" class="dash">About</to-link>
    </div>
    <div></div>
    <div>
      <!-- TODO -->
      <to-link to="/tos">Terms</to-link>
      <to-link to="/privacy" class="dash">Privacy</to-link>
    </div>
  </footer>
</template>

<script>
import { defineComponent } from "vue";
import ToLink from "./ToLink.vue";
export default defineComponent({
  components: { ToLink },
});
</script>

<style lang="scss" scoped>
@tailwind components;

@layer components {
  #footer {
    @apply flex justify-between items-center py-3 px-5;
    @apply bg-footer text-white h-footer;
    > div {
      @apply flex items-center;
    }
    &,
    a {
      @apply text-sm text-white;
    }
    a {
      @apply font-semibold no-underline cursor-pointer;
    }
    .dash {
      &::before {
        @apply mx-1;
        content: "-";
      }
    }
    .copyright {
      @apply font-semibold self-center;
    }
  }
}
</style>
