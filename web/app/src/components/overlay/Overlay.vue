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

<script lang="ts">
import { defineComponent, reactive, toRefs } from "vue";

function getElements(): { h: HTMLInputElement; navS: CSSStyleDeclaration } {
  return {
    h: document.querySelector<HTMLInputElement>("html")!,
    navS: document.querySelector<HTMLInputElement>("#nav")!.style,
  };
}

function offsetWidth() {
  let { h } = getElements();
  return window.innerWidth - h.clientWidth;
}

function block() {
  let { h, navS } = getElements();
  h.classList.add("overflow-hidden");
  h.style.marginRight = navS.paddingRight = `${offsetWidth()}px`;
}
function unblock() {
  let { h, navS } = getElements();
  if (h.classList.contains("overflow-hidden")) {
    h.classList.remove("overflow-hidden");
    h.style.removeProperty("marginRight");
    navS.removeProperty("paddingRight");
  }
}

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
    });

    return {
      ...toRefs(data),
    };
  },
  beforeUnmount() {
    unblock();
  },
  watch: {
    modelValue(newValue) {
      this.show = newValue;
      if (this.blockScroll) {
        if (this.show) {
          block();
        } else {
          unblock();
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
