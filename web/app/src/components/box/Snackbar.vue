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
      class="fixed bottom-0 left-0 right-0 z-auto w-full xs:bottom-4"
    >
      <div class="m-auto w-full xs:w-max">
        <box
          class="flex items-center justify-between rounded-t-lg xs:rounded-md"
          :class="closeBtn ? 'px-3 py-2' : 'px2.5 py-1.5 xs:px-4 xs:py-3'"
        >
          <slot></slot>
          <basic-btn v-if="closeBtn" class="ml-1.5" @click="close">
            <XMarkIcon class="h-5 w-5"></XMarkIcon>
          </basic-btn>
        </box>
      </div>
    </div>
  </transition>
</template>

<script>
import { defineComponent, reactive, toRefs } from "vue";
import { XMarkIcon } from "@heroicons/vue/20/solid";
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
    XMarkIcon,
    BasicBtn,
  },
});
</script>

<style></style>
