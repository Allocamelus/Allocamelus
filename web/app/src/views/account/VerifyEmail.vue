<template>
  <div>
    <div v-if="!hasSelectorToken">
    </div>
    Test
    <div>{{ selector }}</div>
    <div>{{ token }}</div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive, computed } from "vue";
export default defineComponent({
  props: {
    selector: {
      type: String,
      defualt: "",
    },
    token: {
      type: String,
      defualt: "",
    },
  },
  setup(props) {
    const data = reactive({});

    const hasSelectorToken = computed(() => {
      return props.selector != undefined && props.token != undefined
        ? props.selector.length != 0 && props.token.length != 0
          ? true
          : false
        : false;
    });

    return {
      ...toRefs(data),
      hasSelectorToken,
    };
  },
});
</script>

<style>
</style>