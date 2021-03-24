<template>
  <div class="min-w-max pr-1 cursor-pointer select-none" @click="toggleCheck()">
    <input type="checkbox" v-model="checked" :name="name" @click.capture.stop />
    <slot>Checkbox</slot>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
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
});
</script>

<style lang="scss" scoped>
@mixin before($content) {
  &::before {
    content: $content;
  }
}
@layer components {
  input,
  :slotted(label),
  :slotted(div) {
    @apply cursor-pointer select-none;
  }
  input[type="checkbox"] {
    @apply font-awesome appearance-none font-normal mr-1;
    @apply dark:text-warm-gray-50 focus:outline-none;
    font-size: 1.01em;
    @include before("\f0c8");
    &:checked {
      @include before("\f14a");
    }
  }
}
</style>