<template>
  <overlay v-model="visable">
    <box
      class="flex h-full max-h-screen w-full flex-col rounded-none bg-secondary-800 shadow-lg focus:outline-none xs:m-3 xs:h-fit xs:rounded-md"
    >
      <snackbar v-model="err.snackbar.show" :close-btn="true">
        {{ err.snackbar.msg }}
      </snackbar>
      <div
        class="flex w-full shrink-0 items-end border-b border-secondary-600 p-3"
      >
        <div class="flex flex-1 justify-start">
          <basic-btn @click="visable = false">
            <XMarkIcon class="xIcon h-5 w-5"></XMarkIcon>
          </basic-btn>
        </div>
        <div class="flex flex-1 justify-center">
          <div
            class="text-base leading-4 font-medium text-neutral-900 dark:text-neutral-100"
          >
            Edit Profile
          </div>
        </div>
        <div class="flex flex-1 justify-end">
          <basic-btn class="link" @click="submit">Save</basic-btn>
        </div>
      </div>
      <div class="flex grow flex-col overflow-y-auto">
        <div class="flex shrink-0 grow flex-col px-6 py-6 xs:px-8">
          <div class="flex items-center">
            <user-avatar
              class="h-11 w-11"
              :user="user"
              :is-link="false"
            ></user-avatar>
            <change-avatar
              class="ml-2 w-full"
              :user="user"
              :block-scroll="false"
            >
              <basic-btn class="link">Change Avatar</basic-btn>
            </change-avatar>
          </div>
          <div class="mt-3.5">
            <checkbox v-model="privateUser" name="private">
              Private Account
            </checkbox>
          </div>
          <div class="mt-3">
            <input-label for="name" :err="err.name"> Name </input-label>
            <text-input
              v-model="name"
              name="name"
              :check="true"
              :min-len="0"
              :max-len="128"
              placeholder="Mary Smith"
              :regex="/^[^<>\[\]]*$/"
              :regex-msg="InvalidCharacters"
              @error="err.name = $event"
            ></text-input>
          </div>
          <div class="mt-3">
            <input-label for="bio" :err="err.bio"> Bio </input-label>
            <text-area
              v-model="bio"
              name="bio"
              :check="true"
              :min-len="0"
              :max-len="255"
              placeholder="Say Something?"
              :regex="/^[^<>\[\]]*$/"
              :regex-msg="InvalidCharacters"
              @error="err.bio = $event"
            ></text-area>
          </div>
        </div>
      </div>
    </box>
  </overlay>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from "vue";
import { useSessionStore } from "@/store/session";

import { User } from "@/models/user";
import { RespToError } from "@/models/responses";
import { InvalidCharacters, SomethingWentWrong } from "../form/errors";

import { bio as UpdateBio } from "@/api/user/update/bio";
import { name as UpdateName } from "@/api/user/update/name";
import {
  type as UpdateType,
  TYPE_PRIVATE,
  TYPE_PUBLIC,
} from "@/api/user/update/type";

import { XMarkIcon } from "@heroicons/vue/20/solid";

import Box from "../box/Box.vue";
import BasicBtn from "../button/BasicBtn.vue";
import InputLabel from "../form/InputLabel.vue";
import TextArea from "../form/TextArea.vue";
import TextInput from "../form/TextInput.vue";
import Overlay from "../overlay/Overlay.vue";
import UserAvatar from "./Avatar.vue";
import ChangeAvatar from "./ChangeAvatar.vue";
import Snackbar from "../box/Snackbar.vue";
import Checkbox from "../form/Checkbox.vue";

export default defineComponent({
  props: {
    user: {
      type: User,
      required: true,
    },
    show: {
      type: Boolean,
      default: false,
    },
  },
  emits: ["close"],
  setup(props) {
    const session = useSessionStore();
    const data = reactive({
      visable: props.show,
      name: props.user.name || "",
      bio: props.user.bio || "",
      privateUser: props.user.type == TYPE_PRIVATE,
      err: {
        name: "",
        bio: "",
        snackbar: {
          show: false,
          msg: "",
        },
      },
    });

    return {
      ...toRefs(data),
      InvalidCharacters,
      updateStoreBio: (bio: string) => session.$patch({ user: { bio } }),
      updateStoreName: (name: string) => session.$patch({ user: { name } }),
      updateStoreType: (type: number) => session.$patch({ user: { type } }),
    };
  },
  watch: {
    show(newValue) {
      this.visable = newValue;
    },
    visable(newValue) {
      if (!newValue) {
        this.close();
      }
    },
  },
  methods: {
    close() {
      this.$emit("close");
    },
    submit() {
      this.err.snackbar.msg = "";
      if (!this.noErrs()) {
        return;
      }

      (async () => {
        if (this.name != this.user.name) {
          UpdateName(this.user.userName, this.name)
            .then((r) => {
              if (r.success) {
                this.updateStoreName(this.name);
              } else {
                var errText = RespToError(r.error);
                if (errText.length > 0) {
                  this.err.name = errText;
                } else {
                  this.snackbarErr(SomethingWentWrong);
                }
              }
            })
            .catch(() => {
              this.snackbarErr(SomethingWentWrong);
            });
        }
        if (this.bio != this.user.bio) {
          UpdateBio(this.user.userName, this.bio)
            .then((r) => {
              if (r.success) {
                this.updateStoreBio(this.bio);
              } else {
                var errText = RespToError(r.error);
                if (errText.length > 0) {
                  this.err.bio = errText;
                } else {
                  this.snackbarErr(SomethingWentWrong);
                }
              }
            })
            .catch(() => {
              this.snackbarErr(SomethingWentWrong);
            });
        }
        if (this.privateUser != (this.user.type == TYPE_PRIVATE)) {
          var newType = this.privateUser ? TYPE_PRIVATE : TYPE_PUBLIC;
          UpdateType(this.user.userName, newType)
            .then((r) => {
              if (r.success) {
                this.updateStoreType(newType);
              } else {
                this.snackbarErr(SomethingWentWrong);
              }
            })
            .catch(() => {
              this.snackbarErr(SomethingWentWrong);
            });
        }
      })();

      if (this.noErrs()) {
        this.visable = false;
      }
    },
    noErrs() {
      if (
        this.err.name.length != 0 ||
        this.err.bio.length != 0 ||
        this.err.snackbar.msg.length != 0
      ) {
        return false;
      }
      return true;
    },
    snackbarErr(err: string) {
      this.err.snackbar.msg = "";
      if (err.length > 0) {
        this.err.snackbar.msg = err;
        this.err.snackbar.show = true;
      }
    },
  },
  components: {
    XMarkIcon,
    Box,
    BasicBtn,
    InputLabel,
    TextArea,
    TextInput,
    UserAvatar,
    ChangeAvatar,
    Overlay,
    Snackbar,
    Checkbox,
  },
});
</script>
