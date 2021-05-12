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
      :multiple="multiple"
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
      default: 52428800, // 50MB
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
    multiple: {
      type: Boolean,
      default: false,
    },
    maxFiles: {
      type: Number,
      default: 5,
    },
    fileCount: Number,
  },
  emits: ["filesChange", "error"],
  setup() {
    const data = reactive({
      files: null,
    });
    return {
      ...toRefs(data),
    };
  },
  watch: {
    "$refs.input": (newInput, _old) => {
      console.log(newInput);
    },
  },
  methods: {
    onClick() {
      this.files = this.$refs.input.value = null;
      this.$refs.input.click();
    },
    validate() {
      if (this.check) {
        if (this.multiple) {
          var filesL = this.files.length;

          if (this.fileCount + filesL > this.maxFiles) {
            this.files = [];
            return Errs.ErrMsg(Errs.ErrMaxLength, this.maxFiles, "file count");
          }

          var file,
            err = "",
            lastErr = "",
            newFiles = [];

          for (let i = 0; i < filesL; i++) {
            file = this.files[i];
            err = this.validateFile(file);
            if (err.length > 0) {
              lastErr = `${file.name}: ${err}`;
            } else {
              newFiles.push(file);
            }
          }

          this.files = newFiles;
          return lastErr;
        } else {
          return this.validateFile(this.files);
        }
      }
      return "";
    },
    validateFile(file) {
      if (this.required && file?.size <= 0) {
        return Errs.ErrMsg(Errs.ErrRequired);
      }
      if (file?.size > this.maxSize) {
        return Errs.ErrMsg(
          Errs.ErrMaxLength,
          FormatBytes(this.maxSize),
          "upload size"
        );
      }
      return "";
    },
    emiter(event) {
      var filesInput;
      if (this.multiple) {
        filesInput = [];
        for (let i = 0; i < event.target.files.length; i++) {
          filesInput.push(event.target.files[i]);
        }
      } else {
        filesInput = event.target.files[0];
      }

      if (this.files == null && filesInput.length  != 0) {
        this.files = filesInput;
        if (this.check) {
          this.$emit("error", this.validate());
        }
        this.$emit("filesChange", this.files);
      }
    },
  },
});
</script>
