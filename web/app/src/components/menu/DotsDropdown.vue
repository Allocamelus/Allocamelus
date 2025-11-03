<template>
  <div class="relative shrink-0">
    <circle-bg class="hover:bg-rose-800" @click="show = !show">
      <EllipsisVerticalIcon
        class="h-4.5 w-4.5 text-neutral-800 group-hover:text-rose-700 dark:text-neutral-200"
      ></EllipsisVerticalIcon>
    </circle-bg>
    <dropdown v-model="show" class="w-44">
      <slot></slot>
    </dropdown>
  </div>
</template>

<script>
import { defineComponent, reactive, toRefs } from "vue";

import { EllipsisVerticalIcon } from "@heroicons/vue/20/solid";
import CircleBg from "../button/CircleBg.vue";
import Dropdown from "./Dropdown.vue";

export default defineComponent({
  setup() {
    const data = reactive({
      show: false,
    });

    return {
      ...toRefs(data),
    };
  },
  components: { EllipsisVerticalIcon, CircleBg, Dropdown },
});
</script>
