<template>
  <center-form-box>
    <div v-if="hasSelectorToken">
      <div v-if="loading" class="flex justify-center items-center">
        <spin-loader class="w-5 h-5 mr-1.5" />
        <div class="font-medium">Checking...</div>
      </div>
      <div v-else>
        <div v-if="success">
          <div class="text-lg font-semibold">Success</div>
          <div>
            You can now
            <to-link to="/login" class="link font-medium">login</to-link>
          </div>
        </div>
        <div v-else>
          <div class="text-lg font-semibold">Error verifying email</div>
          <div v-if="err == UserResp.EmailToken.Validate.Expired">
            <text-small>Expired Email Token</text-small>
            <text-small> Your verification token/link has expired </text-small>
          </div>
          <div v-else>
            <text-small>Invalid Email Token</text-small>
            <div class="font-medium mt-1.5">Why am I seeing this?</div>
            <ul class="list-disc list-inside">
              <text-small>
                <li>You may have already verified your email</li>
                <li>Something Went Wrong</li>
              </text-small>
            </ul>
            <text-small class="mt-1.5">
              Try refreshing this page or request verification
            </text-small>
          </div>
          <div class="mt-1.5">
            <to-link to="/account/verify-email" class="link font-medium">
              Resend Verification Email
            </to-link>
          </div>
        </div>
      </div>
    </div>
    <div v-else>
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
        <div v-if="showForm">
          <h2 class="text-2xl font-medium">Resend Verification Email</h2>
          <div
            v-if="err.create.length > 0"
            class="mt-3"
            v-html="err.create /* skipcq: JS-0693 */"
          ></div>
          <form @submit.prevent="onSubmit" ref="form" class="mt-3">
            <input-label for="email" :err="err.email">Email</input-label>
            <email-input
              v-model="email"
              :check="true"
              :required="true"
              @error="err.email = $event"
            ></email-input>

            <div class="flex justify-end mt-4">
              <submit
                class="whitespace-nowrap"
                :title="!captcha.loaded ? 'Loading Captcha' : 'Sign Up'"
                :disabled="!captcha.loaded"
              >
                Submit
              </submit>
            </div>
          </form>
        </div>
        <div v-else>
          <h2 class="text-2xl font-medium">Success</h2>
          <div>A verification email to {{ email }} was sent*</div>
          <text-small class="mt-1.5"
            >*If an unverified account with attached email exist</text-small
          >
        </div>
      </div>
    </div>
  </center-form-box>
</template>

<script>
import { defineComponent, toRefs, reactive, computed } from "vue";
import { useStore } from "vuex";

import ApiResp, { RespToError } from "../../models/responses";
import { validate } from "../../api/user/email-token/validate";
import { create } from "../../api/user/email-token/create";
import { siteKeys } from "../../api/meta/captcha/siteKeys";
import {
  HtmlLoadingCaptcha,
  HtmlSomethingWentWrong,
} from "../../components/htmlErrors";

import VueHcaptcha from "@hcaptcha/vue3-hcaptcha";

import SpinLoader from "../../components/icons/SpinLoader.vue";
import ToLink from "../../components/ToLink.vue";
import TextSmall from "../../components/text/Small.vue";
import EmailInput from "../../components/form/EmailInput.vue";
import CenterFormBox from "../../components/form/CenterFormBox.vue";
import Submit from "../../components/form/Submit.vue";
import InputLabel from "../../components/form/InputLabel.vue";
import ChevronLeftIcon from "@heroicons/vue/solid/ChevronLeftIcon";

function hasST(selector, token) {
  return selector != undefined && token != undefined
    ? selector.length != 0 && token.length != 0
      ? true
      : false
    : false;
}

export default defineComponent({
  props: {
    selector: {
      type: String,
      default: "",
    },
    token: {
      type: String,
      default: "",
    },
  },
  setup(props) {
    const store = useStore();
    const data = reactive({
      loading: true,
      success: false,
      err: "",
      email: "",
      err: {
        email: "",
        create: "",
      },
      captcha: {
        show: false,
        loaded: false,
        siteKey: "",
        token: "",
        theme: store.getters.theme,
      },
      showForm: true,
    });

    const hasSelectorToken = computed(() => {
      return hasST(props.selector, props.token);
    });

    if (hasSelectorToken.value) {
      validate(props.selector, props.token)
        .then((r) => {
          data.success = r.success;
          data.err = r.error;
        })
        .finally(() => {
          data.loading = false;
        });
    }

    (async () => {
      siteKeys().then((sk) => {
        data.captcha.siteKey = sk.siteKey(sk.difficulties.user.emailToken);
      });
      // TODO: Error handling
    })();

    return {
      ...toRefs(data),
      hasSelectorToken,
      UserResp: ApiResp.User,
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
    onSubmit() {
      if (this.err.email.length != 0) {
        return;
      }

      if (this.captcha.token.length <= 0) {
        this.captcha.show = true;
        return;
      }

      create(this.email, this.captcha.token)
        .then((r) => {
          this.captcha.show = false;
          if (r.success) {
            this.showForm = false;
          } else {
            switch (r.error) {
              case ApiResp.User.Validate.Email.Invalid:
                this.err.email = RespToError(r.error);
                return;
              case ApiResp.Shared.InvalidCaptcha:
                this.captcha.siteKey = "";
                this.err.create = HtmlLoadingCaptcha;
                siteKeys().then((sk) => {
                  this.captcha.siteKey = sk.siteKey(
                    sk.difficulties.user.emailToken
                  );
                });
                this.captcha.show = true;
                this.err.create = "";
                return;
              default:
                this.err.create = HtmlSomethingWentWrong;
                return;
            }
          }
        })
        .catch(() => {
          this.captcha.show = false;
          this.err.create = HtmlSomethingWentWrong;
        });
    },
  },
  async beforeRouteUpdate(to) {
    this.loading = true;
    this.success = false;
    this.err = "";
    this.captcha.show = this.captcha.loaded = false;

    if (hasST(to.query.selector, to.query.token)) {
      validate(to.query.selector, to.query.token)
        .then((r) => {
          this.success = r.success;
          this.err = r.error;
        })
        .finally(() => {
          this.loading = false;
        });
    }
  },
  components: {
    SpinLoader,
    ToLink,
    TextSmall,
    EmailInput,
    CenterFormBox,
    Submit,
    InputLabel,
    VueHcaptcha,
    ChevronLeftIcon,
  },
});
</script>
