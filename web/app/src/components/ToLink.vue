<template>
  <component
    :is="local ? 'router-link' : 'a'"
    v-bind="attributes"
    @click="userEvent"
  >
    <slot></slot>
  </component>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { RouteLocationRaw } from "vue-router";
import { useStateStore } from "@/store";

export default defineComponent({
  props: {
    to: {
      type: [Object, String] as PropType<RouteLocationRaw | String>,
      default: "",
    },
  },
  setup() {
    const state = useStateStore();

    return {
      updateViewKey: () => state.updateViewKey(),
    };
  },
  methods: {
    userEvent() {
      if (
        this.to == this.$route.path ||
        (typeof this.to !== "string" &&
          "path" in this.to &&
          this.to.path == this.$route.path)
      ) {
        this.updateViewKey();
      }
    },
  },
  computed: {
    local() {
      if (typeof this.to !== "string") {
        return true;
      }
      try {
        let url = new URL(this.to, window.location.origin);
        if (url.host === window.location.host) {
          return true;
        }
      } catch (_) {
        return true;
      }
      return false;
    },
    attributes() {
      if (this.local) {
        return {
          ["to"]: this.to,
        };
      }
      return {
        ["href"]: this.to,
        ["rel"]: "noopener noreferrer",
        ["target"]: "_blank",
      };
    },
  },
});
</script>
