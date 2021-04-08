<template>
  <div class="text-center flex flex-col py-8 px-6 xs:px-8">
    <div class="text-xl font-medium flex">
      <slot></slot>
    </div>
    <div class="flex flex-col items-center mt-8">
      <basic-btn
        to="/signup"
        class="w-full text-white bg-secondary-700 hover:bg-secondary-800 py-2.5 px-3.5 mb-4"
      >
        Sign Up
      </basic-btn>
      <basic-btn
        to="/login"
        class="w-full py-2.5 px-3 link border border-rose-800 dark:border-rose-500 hover:border-rose-900 dark:hover:border-rose-600"
      >
        Login
      </basic-btn>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import BasicBtn from "../button/BasicBtn.vue";
export default defineComponent({
  name: "sign-up-in-inner-overlay",
  components: { BasicBtn },
});
</script>