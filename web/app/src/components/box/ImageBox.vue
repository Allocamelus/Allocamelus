<template>
  <div
    class="group relative flex overflow-hidden"
    :class="[
      totalNumber == 1 ? 'w-full rounded-lg' : '',
      totalNumber == 2
        ? [
            index == 0 || index == 1 ? 'w-1/2' : '',
            index == 0 ? 'rounded-l-lg' : 'rounded-r-lg',
          ]
        : '',
      totalNumber == 3
        ? [
            index == 0
              ? 'w-full rounded-t-lg'
              : ['w-1/2', index == 1 ? 'rounded-bl-lg' : ' rounded-br-lg'],
          ]
        : '',
      totalNumber == 4
        ? [
            'w-1/2',
            index == 0
              ? 'rounded-tl-lg'
              : [
                  index == 1
                    ? 'rounded-tr-lg'
                    : [index == 2 ? 'rounded-bl-lg' : 'rounded-br-lg'],
                ],
          ]
        : '',
    ]"
  >
    <slot></slot>
    <img :src="url" class="w-full object-cover"/>
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    url: String,
    index: Number, // Index
    totalNumber: Number, // How many other images are there
    loading: {
      type: String,
      default: "auto"
    }
  },
});
</script>