<template>
  <transition
    :duration="{ enter: 25, leave: 25 }"
    enter-from-class="transform opacity-0 scale-90"
    enter-active-class="transition ease-out duration-25"
    enter-to-class="transform opacity-100 scale-100"
    leave-from-class="transform opacity-100 scale-100"
    leave-active-class="transition ease-in duration-25"
    leave-to-class="transform opacity-0 scale-90"
  >
    <div
      v-show="show"
      class="origin-top-right absolute flex flex-col right-0 pt-0.5"
    >
      <div
        class="top-0 bottom-0 right-0 left-0 fixed z-30"
        @click="toggleShow"
      ></div>
      <div
        class="
          w-full
          z-40
          relative
          rounded-md
          shadow-lg
          focus:outline-none
          overflow-hidden
        "
      >
        <slot></slot>
      </div></div
  ></transition>
</template>

<script>
import { defineComponent, reactive, toRefs } from "vue";
export default defineComponent({
  name: "dropdown",
  props: {
    modelValue: Boolean,
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
    },
  },
  methods: {
    toggleShow() {
      this.show = !this.show;
      this.$emit("update:modelValue", this.show);
    },
  },
});
</script>
