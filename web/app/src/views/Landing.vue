<template>
  <div class="container flex py-5">
    <div class="grid grid-cols-5 pt-10 mx-auto">
      <div
        class="z-10 flex flex-col col-start-1 col-end-6 row-start-1 sm:col-start-3"
      >
        <h1 class="text-4xl font-semibold text-center sm:text-5xl">
          Allocamelus
        </h1>
        <h2 class="mt-5 text-2xl font-semibold text-center sm:text-3xl">
          A social media platform
        </h2>
        <div class="flex justify-center mt-6 space-x-3">
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
              <CodeIcon class="w-4 h-4 mr-1" />
              <div>Open Source</div>
            </div>
          </basic-btn>
        </div>
        <div class="flex items-baseline justify-center mt-3 text-sm">
          <div class="text-gray-800 dark:text-gray-200">Have an account</div>
          <div class="flex ml-1">
            <to-link to="/login" class="link">Login</to-link>
          </div>
        </div>
        <div class="flex"></div>
      </div>
      <div
        class="col-start-1 col-end-6 row-start-2 mt-6 sm:mt-0 sm:col-end-3 sm:row-start-1"
      >
        <div class="flex w-full">
          <allocamelus
            class="object-cover w-full h-full opacity-50 dark:opacity-50 sm:opacity-100 text-primary-lighter"
          ></allocamelus>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

import { CodeIcon } from "@heroicons/vue/solid";

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
  components: { Allocamelus, BasicBtn, CodeIcon, ToLink },
});
</script>

<style
  src="@/scss/modules/button.modules.scss"
  lang="scss"
  module="buttonStyle"
  scoped
></style>
