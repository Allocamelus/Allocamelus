<template>
  <box class="mb-3 rounded-xl p-3.5">
    <text-input
      placeholder="New Post"
      :readonly="true"
      @click="$router.push('/post/new')"
    >
    </text-input>
  </box>
</template>

<script>
import { defineComponent } from "vue";
import Box from "../box/Box.vue";
import TextInput from "../form/TextInput.vue";
export default defineComponent({
  components: { Box, TextInput },
});
</script>
