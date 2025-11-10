<template>
  <div class="flex flex-col">
    <strong>{{ errorText[0] }}</strong>
    <div v-if="errorText.length > 1">{{ errorText[1] }}</div>
    <slot></slot>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, PropType } from "vue";

function getPath(path = ""): string {
  if (path !== "") {
    return path;
  }
  return window.location.pathname;
}

export const SomethingWentWrong = "SomethingWentWrong";
export const LoadingCaptcha = "LoadingCaptcha";

export default defineComponent({
  props: {
    error: {
      type: [Number, String, Array] as PropType<number | string | string[]>,
      default: "",
    },
    path: {
      type: String,
      default: "",
    },
  },
  setup(props) {
    return {
      errorText: computed(() => {
        let path = getPath(props.path);
        if (typeof props.error == "object") {
          return props.error;
        }
        switch (props.error) {
          case 403:
          case "403":
            return [`Error: 403 Forbidden`, `This page is Private ${path}`];
          case 404:
          case "404":
            return [`Error: 404`, `${path} Not Found`];
          case 422:
          case "422":
            return [`Error: 422 Unprocessable Entity`];
          case 429:
          case "429":
            return [`Error: 429 Too Many Requests`, `Try again later`];
          case LoadingCaptcha:
            return [`Loading captcha...`];
          case SomethingWentWrong:
          default:
            return [`Something went wrong`, `Try again later`];
        }
      }),
    };
  },
});
</script>
