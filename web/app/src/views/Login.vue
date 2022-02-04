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
            (token: string) => {
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
              <to-link class="link whitespace-nowrap" to="/signup">
                Sign Up
              </to-link>
            </text-small>
          </div>
          <submit class="mt-3 self-end" title="Login">Login</submit>
        </div>
      </form>
    </div>
  </center-form-box>
</template>

<script lang="ts">
import { defineComponent, toRefs, reactive } from "vue";
import { redirectUrl } from "@/router";
import { useStateStore } from "@/store";
import { Session, useSessionStore } from "@/store/session";

import CenterFormBox from "@/components/form/CenterFormBox.vue";
import TextInput from "@/components/form/TextInput.vue";
import PasswordInput from "@/components/form/PasswordInput.vue";
import Checkbox from "@/components/form/Checkbox.vue";
import Submit from "@/components/form/Submit.vue";
import InputLabel from "@/components/form/InputLabel.vue";
import TextSmall from "@/components/text/Small.vue";
import ChevronLeftIcon from "@heroicons/vue/solid/ChevronLeftIcon";
import ToLink from "@/components/ToLink.vue";

// @ts-ignore
import VueHcaptcha from "@hcaptcha/vue3-hcaptcha";

import { auth } from "@/api/account/auth";
import ApiResp from "@/models/responses";
import {
  htmlErrBuilder,
  HtmlSomethingWentWrong,
  HtmlLoadingCaptcha,
} from "@/components/htmlErrors";
import { salt as getSalt } from "@/api/account/salt";
import { hashSalt, parse } from "@/pkg/crypto/argon2id";
import { getKeys } from "@/pkg/crypto/userKeys";

const HtmlInvalidUsernamePassword = htmlErrBuilder(
    `Invalid Username/Email or Password`,
    `Forgot your password? <a href="/account/reset_password" class="link">Click Here</a> to reset it.`
  ),
  HtmlUnverifiedEmail = htmlErrBuilder(
    `Please verify your email to login`,
    `Don't see the verification email? <a class="link" href="/account/verify-email">Resend It</a>`
  );

export default defineComponent({
  props: {
    redirect: {
      type: String,
      default: "",
    },
  },
  setup() {
    const state = useStateStore();
    const session = useSessionStore();
    const data = reactive({
      err: {
        login: "",
        username: "",
        password: "",
      },
      username: "",
      password: "",
      keys: {
        authKey: "",
        pgpPassphrase: "",
      },
      remember: false,
      captcha: {
        show: false,
        loaded: false,
        siteKey: "",
        token: "",
        theme: state.theme,
      },
    });

    document.title = `Login - ${import.meta.env.VITE_SITE_NAME}`;

    return {
      ...toRefs(data),
      session,
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
    async onSubmit() {
      if (this.err.username.length != 0 || this.err.password.length != 0) {
        return;
      }

      if (this.keys.authKey == "" || this.keys.pgpPassphrase == "") {
        let salt = await getSalt(this.username).catch(() => {
          this.handleErr();
          return undefined;
        });
        if (salt == undefined) return;
        if (salt.error != undefined) {
          return this.handleErr(salt.error);
        }

        this.keys = await getKeys(this.password, salt.salt);
      }

      auth({
        userName: this.username,
        authKey: this.keys.authKey,
        remember: this.remember,
        captcha: this.captcha.token,
      })
        .then((r) => {
          if (!r.success) {
            return this.handleErr(r.error, r.captcha);
          } else {
            this.session.$patch(
              new Session({
                loggedIn: true,
                user: r.user,
              })
            );
            this.$router.push(redirectUrl(this.redirect));
          }
        })
        .catch(() => {
          this.handleErr();
        });
    },
    handleErr(err?: string, captcha?: string) {
      this.captcha.show = false;

      switch (err) {
        case ApiResp.Account.Auth.InvalidUsernamePassword:
          this.resetSensitive();
          this.err.login = HtmlInvalidUsernamePassword;
          return;
        case ApiResp.Account.Auth.UnverifiedEmail:
          this.resetSensitive();
          this.err.login = HtmlUnverifiedEmail;
          return;
        case ApiResp.Shared.InvalidCaptcha:
          if (captcha == undefined) {
            this.err.login = HtmlSomethingWentWrong;
            throw new Error("login: Error missing captcha siteKey");
          }
          this.captcha.show = true;
          this.captcha.siteKey = captcha;
          this.err.login = HtmlLoadingCaptcha;
          return;
        default:
          this.resetSensitive();
          this.err.login = HtmlSomethingWentWrong;
          throw new Error(err);
      }
    },
    resetSensitive() {
      this.password = this.keys.authKey = this.keys.pgpPassphrase = "";
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
    ToLink,
  },
});
</script>
