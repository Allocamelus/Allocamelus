<template>
  <div
    v-show="show"
    class="top-0 bottom-0 right-0 left-0 fixed w-full h-full z-30"
  >
    <div
      class="flex justify-center items-center m-auto h-full w-full xs:max-w-md"
    >
      <div
        class="top-0 bottom-0 right-0 left-0 opacity-75 bg-black fixed"
        @click="toggleShow"
      ></div>
      <div
        class="relative flex items-center justify-center w-full"
        :class="[
          xsFullHeight ? 'h-full xs:h-auto' : '',
          xsSelfEnd ? 'self-end xs:self-center' : '',
        ]"
      >
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive, toRefs } from "vue";
export default defineComponent({
  name: "overlay",
  props: {
    modelValue: Boolean,
    blockScroll: {
      type: Boolean,
      default: true,
    },
    xsFullHeight: {
      type: Boolean,
      default: true,
    },
    xsSelfEnd: {
      type: Boolean,
      default: false,
    },
  },
  emits: ["update:modelValue"],
  setup(props) {
    const data = reactive({
      show: props.modelValue,
      offsetWidth: 0,
    });

    return {
      ...toRefs(data),
    };
  },
  beforeMount() {
    if (this.blockScroll) {
      // Get scrollbar width
      this.offsetWidth = window.innerWidth - document.querySelector('html').clientWidth;
    }
  },
  watch: {
    modelValue(newValue) {
      this.show = newValue;
      var h = document.querySelector("html"),
        navS = document.querySelector("#nav").style;
      if (this.blockScroll) {
        if (this.show) {
          h.classList.add("overflow-hidden");
          h.style.marginRight = navS.paddingRight = `${this.offsetWidth}px`;
        } else if (h.classList.contains("overflow-hidden")) {
          h.classList.remove("overflow-hidden");
          h.style.marginRight = navS.paddingRight = null;
        }
      }
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