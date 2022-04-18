<template>
  <text-input
    v-model="email"
    :watch-model="watchModel"
    :check="checkC"
    :required="required"
    :max-len="254"
    type="email"
    name="email"
    :regex="/.+@.+\..+/"
    regex-msg="Invalid Email"
    placeholder="mary@example.com"
    @input="emitter"
    @error="$emit('error', $event)"
  >
    <slot></slot>
  </text-input>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import TextInput from "./TextInput.vue";

export default defineComponent({
  name: "EmailInput",
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
  computed: {
    checkC() {
      if (this.required) {
        return true;
      }
      return this.check;
    },
  },
  watch: {
    modelValue(newValue) {
      if (this.watchModel) {
        this.email = newValue;
      }
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
