<template>
  <div class="input-container" :class="focused ? 'focus' : ''">
    <input
      v-model.trim="text"
      @input="emiter"
      :type="type"
      :name="name"
      :minlength="minLenC"
      :maxlength="maxLenC+1"
      class="input"
      @focus="focused = true"
      @blur="focused = false"
      :required="requiredC"
    />
    <slot></slot>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import * as Errs from "./errors";

export default defineComponent({
  props: {
    modelValue: String,
    name: {
      type: String,
      default: "text",
    },
    type: {
      type: String,
      default: "text",
    },
    check: {
      type: Boolean,
      default: false,
    },
    minLen: {
      type: Number,
      default: 0,
    },
    maxLen: {
      type: Number,
      default: 16383,
    },
    required: {
      type: Boolean,
      default: false,
    },
    regex: RegExp,
    regexMsg: String,
  },
  emits: ["update:modelValue", "error"],
  setup(props) {
    const data = reactive({
      text: props.modelValue,
      focused: false,
    });
    return {
      ...toRefs(data),
    };
  },
  computed: {
    minLenC() {
      if (this.check) {
        return this.minLen;
      }
      return 0;
    },
    maxLenC() {
      if (this.check) {
        return this.maxLen;
      }
      return 16383;
    },
    requiredC() {
      if (this.check) {
        return this.required;
      }
      return false;
    },
  },
  methods: {
    validate() {
      var vm = this,
        l = vm.text.length;

      if (vm.required && l <= 0) {
        return Errs.ErrMsg(Errs.ErrRequired);
      }

      if (l < vm.minLenC) {
        return Errs.ErrMsg(Errs.ErrMinLength, vm.minLenC);
      }

      if (l > vm.maxLenC) {
        return Errs.ErrMsg(Errs.ErrMaxLength, vm.maxLenC);
      }

      if (vm.regex != undefined) {
        if (!vm.regex.test(String(vm.text))) {
          return Errs.ErrMsg(Errs.ErrRegex, vm.regexMsg);
        }
      }
      return "";
    },
    emiter() {
      this.$emit("update:modelValue", this.text);
      if (this.check) {
        this.$emit("error", this.validate());
      }
    },
  },
});
</script>

<style src="./input.scss" lang="scss" scoped>
</style>