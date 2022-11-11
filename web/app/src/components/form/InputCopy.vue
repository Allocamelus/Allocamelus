<template>
  <text-input
    ref="input"
    v-model="value"
    :watch-model="true"
    :readonly="true"
    class="py-0.5"
  >
    <div class="mr-0.5 flex items-center">
      <circle-bg class="p-1.5 hover:bg-rose-800" title="Copy" @click="doCopy">
        <ClipboardDocumentListIcon
          class="h-5 w-5 text-current group-hover:text-rose-700"
        ></ClipboardDocumentListIcon>
      </circle-bg>
    </div>
  </text-input>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

import { ClipboardDocumentListIcon } from "@heroicons/vue/24/outline";
import TextInput from "./TextInput.vue";
import CircleBg from "../button/CircleBg.vue";

export default defineComponent({
  name: "ButtonCopy",
  props: {
    modelValue: {
      type: String,
      default: "",
    },
  },
  setup(props) {
    const data = reactive({
      value: props.modelValue,
    });

    return {
      ...toRefs(data),
    };
  },
  watch: {
    modelValue(newValue) {
      this.value = newValue;
    },
  },
  methods: {
    doCopy() {
      this.$refs.input.$refs.input.select();
      navigator.clipboard.writeText(this.value);
    },
  },
  components: { ClipboardDocumentListIcon, TextInput, CircleBg },
});
</script>
