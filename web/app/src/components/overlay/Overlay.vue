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
          xsFullHeigth ? 'h-full xs:h-auto' : '',
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
import Box from "../box/Box.vue";
export default defineComponent({
  components: { Box },
  name: "overlay",
  props: {
    modelValue: Boolean,
    blockScrool: {
      type: Boolean,
      default: true,
    },
    xsFullHeigth: {
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
    if (this.blockScrool) {
      // Creating invisible container
      const outer = document.createElement("div");
      outer.style.visibility = "hidden";
      outer.style.overflow = "scroll"; // forcing scrollbar to appear
      outer.style.msOverflowStyle = "scrollbar"; // needed for WinJS apps
      document.body.appendChild(outer);

      // Creating inner element and placing it in the container
      const inner = document.createElement("div");
      outer.appendChild(inner);

      // Calculating difference between container's full width and the child width
      this.offsetWidth = outer.offsetWidth - inner.offsetWidth;

      // Removing temporary elements from the DOM
      outer.parentNode.removeChild(outer);
    }
  },
  watch: {
    modelValue(newValue, old) {
      this.show = newValue;
      var h = document.querySelector("html"),
        navS = document.querySelector("#nav").style;
      if (this.blockScrool) {
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