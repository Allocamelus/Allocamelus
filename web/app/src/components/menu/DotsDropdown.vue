<template>
  <div class="flex-shrink-0 relative">
    <circle-bg class="hover:bg-rose-800" @click="show = !show">
      <DotsVerticalIcon
        class="
          h-4.5
          w-4.5
          text-gray-800
          dark:text-gray-200
          group-hover:text-rose-700
        "
      ></DotsVerticalIcon>
    </circle-bg>
    <dropdown v-model="show" class="w-44">
        <slot></slot>
    </dropdown>
  </div>
</template>

<script>
import { defineComponent, reactive, toRefs } from "vue";

import DotsVerticalIcon from "@heroicons/vue/outline/DotsVerticalIcon";
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
  components: { DotsVerticalIcon, CircleBg, Dropdown },
});
</script>