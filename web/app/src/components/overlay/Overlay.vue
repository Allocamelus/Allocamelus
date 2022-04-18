<template>
  <div
    v-show="show"
    class="fixed top-0 bottom-0 right-0 left-0 z-30 h-full w-full"
  >
    <div
      class="m-auto flex h-full w-full items-center justify-center xs:max-w-md"
    >
      <div
        class="fixed top-0 bottom-0 right-0 left-0 bg-black opacity-75"
        @click="toggleShow"
      ></div>
      <div
        class="relative flex w-full items-center justify-center"
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
  name: "Overlay",
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
  beforeUnmount() {
    unblock();
  },
  methods: {
    toggleShow() {
      this.show = !this.show;
      this.$emit("update:modelValue", this.show);
    },
  },
});
</script>
