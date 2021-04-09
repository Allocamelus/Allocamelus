<template>
  <textarea
    v-model.trim="text"
    @input="emiter"
    :type="type"
    :name="name"
    :minlength="minLenC"
    :maxlength="maxLenC + 1"
    class="w-full rounded-sm box-border border border-solid"
    :class="[
      'py-1.5 px-2.5 my-1',
      'bg-gray-200 xs-max:bg-gray-300 border-warm-gray-400 focus:border-secondary-600 text-black-lighter',
      'dark:bg-gray-800 dark:border-warm-gray-500 dark:focus:border-secondary-600 dark:text-white',
      'placeholder-warm-gray-800 placeholder-opacity-80 dark:placeholder-warm-gray-400 dark:placeholder-opacity-80'
    ]"
    :required="requiredC"
    :placeholder="placeholder"
    :disabled="disabled"
    :readonly="readonly"
    ref="input"
  />
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import * as Errs from "./errors";

export default defineComponent({
  props: {
    modelValue: String,
    watchModel: {
      type: Boolean,
      default: false,
    },
    name: {
      type: String,
      default: "textarea",
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
    placeholder: {
      type: String,
      default: "",
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    readonly: {
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
    });
    return {
      ...toRefs(data),
    };
  },
  watch: {
    modelValue(newValue, old) {
      if (this.watchModel) {
        this.text = newValue;
      }
    },
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