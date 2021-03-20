<template>
  <center-form-box>
    <div v-show="showCaptcha">
      <div id="loginCaptchaContainer"></div>
      <div class="mt-3 link" @click="captcha.show = false">
        <i class="fas fa-chevron-left"></i> Back
      </div>
    </div>
    <div v-show="!showCaptcha">
      <h2>Login</h2>
      <div v-if="err.login.length > 0" class="mt-3" v-html="err.login"></div>
      <form @submit.prevent="onSubmit" ref="form" class="form mt-3">
        <div>
          <input-label :for="'name'" :err="err.username">Username</input-label>
          <text-input
            v-model="username"
            name="name"
            :check="true"
            :required="true"
            @error="err.username = $event"
          ></text-input>
        </div>
        <div class="mt-2">
          <input-label :for="'password'" :err="err.password"
            >Password</input-label
          >
          <password-input
            v-model="password"
            :required="true"
            @error="err.password = $event"
          ></password-input>
        </div>
        <div class="bottom-wrapper">
          <div class="bottom-text">
            <checkbox v-model="remember" :name="'remember'">
              <label for="remember">Remember Me</label>
            </checkbox>
            <div class="box__text--medium mt-3">
              Don't have an account?
              <a class="link" href="/signup">Sign Up</a>
            </div>
          </div>
          <submit class="mt-3" :title="'Login'">Login</submit>
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
import { API_AuthA10Token } from "../models/api_account_gen";
import { authA10 } from "../api/account/auth";
import InputLabel from "../components/form/InputLabel.vue";
import ApiResp from "../models/responses";

function htmlErrBuilder(first, second = "") {
  var err = `<strong>`;
  err += first;
  err += `</strong>`;
  if (second.length > 0) {
    err += `<br>` + second;
  }
  return err;
}

const HtmlInvalidUsernamePassword = htmlErrBuilder(
    `Invalid Username/Email or Password`,
    `Forgot your password? <a href="/account/reset_password" class="link">Click Here</a> to reset it.`
  ),
  HtmlUnverifiedEmail = htmlErrBuilder(
    `Please verify your email to login`,
    `Don't see the verification email? <a class="link" href="/account/verify_email">Resend It</a>`
  ),
  HtmlSomthingWentWrong = htmlErrBuilder(
    `Something went wrong`,
    `Try again later`
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
      },
    });

    if (store.getters.loggedIn) {
      gotoRedirect(router, props.redirect)
    }

    return {
      ...toRefs(data),
    };
  },
  created() {
    window.loginCaptachaSubmit = () => {
      this.onSubmit();
    };
    window.loginCaptachaLoaded = () => {
      this.captcha.loaded = true;
      window.hcaptcha.render("loginCaptchaContainer", {
        sitekey: this.captcha.siteKey,
        theme: "dark",
        callback: "loginCaptachaSubmit",
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
          uniqueName: vm.username,
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
                if (!vm.captcha.loaded) {
                  var srcCap = document.createElement("script");
                  srcCap.src =
                    "https://hcaptcha.com/1/api.js?onload=loginCaptachaLoaded&render=explicit";
                  document.head.appendChild(srcCap);
                }
                return;
              default:
                vm.err.login = HtmlSomthingWentWrong;
                return;
            }
          } else {
            vm.$store.dispatch("newLoginSession", {
              userId: r.userId,
              authToken: vm.remember,
            });
            gotoRedirect(vm.$router, vm.redirect)
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
  },
});
</script>

<style lang="scss" scoped>
.bottom-wrapper {
  display: flex;
  justify-content: space-between;
}
.submit {
  align-self: flex-end;
}
</style>