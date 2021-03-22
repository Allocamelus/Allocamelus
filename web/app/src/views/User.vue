<template>
  <div class="container">
    <error-box :error="err" class="py-3 px-4">
      <user-name :user="user" :displayType="TwoLine"></user-name>
      <div></div>
      <div>{{ user.bio }}</div>
    </error-box>
    <div class="flex">
      <div class="container__feed"></div>
      <div class="container__sidebar"></div>
    </div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { useStore } from "vuex";
import { get as getUser } from "../api/user/get";
import { User } from "../models/user_gen";
import UserName, { TwoLine } from "../components/user/Name.vue";
import ErrorBox from "../components/ErrorBox.vue";
import ApiResp from "../models/responses";
import { Html404Func, HtmlSomthingWentWrong } from "../components/htmlErrors";
import { API_Error } from "../models/api_error";
import { useRouter } from "vue-router";

function userErrors(api_error, path) {
  if (api_error instanceof API_Error) {
    switch (api_error.error) {
      case ApiResp.Shared.NotFound:
        return Html404Func(path);
    }
  }
  return HtmlSomthingWentWrong;
}
export default defineComponent({
  props: {
    uniqueName: {
      type: Array,
      required: true,
    },
  },
  setup(props) {
    const route = useRouter();
    const store = useStore();
    const data = reactive({
      user: new User(),
      err: "",
    });

    getUser(props.uniqueName[0])
      .then((r) => {
        data.user = r;
      })
      .catch((e) => {
        data.err = userErrors(e, route.currentRoute.value.fullPath);
      });

    return {
      ...toRefs(data),
      TwoLine,
    };
  },
  watch: {
    user(newUser, old) {
      document.title = `${newUser.name} (@${newUser.uniqueName}) - ${
        import.meta.env.VITE_SITE_NAME
      }`;
    },
  },
  async beforeRouteUpdate(to, from) {
    this.user = new User();

    getUser(to.params.uniqueName[0])
      .then((r) => {
        this.user = r;
      })
      .catch((e) => {
        this.err = userErrors(e, route.currentRoute.value.fullPath);
      });
  },
  components: {
    UserName,
    ErrorBox,
  },
});
</script>

<style>
</style>