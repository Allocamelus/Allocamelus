<template>
  <center-form-box classes="overflow-x-hidden overflow-y-auto">
    <bar-loader :show="loading" />
    <div class="px-5 py-4">
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
          class="link mt-2 flex cursor-pointer items-center"
          @click="captcha.show = false"
        >
          <ChevronLeftIcon class="h-5 w-5"></ChevronLeftIcon> Back
        </div>
      </div>
      <div v-show="!showCaptcha">
        <h2 class="text-2xl font-medium">Login</h2>
        <html-errors
          v-if="err.login.length > 0"
          class="mt-3"
          :error="err.login"
        >
          <div v-if="err.loginType !== ''">
            <div v-if="err.loginType === 'Email'">
              Don't see the verification email?
              <to-link class="link" to="/account/verify-email"
                >Resend It</to-link
              >
            </div>
            <div v-else>
              Forgot your password?
              <to-link to="/account/reset_password" class="link"
                >Click Here</to-link
              >
              to reset it.
            </div>
          </div>
        </html-errors>
        <form ref="form" class="form mt-3" @submit.prevent="onSubmit">
          <div>
            <input-label for="name" :err="err.username">Username</input-label>
            <text-input
              v-model="username"
              name="name"
              :check="true"
              :required="true"
              @error="err.username = $event"
            >
            </text-input>
          </div>
          <div class="mt-3">
            <input-label for="password" :err="err.password"
              >Password</input-label
            >
            <password-input
              v-model="password"
              :required="true"
              @error="err.password = $event"
            ></password-input>
          </div>
          <div class="mt-3 flex justify-between">
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
            <submit class="mt-3 self-end" title="Login" :disabled="loading">
              Login
            </submit>
          </div>
        </form>
      </div>
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
import { ChevronLeftIcon } from "@heroicons/vue/20/solid";
import ToLink from "@/components/ToLink.vue";
import BarLoader from "@/components/overlay/BarLoader.vue";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import VueHcaptcha from "@hcaptcha/vue3-hcaptcha";

import { auth } from "@/api/account/auth";
import ApiResp from "@/models/responses";
import { salt as getSalt } from "@/api/account/salt";
import { getKeys } from "@/pkg/crypto/userKeys";
import { NullError } from "@/models/Error";
import HtmlErrors, {
  LoadingCaptcha,
  SomethingWentWrong,
} from "@/components/HtmlErrors.vue";

const HtmlInvalidUsernamePassword = [`Invalid Username/Email or Password`],
  HtmlUnverifiedEmail = [`Please verify your email to login`];

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
        login: "" as string | string[],
        loginType: "" as "" | "Email" | "UsernamePassword",
        username: "",
        password: "",
      },
      username: "",
      password: "",
      keys: {
        authKey: "",
        pgpPassphrase: "",
        err: null as NullError<any>,
      },
      remember: false,
      captcha: {
        show: false,
        loaded: false,
        siteKey: "",
        token: "",
        theme: state.theme,
      },
      loading: false,
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
      this.loading = true;

      if (this.err.username.length != 0 || this.err.password.length != 0) {
        this.loading = false;
        return;
      }

      if (this.keys.authKey == "" || this.keys.pgpPassphrase == "") {
        let salt = await getSalt(this.username).catch(() => {
          return this.handleErr();
        });
        if (salt == undefined) {
          this.loading = false;
          return;
        }
        if (salt.error != undefined) {
          return this.handleErr(salt.error);
        }

        this.keys = await getKeys(this.password, salt.salt);
        if (this.keys.err !== null) {
          return this.handleErr(String(this.keys.err));
        }
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
          return this.handleErr();
        });
    },
    handleErr(err?: string, captcha?: string) {
      this.loading = false;
      this.captcha.show = false;
      this.err.loginType = "";

      switch (err) {
        case ApiResp.Account.Auth.InvalidUsernamePassword:
          this.resetSensitive();
          this.err.login = HtmlInvalidUsernamePassword;
          this.err.loginType = "UsernamePassword";
          return;
        case ApiResp.Account.Auth.UnverifiedEmail:
          this.resetSensitive();
          this.err.login = HtmlUnverifiedEmail;
          this.err.loginType = "Email";
          return;
        case ApiResp.Shared.InvalidCaptcha:
          if (captcha == undefined) {
            this.err.login = SomethingWentWrong;
            throw new Error("login: Error missing captcha siteKey");
          }
          this.captcha.show = true;
          this.captcha.siteKey = captcha;
          this.err.login = LoadingCaptcha;
          return;
        default:
          this.resetSensitive();
          this.err.login = SomethingWentWrong;
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
    BarLoader,
    HtmlErrors,
  },
});
</script>
