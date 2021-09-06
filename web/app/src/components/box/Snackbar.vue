<template>
  <transition
    :duration="{ enter: 50, leave: 50 }"
    enter-from-class="transform opacity-0 translate-y-0 xs:translate-y-0"
    enter-active-class="transition duration-50"
    enter-to-class="transform opacity-100 translate-y-11 xs:translate-y-4"
    leave-from-class="transform opacity-100 translate-y-0 xs:translate-y-4"
    leave-active-class="transition duration-50"
    leave-to-class="transform opacity-0 translate-y-11 xs:translate-y-0"
  >
    <div
      v-show="modelValue"
      class="w-full fixed z-auto right-0 left-0 bottom-0 xs:bottom-4"
    >
      <div class="m-auto w-full xs:w-max">
        <box
          class="rounded-t-lg xs:rounded-md flex items-center justify-between"
          :class="closeBtn ? 'px-3 py-2' : 'py-1.5 px2.5 xs:py-3 xs:px-4'"
        >
          <slot></slot>
          <basic-btn class="ml-1.5" v-if="closeBtn" @click="close">
            <XIcon class="w-5 h-5"></XIcon>
          </basic-btn>
        </box>
      </div>
    </div>
  </transition>
</template>

<script>
import { defineComponent, reactive, toRefs } from "vue";
import XIcon from "@heroicons/vue/solid/XIcon";
import Box from "./Box.vue";
import BasicBtn from "../button/BasicBtn.vue";

export default defineComponent({
  props: {
    modelValue: Boolean,
    closeBtn: {
      type: Boolean,
      default: false,
    },
    closeTime: {
      type: Number,
      default: 4500, // 4.5sec
    },
  },
  emits: ["update:modelValue"],
  setup(props) {
    const data = reactive({
      show: props.modelValue,
    });

    return {
      ...toRefs(data),
    };
  },
  watch: {
    modelValue(newValue) {
      this.show = newValue;
      if (this.show) {
        setTimeout(() => {
          this.close();
        }, this.closeTime);
      }
    },
  },
  methods: {
    close() {
      this.show = false;
      this.$emit("update:modelValue", false);
    },
  },
  components: {
    Box,
    XIcon,
    BasicBtn,
  },
});
</script>

<style>
</style>