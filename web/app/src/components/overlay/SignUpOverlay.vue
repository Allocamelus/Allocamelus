<template>
  <overlay v-model="visible">
    <box
      class="flex flex-col w-full h-full overflow-hidden bg-opacity-75 rounded-none shadow-lg xs:h-fit xs:m-3 xs:rounded-md dark:bg-opacity-75 focus:outline-none"
    >
      <div class="flex items-end w-full p-3 border-b border-secondary-600">
        <div class="flex justify-end flex-1" :class="iconStyle.xIcon">
          <basic-btn @click="visible = false">
            <XIcon class="w-5 h-5"></XIcon>
          </basic-btn>
        </div>
      </div>
      <div class="flex items-center justify-center flex-grow">
        <div class="flex flex-col px-6 py-8 text-center xs:px-8">
          <div
            class="flex flex-wrap items-center text-xl font-medium justify-evenly"
          >
            <slot></slot>
          </div>
          <div class="flex flex-col items-center mt-8">
            <basic-btn
              to="/signup"
              class="w-full py-2.5 px-3.5 mb-4"
              :class="buttonStyle.secondary"
            >
              Sign Up
            </basic-btn>
            <basic-btn
              :to="{ path: '/login', query: { r: redirect } }"
              class="w-full py-2.5 px-3 link border border-rose-800 dark:border-rose-500 hover:border-rose-900 dark:hover:border-rose-600"
            >
              Login
            </basic-btn>
          </div>
        </div>
      </div>
    </box>
  </overlay>
</template>

<script>
import { defineComponent, reactive, toRefs } from "vue";

import XIcon from "@heroicons/vue/solid/XIcon";

import Box from "../box/Box.vue";
import BasicBtn from "../button/BasicBtn.vue";
import Overlay from "./Overlay.vue";

export default defineComponent({
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    // Login only redirect
    redirect: {
      type: String,
      default: "",
    },
  },
  emits: ["close"],
  setup(props) {
    // visible is used for overlay instead of show to keep parent and overlay in sync
    const data = reactive({
      visible: props.show,
    });

    return {
      ...toRefs(data),
    };
  },
  watch: {
    show(newValue) {
      this.visible = newValue;
    },
    visible(newValue) {
      if (!newValue) {
        this.close();
      }
    },
  },
  methods: {
    close() {
      this.$emit("close");
    },
  },
  components: {
    XIcon,
    Box,
    BasicBtn,
    Overlay,
  },
});
</script>

<style
  src="@/scss/modules/button.modules.scss"
  lang="scss"
  module="buttonStyle"
  scoped
></style>
<style
  src="@/scss/modules/icon.modules.scss"
  lang="scss"
  module="iconStyle"
  scoped
></style>
