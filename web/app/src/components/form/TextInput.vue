<template>
  <div
    class="my-1 box-border flex w-full items-center justify-between rounded-sm border border-solid border-stone-400 bg-neutral-300 p-0 text-black-lighter focus-within:border-secondary-600 xs:bg-neutral-200 dark:border-stone-500 dark:bg-neutral-800 dark:text-white dark:focus-within:border-secondary-600"
  >
    <input
      ref="input"
      v-model.trim="text"
      :type="type"
      :name="name"
      :minlength="minLenC"
      :maxlength="maxLenC + 1"
      class="placeholder-opacity-80 dark:placeholder-opacity-80 mr-2.5 box-content flex-1 rounded-l-sm border-none bg-transparent py-1.5 pl-2.5 placeholder-stone-800 shadow-none outline-none autofill:text-black-lighter autofill:shadow-[inset_0_0_0_100px] autofill:shadow-neutral-300 hover:autofill:bg-transparent focus:outline-none focus:autofill:bg-transparent active:autofill:bg-transparent xs:autofill:shadow-neutral-200 dark:placeholder-stone-400 dark:autofill:text-white dark:autofill:shadow-neutral-800"
      :required="requiredC"
      :placeholder="placeholder"
      :disabled="disabled"
      :readonly="readonly"
      @input="emitter"
    />
    <slot></slot>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import * as Errs from "./errors";

export default defineComponent({
  props: {
    modelValue: {
      type: String,
      default: "",
    },
    watchModel: {
      type: Boolean,
      default: false,
    },
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
    regexMsg: {
      type: String,
      default: "",
    },
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
  watch: {
    modelValue(newValue) {
      if (this.watchModel) {
        this.text = newValue;
      }
    },
  },
  methods: {
    validate() {
      var l = this.text.length;

      if (this.required && l <= 0) {
        return Errs.ErrMsg(Errs.ErrRequired);
      }

      if (l < this.minLenC) {
        return Errs.ErrMsg(Errs.ErrMinLength, this.minLenC);
      }

      if (l > this.maxLenC) {
        return Errs.ErrMsg(Errs.ErrMaxLength, this.maxLenC);
      }

      if (this.regex != undefined) {
        if (!this.regex.test(String(this.text))) {
          return Errs.ErrMsg(Errs.ErrRegex, this.regexMsg);
        }
      }
      return "";
    },
    emitter() {
      this.$emit("update:modelValue", this.text);
      if (this.check) {
        this.$emit("error", this.validate());
      }
    },
  },
});
</script>
