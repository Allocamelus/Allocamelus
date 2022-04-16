<template>
  <component
    :is="isLink ? 'to-link' : 'div'"
    class="flex truncate text-gray-800 dark:text-gray-200"
    :class="[
      isLink ? 'group no-underline' : '',
      twoLine ? 'flex-col' : 'items-center',
    ]"
    :to="'/u/' + user.userName"
  >
    <div
      v-if="!noName"
      class="flex-shrink truncate font-semibold"
      :class="[
        isLink ? 'group-hover:underline' : '',
        twoLine ? 'mb-0.5 text-xl' : '',
      ]"
    >
      {{ user.name }}
    </div>
    <text-small class="flex-none font-normal" :class="!twoLine ? 'ml-1' : ''">
      @{{ user.userName }}
    </text-small>
  </component>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { User } from "../../models/user";

import ToLink from "../ToLink.vue";
import TextSmall from "../text/Small.vue";

export default defineComponent({
  props: {
    user: {
      type: User,
      required: true,
    },
    isLink: {
      type: Boolean,
      default: true,
    },
    noName: {
      type: Boolean,
      default: false,
    },
    twoLine: {
      type: Boolean,
      default: false,
    },
  },
  computed: {},
  components: { ToLink, TextSmall },
});
</script>
