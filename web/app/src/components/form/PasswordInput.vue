<template>
  <text-input
    v-model="password"
    :watchModel="watchModel"
    :check="checkC"
    @input="emiter"
    @error="$emit('error', $event)"
    :type="show ? 'text' : 'password'"
    :name="password"
    :minLen="minLenC"
    :maxLen="maxLenC"
    :required="required"
  >
    <div class="flex items-center mr-1.5">
      <div
        v-if="check"
        class="justify-between mr-1.5 grid grid-rows-2 grid-flow-col gap-0.5"
        :class="strengthClass"
        title="Password Strength"
      >
        <div class="ps-4"></div>
        <div class="ps-3"></div>
        <div class="ps-1"></div>
        <div class="ps-2"></div>
      </div>
      <component
        class="cursor-pointer text-secondary-600 hover:text-secondary-700"
        @click="togglePass"
        :is="show ? 'eye-off-sm' : 'eye-sm'"
        title="Toggle Visibility"
      ></component>
    </div>
  </text-input>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { debounce } from "debounce";
import TextInput from "./TextInput.vue";
import EyeSm from "../icon/EyeSm.vue";
import EyeOffSm from "../icon/EyeOffSm.vue";

export default defineComponent({
  name: "password-input",
  props: {
    modelValue: String,
    watchModel: {
      type: Boolean,
      default: false,
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
    modelValue(newValue, old) {
      if (this.watchModel) {
        this.text = newValue;
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
  methods: {
    scoreDeb() {
      this.score = this.zxcvbn(this.password.substring(0, 64)).score;
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
    EyeSm,
    EyeOffSm,
  },
});
</script>

<style lang="scss" scoped>
@layer components {
  .ps-1,
  .ps-2,
  .ps-3,
  .ps-4 {
    @apply h-1.5 w-1.5 bg-gray-500;
  }
  .s1 {
    .ps-1 {
      @apply bg-red-600;
    }
  }
  .s2 {
    .ps-1,
    .ps-2 {
      @apply bg-yellow-400;
    }
  }
  .s3 {
    div {
      @apply bg-orange-600;
    }
    .ps-4 {
      @apply bg-gray-500;
    }
  }
  .s4 {
    div {
      @apply bg-green-600;
    }
  }
}
</style>