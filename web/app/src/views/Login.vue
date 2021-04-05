<template>
  <center-form-box>
    <div v-show="showCaptcha">
      <div v-if="captcha.siteKey.length > 0">
        <vue-hcaptcha
          class="mx-auto max-w-max"
          :sitekey="captcha.siteKey"
          :theme="captcha.theme"
          @rendered="captcha.loaded = true"
          @verify="
            (token, eKey) => {
              captcha.token = token;
              onSubmit();
            }
          "
          @expired="captcha.token = ''"
        ></vue-hcaptcha>
      </div>
      <div
        class="mt-2 link flex items-center cursor-pointer"
        @click="captcha.show = false"
      >
        <ChevronLeftIcon class="w-5 h-5"></ChevronLeftIcon> Back
      </div>
    </div>
    <div v-show="!showCaptcha">
      <h2 class="text-2xl font-medium">Login</h2>
      <div v-if="err.login.length > 0" class="mt-3" v-html="err.login"></div>
      <form @submit.prevent="onSubmit" ref="form" class="form mt-3">
        <div>
          <input-label for="name" :err="err.username">Username</input-label>
          <text-input
            v-model="username"
            name="name"
            :check="true"
            :required="true"
            @error="err.username = $event"
          ></text-input>
        </div>
        <div class="mt-3">
          <input-label for="password" :err="err.password">Password</input-label>
          <password-input
            v-model="password"
            :required="true"
            @error="err.password = $event"
          ></password-input>
        </div>
        <div class="flex justify-between mt-3">
          <div class="flex flex-col">
            <checkbox v-model="remember" name="remember">
              <label for="remember">Remember Me</label>
            </checkbox>
            <text-small class="mt-2 mr-3">
              Don't have an account?
              <router-link class="link whitespace-nowrap" to="/signup">
                Sign Up
              </router-link>
            </text-small>
          </div>
          <submit class="mt-3 self-end" title="Login">Login</submit>
        </div>
      </form>
    </div>
  </center-form-box>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { useRouter } from "vue-router";
import { redirectUrl } from "../router";
import { useStore } from "vuex";

import CenterFormBox from "../components/form/CenterFormBox.vue";
import TextInput from "../components/form/TextInput.vue";
import PasswordInput from "../components/form/PasswordInput.vue";
import Checkbox from "../components/form/Checkbox.vue";
import Submit from "../components/form/Submit.vue";
import InputLabel from "../components/form/InputLabel.vue";
import TextSmall from "../components/text/Small.vue";
import ChevronLeftIcon from "@heroicons/vue/solid/ChevronLeftIcon";

import VueHcaptcha from "@jdinabox/vue-3-hcaptcha";

import { GEN_AuthA10Token } from "../models/go_structs_gen";
import { authA10 } from "../api/account/auth";
import ApiResp from "../models/responses";
import {
  htmlErrBuilder,
  HtmlSomthingWentWrong,
  HtmlLoadingCaptcha,
} from "../components/htmlErrors";

const HtmlInvalidUsernamePassword = htmlErrBuilder(
    `Invalid Username/Email or Password`,
    `Forgot your password? <a href="/account/reset_password" class="link">Click Here</a> to reset it.`
  ),
  HtmlUnverifiedEmail = htmlErrBuilder(
    `Please verify your email to login`,
    `Don't see the verification email? <a class="link" href="/account/verify_email">Resend It</a>`
  );

export default defineComponent({
  props: {
    redirect: {
      type: String,
    },
  },
  setup(props) {
    const store = useStore();
    const router = useRouter();
    const data = reactive({
      err: {
        login: "",
        username: "",
        password: "",
      },
      username: "",
      password: "",
      remember: false,
      captcha: {
        show: false,
        loaded: false,
        siteKey: "",
        token: "",
        theme: store.getters.theme,
      },
    });

    document.title = `Login - ${import.meta.env.VITE_SITE_NAME}`;

    return {
      ...toRefs(data),
    };
  },
  computed: {
    showCaptcha() {
      if (this.captcha.show) {
        return this.captcha.loaded;
      }
      return false;
    },
  },
  methods: {
    onSubmit(e) {
      var vm = this;
      if (vm.err.username.length != 0 || vm.err.password.length != 0) {
        return;
      }
      try {
        vm.captcha.token = document.getElementsByName(
          "h-captcha-response"
        )[0].value;
      } catch (e) {
        vm.captcha.token = "";
      }
      authA10(
        GEN_AuthA10Token.createFrom({
          userName: vm.username,
          password: vm.password,
          remember: vm.remember,
          captcha: vm.captcha.token,
        })
      )
        .then((r) => {
          vm.captcha.show = false;
          if (!r.success) {
            switch (r.error) {
              case ApiResp.Account.Auth.InvalidUsernamePassword:
                vm.err.login = HtmlInvalidUsernamePassword;
                return;
              case ApiResp.Account.Auth.UnverifiedEmail:
                vm.err.login = HtmlUnverifiedEmail;
                return;
              case ApiResp.Shared.InvalidCaptcha:
                vm.captcha.show = true;
                vm.captcha.siteKey = r.captcha;
                vm.err.login = HtmlLoadingCaptcha;
                return;
              default:
                vm.err.login = HtmlSomthingWentWrong;
                return;
            }
          } else {
            vm.$store.dispatch("newLoginSession", {
              user: r.user,
              authToken: vm.remember,
            });
            this.$router.push(redirectUrl(vm.redirect));
          }
        })
        .catch((e) => {
          vm.captcha.show = false;
          vm.err.login = HtmlSomthingWentWrong;
        });
    },
  },
  components: {
    CenterFormBox,
    TextInput,
    PasswordInput,
    Checkbox,
    Submit,
    InputLabel,
    TextSmall,
    ChevronLeftIcon,
    VueHcaptcha,
  },
});
</script>
