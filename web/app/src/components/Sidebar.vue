<template>
  <div class="hidden ml-6 w-64 items-stretch md:flex flex-col">
    <box v-if="loggedIn" class="py-3 px-4 rounded-xl">
      <div class="w-max mx-auto">
        <basic-btn
          to="/post/new"
          class="text-white bg-secondary-700 hover:bg-secondary-800 py-2 px-3.5"
        >
          New Post
        </basic-btn>
      </div>
    </box>
    <box v-else class="py-2 px-3 rounded-xl">
      <div class="flex w-max mx-auto">
        <basic-btn
          to="/signup"
          class="text-white bg-secondary-700 hover:bg-secondary-800 py-2 px-3.5"
        >
          Sign Up
        </basic-btn>
        <basic-btn to="/login" class="ml-1.5 py-2 px-3 link"> Login </basic-btn>
      </div>
    </box>
    <slot></slot>
  </div>
</template>

<script>
import { defineComponent, computed, toRefs, reactive } from "vue";
import { useStore } from "vuex";

import Box from "./box/Box.vue";
import BasicBtn from "./button/BasicBtn.vue";

export default defineComponent({
  setup(props) {
    const store = useStore();
    return {
      loggedIn: computed(() => store.getters.loggedIn),
    };
  },
  components: { Box, BasicBtn },
});
</script>