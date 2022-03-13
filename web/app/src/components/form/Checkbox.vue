<template>
  <div
    class="flex items-center pr-1 cursor-pointer select-none min-w-max"
    @click="toggleCheck()"
  >
    <component
      :is="checked ? 'radix-checkbox' : 'radix-box'"
      class="w-4 h-4"
    ></component>
    <input type="checkbox" v-model="checked" :name="name" @click.capture.stop />
    <slot>Checkbox</slot>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

import RadixBox from "../icons/RadixBox.vue";
import RadixCheckbox from "../icons/RadixCheckbox.vue";

export default defineComponent({
  props: {
    modelValue: Boolean,
    name: {
      type: String,
      default: "checkbox",
    },
  },
  emits: ["update:modelValue"],
  setup(props) {
    const data = reactive({
      checked: props.modelValue,
    });
    return {
      ...toRefs(data),
    };
  },
  methods: {
    toggleCheck() {
      this.checked = !this.checked;
      this.$emit("update:modelValue", this.checked);
    },
  },
  components: { RadixBox, RadixCheckbox },
});
</script>

<style lang="scss" scoped>
@tailwind components;

@mixin before($content) {
  &::before {
    content: $content;
  }
}
input,
:slotted(label),
:slotted(div) {
  @apply cursor-pointer select-none;
}
input[type="checkbox"] {
  @apply appearance-none font-normal mr-1;
  @apply dark:text-warm-gray-50 focus:outline-none;
  &:checked {
    @apply appearance-none;
  }
}
</style>
