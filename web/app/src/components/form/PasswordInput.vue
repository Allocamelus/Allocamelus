<template>
  <text-input
    v-model="password"
    :watch-model="watchModel"
    :check="checkC"
    :type="show ? 'text' : 'password'"
    name="password"
    :min-len="minLenC"
    :max-len="maxLenC"
    :required="required"
    @input="emitter"
    @error="$emit('error', $event)"
  >
    <div class="mr-1.5 flex items-center">
      <div
        v-if="check"
        class="mr-1.5 grid grid-flow-col grid-rows-2 justify-between gap-0.5"
        :class="strengthClass"
        title="Password Strength"
      >
        <div class="pass-4"></div>
        <div class="pass-3"></div>
        <div class="pass-1"></div>
        <div class="pass-2"></div>
      </div>
      <div
        title="Toggle Visibility"
        class="cursor-pointer text-secondary-600 hover:text-secondary-700"
      >
        <component
          :is="show ? 'EyeSlashIcon' : 'EyeIcon'"
          class="w5 h-5"
          @click="togglePass"
        ></component>
      </div>
    </div>
  </text-input>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { debounce } from "throttle-debounce";
import TextInput from "./TextInput.vue";
import { EyeIcon, EyeSlashIcon } from "@heroicons/vue/20/solid";

import { zxcvbn, zxcvbnOptions } from "@zxcvbn-ts/core";
import * as zxcvbnCommonPackage from "@zxcvbn-ts/language-common";
import * as zxcvbnEnPackage from "@zxcvbn-ts/language-en";

export default defineComponent({
  name: "PasswordInput",
  props: {
    modelValue: {
      type: String,
      default: "",
    },
    watchModel: {
      type: Boolean,
      default: true,
    },
    check: {
      type: Boolean,
      default: false,
    },
    minLen: {
      type: Number,
      default: 8,
    },
    maxLen: {
      type: Number,
      default: 1024,
    },
    required: {
      type: Boolean,
      default: true,
    },
  },
  emits: ["update:modelValue", "error"],
  setup(props) {
    const data = reactive({
      password: props.modelValue,
      show: false,
      score: 0,
      debouncedCheck: undefined,
    });

    zxcvbnOptions.setOptions({
      graphs: zxcvbnCommonPackage.adjacencyGraphs,
      useLevenshteinDistance: true,
      dictionary: {
        ...zxcvbnCommonPackage.dictionary,
        ...zxcvbnEnPackage.dictionary,
      },
    });

    return {
      ...toRefs(data),
    };
  },
  computed: {
    checkC() {
      if (this.required) {
        return true;
      }
      return this.check;
    },
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
      return 1024;
    },
    strengthClass() {
      if (this.check) {
        if (this.password.length == 0) {
          return "";
        }
        return {
          s1: this.score == 1,
          s2: this.score == 2,
          s3: this.score == 3,
          s4: this.score == 4,
        };
      }
      return "";
    },
  },
  watch: {
    password() {
      if (this.check) {
        this.debouncedCheck();
      }
    },
    modelValue(newValue) {
      if (this.watchModel) {
        this.password = newValue;
      }
    },
  },
  created() {
    this.password = this.modelValue;
    if (this.check) {
      this.debouncedCheck = debounce(200, this.scoreDeb, { atBegin: true });
    }
  },
  methods: {
    scoreDeb() {
      this.score = zxcvbn(this.password.substring(0, 64)).score;
    },
    togglePass() {
      this.show = !this.show;
    },
    emitter() {
      this.$emit("update:modelValue", this.password);
    },
  },
  components: {
    TextInput,
    EyeIcon,
    EyeSlashIcon,
  },
});
</script>
