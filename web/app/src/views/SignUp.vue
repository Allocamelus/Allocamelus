<template>
  <center-form-box>
    <div v-show="showCaptcha">
      <div id="signUpCaptchaContainer"></div>
      <div
        class="mt-2 link flex items-center cursor-pointer"
        @click="captcha.show = false"
      >
        <chevron-left-sm></chevron-left-sm> Back
      </div>
    </div>
    <div v-show="!showCaptcha">
      <h2 class="text-2xl font-medium">Sign Up</h2>
      <div v-if="err.signUp.length > 0" class="mt-3" v-html="err.signUp"></div>
      <form @submit.prevent="onSubmit" ref="form" class="form mt-3">
        <div>
          <input-label for="name" :err="err.username">Username</input-label>
          <text-input
            v-model="username"
            name="name"
            :check="true"
            :required="true"
            placeholder="mary-smith"
            :regex="/^[a-zA-Z0-9_-]*$/"
            regexMsg="Alphanumeric characters, Underscores and Dashes Only"
            @error="err.username = $event"
          ></text-input>
        </div>
        <div class="mt-3">
          <input-label for="email" :err="err.email">Email</input-label>
          <text-input
            v-model="email"
            :check="true"
            :required="true"
            :regex="/.+@.+\..+/"
            regexMsg="Invalid Email"
            placeholder="mary@example.com"
            @error="err.email = $event"
          ></text-input>
        </div>
        <div class="mt-3">
          <input-label for="password" :err="err.password">Password</input-label>
          <password-input
            v-model="password"
            :check="true"
            :required="true"
            @error="err.password = $event"
          ></password-input>
        </div>

        <div class="flex justify-between mt-3">
          <div class="flex flex-col justify-end">
            <small-text class="mr-3">
              By Signing Up, you agree to the
              <router-link class="link whitespace-nowrap" to="/tos"
                >Terms of Service</router-link
              >
            </small-text>
            <small-text class="mt-1 mr-3">
              Have an account?
              <router-link class="link whitespace-nowrap" to="/login"
                >Login</router-link
              >
            </small-text>
          </div>
          <submit class="mt-3 self-end whitespace-nowrap" title="Sign Up"
            >Sign Up</submit
          >
        </div>
      </form>
    </div>
  </center-form-box>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";

import CenterFormBox from "../components/form/CenterFormBox.vue";
import TextInput from "../components/form/TextInput.vue";
import PasswordInput from "../components/form/PasswordInput.vue";
import Checkbox from "../components/form/Checkbox.vue";
import Submit from "../components/form/Submit.vue";
import InputLabel from "../components/form/InputLabel.vue";
import SmallText from "../components/form/SmallText.vue";
import ChevronLeftSm from "../components/icon/ChevronLeftSm.vue";

import { API_AuthA10Token } from "../models/api_account_gen";
import { authA10 } from "../api/account/auth";
import ApiResp from "../models/responses";
import { siteKeys } from "../api/meta/captcha/siteKeys";
import {
  htmlErrBuilder,
  HtmlSomthingWentWrong,
} from "../components/htmlErrors";

const HtmlInvalidUsernamePassword = htmlErrBuilder(
    `Invalid Username/Email or Password`,
    `Forgot your password? <a href="/account/reset_password" class="link">Click Here</a> to reset it.`
  ),
  HtmlUnverifiedEmail = htmlErrBuilder(
    `Please verify your email to login`,
    `Don't see the verification email? <a class="link" href="/account/verify_email">Resend It</a>`
  ),
  HtmlLoadingCaptcha = htmlErrBuilder("Loading captcha...");

function gotoRedirect(router, redirect) {
  var url = "/";
  if (redirect?.length > 0) {
    url = props.redirect;
  }
  router.push(url);
}

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
        signUp: "",
        username: "",
        email: "",
        password: "",
      },
      username: "",
      email: "",
      password: "",
      remember: false,
      captcha: {
        show: false,
        loaded: false,
        siteKey: "",
        token: "",
      },
    });

    if (store.getters.loggedIn) {
      gotoRedirect(router, props.redirect);
    }

    document.title = `Sign Up - ${import.meta.env.VITE_SITE_NAME}`;

    return {
      ...toRefs(data),
    };
  },
  created() {
    (async () => {
      siteKeys().then((sk) => {
        this.captcha.siteKey = sk.siteKey(sk.difficulties.user.create);
        var srcCap = document.createElement("script");
        srcCap.src =
          "https://hcaptcha.com/1/api.js?onload=signUpCaptachaLoaded&render=explicit";
        document.head.appendChild(srcCap);
      });
      // TODO: Error handling
    })();
    window.signUpCaptachaSubmit = () => {
      this.onSubmit();
    };
    window.signUpCaptachaLoaded = () => {
      this.captcha.loaded = true;
      window.hcaptcha.render("signUpCaptchaContainer", {
        sitekey: this.captcha.siteKey,
        theme: "dark",
        callback: "signUpCaptachaSubmit",
      });
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
        API_AuthA10Token.createFrom({
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
                vm.err.signUp = HtmlInvalidUsernamePassword;
                return;
              case ApiResp.Account.Auth.UnverifiedEmail:
                vm.err.signUp = HtmlUnverifiedEmail;
                return;
              case ApiResp.Shared.InvalidCaptcha:
                vm.captcha.show = true;
                vm.captcha.siteKey = r.captcha;
                vm.err.signUp = HtmlLoadingCaptcha;
                if (!vm.captcha.loaded) {
                  var srcCap = document.createElement("script");
                  srcCap.src =
                    "https://hcaptcha.com/1/api.js?onload=signUpCaptachaLoaded&render=explicit";
                  document.head.appendChild(srcCap);
                }
                return;
              default:
                vm.err.signUp = HtmlSomthingWentWrong;
                return;
            }
          } else {
            vm.$store.dispatch("newLoginSession", {
              userId: r.userId,
              authToken: vm.remember,
            });
            gotoRedirect(vm.$router, vm.redirect);
          }
        })
        .catch((e) => {
          vm.captcha.show = false;
          vm.err.signUp = HtmlSomthingWentWrong;
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
    SmallText,
    ChevronLeftSm,
  },
});
</script>
