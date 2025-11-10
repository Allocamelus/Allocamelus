<template>
  <overlay v-model="visible">
    <box
      class="bg-opacity-75 dark:bg-opacity-75 flex h-full w-full flex-col overflow-hidden rounded-none shadow-lg focus:outline-none xs:m-3 xs:h-fit xs:rounded-md"
    >
      <div class="flex w-full items-end border-b border-secondary-600 p-3">
        <div class="xIcon flex flex-1 justify-end">
          <basic-btn @click="visible = false">
            <XMarkIcon class="h-5 w-5"></XMarkIcon>
          </basic-btn>
        </div>
      </div>
      <div class="flex grow items-center justify-center">
        <div class="flex flex-col px-6 py-8 text-center xs:px-8">
          <div
            class="flex flex-wrap items-center justify-evenly text-xl font-medium"
          >
            <slot></slot>
          </div>
          <div class="mt-8 flex flex-col items-center">
            <basic-btn
              to="/signup"
              class="btn-secondary mb-4 w-full px-3.5 py-2.5"
            >
              Sign Up
            </basic-btn>
            <basic-btn
              :to="{ path: '/login', query: { r: redirect } }"
              class="link w-full border border-rose-800 px-3 py-2.5 hover:border-rose-900 dark:border-rose-500 dark:hover:border-rose-600"
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

import { XMarkIcon } from "@heroicons/vue/20/solid";

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
    XMarkIcon,
    Box,
    BasicBtn,
    Overlay,
  },
});
</script>
