<template>
  <center-form-box classes="overflow-x-hidden overflow-y-auto">
    <bar-loader :show="loading" />
    <div class="py-4 px-5">
      <div v-show="showCaptcha">
        <div v-if="captcha.siteKey.length > 0">
          <vue-hcaptcha
            class="mx-auto max-w-max"
            :sitekey="captcha.siteKey"
            :theme="captcha.theme"
            @rendered="captcha.loaded = true"
            @verify="
              (token) => {
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
                :regexMsg="errMsgUserName"
                @error="err.userName = $event"
              ></text-input>
            </div>
            <div class="mt-3">
              <input-label for="email" :err="err.email">Email</input-label>
              <email-input
                v-model="email"
                :check="true"
                :required="true"
                @error="err.email = $event"
              ></email-input>
            </div>
            <div class="mt-3">
              <input-label for="password" :err="err.password">
                Password
              </input-label>
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
                  <to-link class="link whitespace-nowrap" to="/tos">
                    Terms of Service
                  </to-link>
                </text-small>
                <text-small class="mt-1 mr-3">
                  Have an account?
                  <to-link class="link whitespace-nowrap" to="/login">
                    Login
                  </to-link>
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
            Verify your account by clicking the activation link sent to your
            email
          </div>
          <div>Backup Key:</div>
          <text-small>
            Save this, you'll need this to recover your account if you lose your
            password
          </text-small>
          <input-copy v-model="recoveryKey" class="my-2"></input-copy>
        </div>
      </div>
    </div>
  </center-form-box>
</template>

<script lang="ts">
import { defineComponent, toRefs, reactive } from "vue";
import { useStore } from "vuex";
import { redirectUrl } from "../router";

import CenterFormBox from "../components/form/CenterFormBox.vue";
import TextInput from "../components/form/TextInput.vue";
import PasswordInput from "../components/form/PasswordInput.vue";
import Submit from "../components/form/Submit.vue";
import InputLabel from "../components/form/InputLabel.vue";
import TextSmall from "../components/text/Small.vue";
import ChevronLeftIcon from "@heroicons/vue/solid/ChevronLeftIcon";
import InputCopy from "../components/form/InputCopy.vue";
import EmailInput from "../components/form/EmailInput.vue";
import BarLoader from "../components/overlay/BarLoader.vue";

// @ts-ignore
import VueHcaptcha from "@hcaptcha/vue3-hcaptcha";

import { genKeys } from "../pkg/crypto/userKeys";

import { create as createAccount } from "../api/account/create";
import ApiResp, { RespToError } from "../models/responses";
import { siteKeys } from "../api/meta/captcha/siteKeys";
import {
  HtmlSomethingWentWrong,
  HtmlLoadingCaptcha,
} from "../components/htmlErrors";
import ToLink from "../components/ToLink.vue";

export default defineComponent({
  props: {
    redirect: {
      type: String,
      default: "",
    },
  },
  setup() {
    const store = useStore();
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
      captcha: {
        show: false,
        loaded: false,
        siteKey: "",
        token: "",
        theme: store.getters.theme,
      },
      recoveryKey: "",
      showForm: true,
      loading: false,
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
      errMsgUserName: "Alphanumeric characters, Underscores and Dashes Only",
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

      if (
        this.err.userName.length != 0 ||
        this.err.email.length != 0 ||
        this.err.password.length != 0
      ) {
        this.loading = false;
        return;
      }

      if (this.captcha.token.length <= 0) {
        this.captcha.show = true;
        this.loading = false;
        return;
      }
      this.captcha.show = false;

      let keys = await genKeys(this.userName, this.password);
      let userKey = keys.userKey;
      this.recoveryKey = keys.recoveryKey;

      createAccount({
        userName: this.userName,
        email: this.email,
        captcha: this.captcha.token,
        auth: {
          salt: userKey.keySaltEncoded,
          hash: userKey.keyAuthHash,
        },
        key: {
          publicArmored: userKey.publicArmored,
          privateArmored: userKey.privateArmored,
          recoveryHash: userKey.recoveryHash,
          recoveryArmored: userKey.recoveryArmored,
        },
      })
        .then((r) => {
          this.loading = false;
          if (!r.success) {
            if (typeof r.errors === "object") {
              if (Array.isArray(r.errors)) {
                r.errors.forEach((err) => {
                  switch (err) {
                    case ApiResp.User.Create.LoggedIn:
                      this.$router.push(redirectUrl(this.redirect));
                      return;
                    case ApiResp.Shared.InvalidCaptcha:
                      this.loading = true;
                      this.captcha.siteKey = "";
                      this.err.signUp = HtmlLoadingCaptcha;
                      siteKeys().then((sk) => {
                        this.captcha.siteKey = sk.siteKey(
                          sk.difficulties.user.create
                        );
                      });
                      this.captcha.show = true;
                      this.loading = false;
                      this.err.signUp = "";
                      return;
                    default:
                      this.err.signUp = HtmlSomethingWentWrong;
                      return;
                  }
                });
              } else {
                for (const key in r.errors) {
                  if (Object.hasOwnProperty.call(r.errors, key)) {
                    var err = r.errors[key];
                    var errText = RespToError(err);
                    if (errText.length > 0) {
                      switch (key) {
                        case "userName":
                          this.err.userName = errText;
                          break;
                        case "email":
                          this.err.email = errText;
                          break;
                        case "password":
                          this.err.password = errText;
                          break;
                        default:
                          this.err.signUp = HtmlSomethingWentWrong;
                          break;
                      }
                    } else {
                      this.err.signUp = HtmlSomethingWentWrong;
                    }
                  }
                }
              }
            } else {
              this.err.signUp = HtmlSomethingWentWrong;
            }
          } else {
            this.captcha.show = false;
            this.showForm = false;
          }
        })
        .catch(() => {
          this.loading = false;
          this.captcha.show = false;
          this.err.signUp = HtmlSomethingWentWrong;
        });
    },
  },
  components: {
    CenterFormBox,
    TextInput,
    PasswordInput,
    Submit,
    InputLabel,
    TextSmall,
    ChevronLeftIcon,
    VueHcaptcha,
    InputCopy,
    ToLink,
    EmailInput,
    BarLoader,
  },
});
</script>
