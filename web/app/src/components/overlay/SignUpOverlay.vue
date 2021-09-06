<template>
  <overlay v-model="visible">
    <box
      class="
        w-full
        xs-max:h-full
        xs:m-3
        rounded-none
        xs:rounded-md
        shadow-lg
        bg-opacity-75
        dark:bg-opacity-75
        focus:outline-none
        overflow-hidden
        flex flex-col
      "
    >
      <div class="w-full p-3 border-b border-secondary-600 flex items-end">
        <div
          class="
            flex-1 flex
            justify-end
            text-black
            dark:text-gray-100
            hover:text-gray-600
            dark:hover:text-gray-300
          "
        >
          <basic-btn @click="visible = false">
            <XIcon class="w-5 h-5"></XIcon>
          </basic-btn>
        </div>
      </div>
      <div class="flex-grow flex items-center justify-center">
        <div class="text-center flex flex-col py-8 px-6 xs:px-8">
          <div
            class="
              text-xl
              font-medium
              flex flex-wrap
              items-center
              justify-evenly
            "
          >
            <div>Sign Up or Login to Follow {{ user.name }}</div>
            <div class="pl-1 font-normal text-gray-700 dark:text-gray-400">
              @{{ user.userName }}
            </div>
          </div>
          <div class="flex flex-col items-center mt-8">
            <basic-btn
              to="/signup"
              class="
                w-full
                text-white
                bg-secondary-700
                hover:bg-secondary-800
                py-2.5
                px-3.5
                mb-4
              "
            >
              Sign Up
            </basic-btn>
            <basic-btn
              to="/login"
              class="
                w-full
                py-2.5
                px-3
                link
                border border-rose-800
                dark:border-rose-500
                hover:border-rose-900
                dark:hover:border-rose-600
              "
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

import { GEN_User } from "../../models/go_structs_gen";

import XIcon from "@heroicons/vue/solid/XIcon";

import Box from "../box/Box.vue";
import BasicBtn from "../button/BasicBtn.vue";
import Overlay from "./Overlay.vue";

export default defineComponent({
  props: {
    user: {
      type: GEN_User,
      required: true,
    },
    show: {
      type: Boolean,
      default: false,
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