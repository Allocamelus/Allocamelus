<template>
  <text-input
    v-model="password"
    :check="checkC"
    @input="emiter"
    @error="$emit('error', $event)"
    :type="show ? 'text' : 'password'"
    :name="password"
    :minLen="minLenC"
    :maxLen="maxLenC"
    :required="required"
  >
    <div class="iconPassBox">
      <div
        v-if="check"
        class="iconPassSec"
        :class="strengthClass"
        title="Password Strength"
      ></div>
      <i
        class="far iconPass"
        v-on:click="togglePass()"
        :class="show ? 'fa-eye-slash' : 'fa-eye'"
        title="Toggle Visibility"
      ></i></div
  ></text-input>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { debounce } from "debounce";
import TextInput from "./TextInput.vue";

export default defineComponent({
  name: "password-input",
  props: {
    modelValue: String,
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
      zxcvbn: undefined,
    });

    return {
      ...toRefs(data),
    };
  },
  watch: {
    password(newPass, oldPass) {
      if (this.check) {
        this.debouncedCheck();
      }
    },
  },
  created() {
    var vm = this;
    vm.password = vm.modelValue;
    if (vm.check) {
      import("zxcvbn").then((zxcvbn) => {
        vm.zxcvbn = zxcvbn.default;
      });

      vm.debouncedCheck = debounce(vm.scoreDeb, 200, true);
    }
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
          return ""
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
  methods: {
    scoreDeb() {
      this.score = this.zxcvbn(this.password.substring(0,64)).score;
    },
    togglePass() {
      this.show = !this.show;
    },
    emiter() {
      this.$emit("update:modelValue", this.password);
    },
  },
  components: {
    TextInput,
  },
});
</script>

<style src="./PasswordInput.scss" lang="scss" scoped>
</style>