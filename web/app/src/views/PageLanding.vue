<template>
  <div class="container flex py-5">
    <div class="mx-auto grid grid-cols-5 pt-10">
      <div
        class="z-10 col-start-1 col-end-6 row-start-1 flex flex-col sm:col-start-3"
      >
        <h1 class="text-center" :class="textStyle.h1">Allocamelus</h1>
        <h2 class="mt-5 text-center" :class="textStyle.h2">
          A social media platform
        </h2>
        <div class="mt-6 flex justify-center space-x-3">
          <basic-btn
            to="/signup"
            class="px-4 py-2.5 leading-4"
            :class="buttonStyle.secondary"
          >
            Sign Up
          </basic-btn>
          <basic-btn
            to="https://github.com/Allocamelus/Allocamelus"
            class="px-4 py-2.5 text-base leading-4"
            :class="buttonStyle.secondaryBorder"
          >
            <div class="flex items-center">
              <CodeBracketIcon class="mr-1 h-4 w-4" />
              <div>Open Source</div>
            </div>
          </basic-btn>
        </div>
        <div class="mt-3 flex items-baseline justify-center text-sm">
          <div class="text-gray-800 dark:text-gray-200">Have an account</div>
          <div class="ml-1 flex">
            <to-link to="/login" class="link">Login</to-link>
          </div>
        </div>
        <div class="flex"></div>
      </div>
      <div
        class="col-start-1 col-end-6 row-start-2 mt-6 sm:col-end-3 sm:row-start-1 sm:mt-0"
      >
        <div class="flex w-full">
          <allocamelus
            class="h-full w-full object-cover text-primary-lighter opacity-50 dark:opacity-70 sm:opacity-100"
          ></allocamelus>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

import { CodeBracketIcon } from "@heroicons/vue/20/solid";

import Allocamelus from "@/components/icons/Allocamelus.vue";
import BasicBtn from "@/components/button/BasicBtn.vue";
import ToLink from "@/components/ToLink.vue";

export default defineComponent({
  setup() {
    const data = reactive({});

    document.title = import.meta.env.VITE_SITE_NAME;

    return {
      ...toRefs(data),
    };
  },
  components: { Allocamelus, BasicBtn, CodeBracketIcon, ToLink },
});
</script>

<style
  src="@/scss/modules/button.modules.scss"
  lang="scss"
  module="buttonStyle"
  scoped
></style>
<style
  src="@/scss/modules/text.modules.scss"
  lang="scss"
  module="textStyle"
  scoped
></style>
