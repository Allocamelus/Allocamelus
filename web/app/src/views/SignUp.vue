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
      <div v-show="showForm">
        <h2 class="text-2xl font-medium">Sign Up</h2>
        <div
          v-if="err.signUp.length > 0"
          class="mt-3"
          v-html="err.signUp"
        ></div>
        <form @submit.prevent="onSubmit" ref="form" class="form mt-3">
          <div>
            <input-label for="name" :err="err.userName">Username</input-label>
            <text-input
              v-model="userName"
              name="name"
              :check="true"
              :required="true"
              :minLen="5"
              :maxLen="64"
              placeholder="mary-smith"
              :regex="/^[a-zA-Z0-9_-]*$/"
              :regexMsg="errMsg.userName"
              @error="err.userName = $event"
            ></text-input>
          </div>
          <div class="mt-3">
            <input-label for="email" :err="err.email">Email</input-label>
            <text-input
              v-model="email"
              :check="true"
              :required="true"
              type="email"
              :regex="/.+@.+\..+/"
              :regexMsg="errMsg.email"
              placeholder="mary@example.com"
              @error="err.email = $event"
            ></text-input>
          </div>
          <div class="mt-3">
            <input-label for="password" :err="err.password"
              >Password</input-label
            >
            <password-input
              v-model="password"
              :check="true"
              :required="true"
              @error="err.password = $event"
            ></password-input>
          </div>

          <div class="flex justify-between mt-3">
            <div class="flex flex-col justify-end">
              <text-small class="mr-3">
                By Signing Up, you agree to the
                <router-link class="link whitespace-nowrap" to="/tos">
                  Terms of Service
                </router-link>
              </text-small>
              <text-small class="mt-1 mr-3">
                Have an account?
                <router-link class="link whitespace-nowrap" to="/login">
                  Login
                </router-link>
              </text-small>
            </div>
            <submit
              class="mt-3 self-end whitespace-nowrap"
              :title="!captcha.loaded ? 'Loading Captcha' : 'Sign Up'"
              :disabled="!captcha.loaded"
            >
              Sign Up
            </submit>
          </div>
        </form>
      </div>
      <div v-show="!showForm" class="font-medium">
        <div class="text-lg flex">
          Account successfully created
          <div class="text-base">*</div>
        </div>
        <text-small>
          *If an account with the email ({{ email }}) doesn't already exist
        </text-small>
        <div class="font-normal my-2">
          Verify your account by clicking the activation link sent to your email
        </div>
        <div>Backup Key:</div>
        <text-small>
          Save this, you'll need this to recover your account
        </text-small>
        <input-copy
          v-model="backupKey"
          :watchModel="true"
          class="my-2"
        ></input-copy>
      </div>
    </div>
  </center-form-box>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import { redirectUrl } from "../router";

import CenterFormBox from "../components/form/CenterFormBox.vue";
import TextInput from "../components/form/TextInput.vue";
import PasswordInput from "../components/form/PasswordInput.vue";
import Checkbox from "../components/form/Checkbox.vue";
import Submit from "../components/form/Submit.vue";
import InputLabel from "../components/form/InputLabel.vue";
import TextSmall from "../components/text/Small.vue";
import ChevronLeftIcon from "@heroicons/vue/solid/ChevronLeftIcon";
import InputCopy from "../components/form/InputCopy.vue";

import VueHcaptcha from "@jdinabox/vue-3-hcaptcha";

import { API_CreateA10Token } from "../models/api_user_gen";
import { createA9s } from "../api/user/create";
import ApiResp from "../models/responses";
import { siteKeys } from "../api/meta/captcha/siteKeys";
import {
  HtmlSomthingWentWrong,
  HtmlLoadingCaptcha,
} from "../components/htmlErrors";

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
        userName: "",
        email: "",
        password: "",
      },
      userName: "",
      email: "",
      password: "",
      remember: false,
      captcha: {
        show: false,
        loaded: false,
        siteKey: "",
        token: "",
        theme: store.getters.theme,
      },
      backupKey: "",
      showForm: true,
    });

    document.title = `Sign Up - ${import.meta.env.VITE_SITE_NAME}`;

    (async () => {
      siteKeys().then((sk) => {
        data.captcha.siteKey = sk.siteKey(sk.difficulties.user.create);
      });
      // TODO: Error handling
    })();

    return {
      ...toRefs(data),
      errMsg: {
        userName: "Alphanumeric characters, Underscores and Dashes Only",
        email: "Invalid Email",
      },
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
      if (
        vm.err.userName.length != 0 ||
        vm.err.email.length != 0 ||
        vm.err.password.length != 0
      ) {
        return;
      }

      if (vm.captcha.token.length <= 0) {
        vm.captcha.show = true;
        return;
      }

      createA9s(
        API_CreateA10Token.createFrom({
          userName: vm.userName,
          email: vm.email,
          password: vm.password,
          captcha: vm.captcha.token,
        })
      )
        .then((r) => {
          vm.captcha.show = false;
          if (!r.success) {
            if (typeof r.errors === "object") {
              if (Array.isArray(r.errors)) {
                r.errors.forEach((err) => {
                  switch (err) {
                    case ApiResp.User.Create.LoggedIn:
                      this.$router.push(redirectUrl(vm.redirect));
                      return;
                    case ApiResp.Shared.InvalidCaptcha:
                      vm.captcha.siteKey = "";
                      vm.err.signUp = HtmlLoadingCaptcha;
                      siteKeys().then((sk) => {
                        vm.captcha.siteKey = sk.siteKey(
                          sk.difficulties.user.create
                        );
                      });
                      vm.captcha.show = true;
                      vm.err.signUp = "";
                      return;
                    default:
                      vm.err.signUp = HtmlSomthingWentWrong;
                      return;
                  }
                });
              } else {
                for (const key in r.errors) {
                  if (Object.hasOwnProperty.call(r.errors, key)) {
                    var err = r.errors[key];
                    switch (key) {
                      case "userName":
                        switch (err) {
                          case ApiResp.User.Validate.UserName.Taken:
                            vm.err.userName = "Username Taken";
                            break;
                          case ApiResp.User.Validate.UserName.Length:
                            vm.err.userName = "Invalid Length";
                            break;
                          case ApiResp.User.Validate.Invalid: {
                            vm.err.userName = vm.errMsg.userName;
                            break;
                          }

                          default:
                            vm.err.signUp = HtmlSomthingWentWrong;
                            break;
                        }
                        break;
                      case "email":
                        switch (err) {
                          case ApiResp.User.Validate.Email.Invalid:
                            vm.err.email = "Invalid Email";
                            break;
                          case ApiResp.User.Validate.Invalid: {
                            vm.err.email = vm.errMsg.email;
                            break;
                          }

                          default:
                            vm.err.signUp = HtmlSomthingWentWrong;
                            break;
                        }
                        break;
                      case "password":
                        switch (err) {
                          case ApiResp.User.Validate.Password.Length:
                            vm.err.password = "Invalid Length";
                            break;
                          case ApiResp.User.Validate.Password.Strength:
                            vm.err.password = "Weak Password";
                            break;

                          default:
                            vm.err.signUp = HtmlSomthingWentWrong;
                            break;
                        }
                        break;

                      default:
                        vm.err.signUp = HtmlSomthingWentWrong;
                        break;
                    }
                  }
                }
              }
            } else {
              vm.err.signUp = HtmlSomthingWentWrong;
            }
          } else {
            vm.backupKey = r.backupKey;
            vm.captcha.show = false;
            vm.showForm = false;
          }
        })
        .catch((e) => {
          vm.captcha.show = false;
          vm.err.signUp = HtmlSomthingWentWrong;
        });
    },
    bk() {
      this.backupKey = "babcbc";
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
    InputCopy,
  },
});
</script>
