<template>
  <!--TODO:Image Resize-->
  <div>
    <div @click="toggleShow">
      <slot>
        <div class="group relative cursor-pointer">
          <div
            class="absolute hidden h-full w-full items-center justify-center rounded-full bg-black bg-opacity-50 group-hover:flex"
          >
            <CameraIcon class="h-5 w-5 text-white opacity-80"></CameraIcon>
          </div>
          <user-avatar :user="user" class="min-h-full min-w-full"></user-avatar>
        </div>
      </slot>
    </div>
    <overlay
      v-model="show"
      :block-scroll="blockScroll"
      :xs-full-height="false"
      :xs-self-end="true"
    >
      <box
        class="w-full self-end rounded-t-lg text-center xs:m-3 xs:self-center xs:rounded-md"
      >
        <snackbar v-model="err.show" :close-btn="true">{{ err.msg }}</snackbar>
        <div
          class="w-full border-b border-secondary-300 p-4 text-2xl font-medium text-black-lighter dark:border-secondary-700 dark:text-white"
        >
          Change Avatar
        </div>
        <div class="font-semibold">
          <div
            class="cursor-pointer border-b border-secondary-300 p-3 text-blue-500 dark:border-secondary-700"
          >
            <file-input
              class="w-full"
              accept="image/png,image/jpeg,image/gif,image/webp"
              :check="true"
              :max-size="maxImageSize"
              :required="true"
              @filesChange="avatarUpload"
              @error="onErr"
            >
              Upload Image
            </file-input>
          </div>
          <div
            v-if="user.avatar"
            class="cursor-pointer border-b border-secondary-300 p-3 text-red-500 dark:border-secondary-700"
            @click="avatarRemove"
          >
            Remove Image
          </div>
          <div
            class="cursor-pointer p-3 text-sm font-normal"
            @click="toggleShow"
          >
            Cancel
          </div>
        </div>
      </box>
    </overlay>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from "vue";
import { useSessionStore } from "@/store/session";

import { User } from "@/models/user";
import { RespToError } from "@/models/responses";
import { SomethingWentWrong } from "../form/errors";

import { avatar as UploadAvatar, removeAvatar } from "@/api/user/update/avatar";

import CameraIcon from "@heroicons/vue/solid/CameraIcon";

import Box from "../box/Box.vue";
import Overlay from "../overlay/Overlay.vue";
import UserAvatar from "./Avatar.vue";
import FileInput from "../form/FileInput.vue";
import Snackbar from "../box/Snackbar.vue";

export default defineComponent({
  props: {
    blockScroll: {
      type: Boolean,
      default: true,
    },
    user: {
      type: User,
      required: true,
    },
  },
  setup() {
    const session = useSessionStore(),
      maxImageSize = 5242880;

    const data = reactive({
      show: false,
      err: {
        msg: "",
        show: false,
      },
    });

    return {
      ...toRefs(data),
      updateStoreAvatar: (url: string | undefined) =>
        session.$patch({ user: { avatar: url } }),
      maxImageSize,
    };
  },
  watch: {
    modelValue(newValue) {
      this.show = newValue;
    },
  },
  methods: {
    toggleShow() {
      this.show = !this.show;
    },
    avatarUpload(avatar: File) {
      if (this.err.msg == "") {
        UploadAvatar(this.user.userName, avatar)
          .then((r) => {
            if (r.success) {
              this.updateStoreAvatar(r.avatarUrl);
              this.toggleShow();
            } else {
              var errText = RespToError(r.error);
              if (errText.length > 0) {
                this.onErr(errText);
              } else {
                this.onErr(SomethingWentWrong);
              }
            }
          })
          .catch(() => {
            this.onErr(SomethingWentWrong);
          });
      }
    },
    avatarRemove() {
      removeAvatar(this.user.userName)
        .then(() => {
          this.updateStoreAvatar("");
          this.toggleShow();
        })
        .catch(() => {
          this.onErr(SomethingWentWrong);
        });
    },
    onErr(err: string) {
      this.err.msg = "";
      if (err.length > 0) {
        this.err.msg = err;
        this.err.show = true;
      }
    },
  },
  components: { Overlay, Box, CameraIcon, UserAvatar, FileInput, Snackbar },
});
</script>
