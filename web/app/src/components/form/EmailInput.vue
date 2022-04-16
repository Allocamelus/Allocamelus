<template>
  <text-input
    v-model="email"
    :watchModel="watchModel"
    :check="checkC"
    :required="required"
    :maxLen="254"
    @input="emitter"
    type="email"
    name="email"
    :regex="/.+@.+\..+/"
    regexMsg="Invalid Email"
    placeholder="mary@example.com"
    @error="$emit('error', $event)"
  >
    <slot></slot>
  </text-input>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import TextInput from "./TextInput.vue";

export default defineComponent({
  name: "email-input",
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
    required: {
      type: Boolean,
      default: true,
    },
  },
  emits: ["update:modelValue", "error"],
  setup(props) {
    const data = reactive({
      email: props.modelValue,
    });

    return {
      ...toRefs(data),
    };
  },
  watch: {
    modelValue(newValue) {
      if (this.watchModel) {
        this.email = newValue;
      }
    },
  },
  computed: {
    checkC() {
      if (this.required) {
        return true;
      }
      return this.check;
    },
  },
  methods: {
    emitter() {
      this.$emit("update:modelValue", this.email);
    },
  },
  components: {
    TextInput,
  },
});
</script>
