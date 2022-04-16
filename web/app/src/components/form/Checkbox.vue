<template>
  <div
    class="flex min-w-max cursor-pointer select-none items-center pr-1"
    @click="toggleCheck()"
  >
    <component
      :is="checked ? 'radix-checkbox' : 'radix-box'"
      class="h-4 w-4"
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
  @apply mr-1 appearance-none font-normal;
  @apply focus:outline-none dark:text-warm-gray-50;
  &:checked {
    @apply appearance-none;
  }
}
</style>
