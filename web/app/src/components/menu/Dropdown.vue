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
      class="absolute right-0 flex origin-top-right flex-col pt-0.5"
    >
      <div
        class="fixed top-0 bottom-0 right-0 left-0 z-30"
        @click="toggleShow"
      ></div>
      <div
        class="relative z-40 w-full overflow-hidden rounded-md shadow-lg focus:outline-none"
      >
        <slot></slot>
      </div></div
  ></transition>
</template>

<script>
import { defineComponent, reactive, toRefs } from "vue";
export default defineComponent({
  name: "Dropdown",
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
