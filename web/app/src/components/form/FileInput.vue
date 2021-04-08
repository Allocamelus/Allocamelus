<template>
  <div class="">
    <input
      @input="emiter"
      type="file"
      :name="name"
      :accept="accept"
      :required="required"
      :disabled="disabled"
      :readonly="readonly"
      ref="input"
      class="hidden"
      @change="emiter"
      hidden
    />
    <div @click="onClick">
      <slot></slot>
    </div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import * as Errs from "./errors";
import FormatBytes from "../../pkg/fmtBtyes";

export default defineComponent({
  props: {
    name: {
      type: String,
      default: "text",
    },
    accept: {
      type: String,
      required: true,
    },
    check: {
      type: Boolean,
      default: false,
    },
    maxSize: {
      type: Number,
      default: 5242880, // 5MB
    },
    required: {
      type: Boolean,
      default: false,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    readonly: {
      type: Boolean,
      default: false,
    },
  },
  emits: ["fileChange", "error"],
  setup(props) {
    const data = reactive({
      file: null,
    });
    return {
      ...toRefs(data),
    };
  },
  watch: {
    "$refs.input": (newInput, old) => {
      console.log(newInput);
    },
  },
  methods: {
    onClick() {
      this.file = this.$refs.input.value = null;
      this.$refs.input.click();
    },
    validate() {
      if (this.check) {
        if (this.required && this.file?.size <= 0) {
          return Errs.ErrMsg(Errs.ErrRequired);
        }
        if (this.file?.size > this.maxSize) {
          return Errs.ErrMsg(
            Errs.ErrMaxLength,
            FormatBytes(this.maxSize),
            "upload size"
          );
        }
      }
      return "";
    },
    emiter(event) {
      var file = event.target.files[0];
      if (this.file != file) {
        this.file = file;
        if (this.check) {
          this.$emit("error", this.validate());
        }
        this.$emit("fileChange", this.file);
      }
    },
  },
});
</script>
