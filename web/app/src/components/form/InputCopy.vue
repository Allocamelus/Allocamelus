<template>
  <text-input
    v-model="value"
    :watchModel="watchModel"
    :readonly="true"
    class="py-0.5"
    ref="input"
  >
    <div class="flex items-center mr-0.5">
      <circle-bg class="p-1.5" title="Copy" @click="doCopy">
        <clipboard-list-outline
          class="text-current group-hover:text-rose-700"
          sizeClass="w-5 h-5"
        ></clipboard-list-outline>
      </circle-bg>
    </div>
  </text-input>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

import ClipboardListOutline from "../icon/ClipboardListOutline.vue";
import TextInput from "./TextInput.vue";
import CircleBg from "../button/CircleBg.vue";

export default defineComponent({
  name: "button-copy",
  props: {
    modelValue: String,
    watchModel: {
      type: Boolean,
      default: false,
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
    modelValue(newValue, old) {
      if (this.watchModel) {
        this.text = newValue;
      }
    },
  },
  methods: {
    doCopy() {
      this.$refs.input.$refs.input.select();
      document.execCommand("copy");
    },
  },
  components: { ClipboardListOutline, TextInput, CircleBg },
});
</script>