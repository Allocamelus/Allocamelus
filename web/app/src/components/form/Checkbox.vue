<template>
  <div class="mt-3 checkbox-container" @click="toggleCheck()">
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

<style src="./Checkbox.scss" lang="scss" scoped>
</style>