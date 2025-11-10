<template>
  <div
    class="flex min-w-max cursor-pointer items-center pr-1 select-none"
    @click="toggleCheck()"
  >
    <component
      :is="checked ? 'radix-checkbox' : 'radix-box'"
      class="h-4 w-4"
    ></component>
    <input
      v-model="checked"
      class="mr-1 cursor-pointer appearance-none font-normal select-none checked:appearance-none focus:outline-none dark:text-stone-50"
      type="checkbox"
      :name="name"
      @click.capture.stop
    />
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
