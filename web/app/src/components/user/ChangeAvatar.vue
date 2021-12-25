<template>
  <!--TODO:Image Resize-->
  <div>
    <div @click="toggleShow">
      <slot>
        <div class="group relative cursor-pointer">
          <div
            class="absolute w-full h-full rounded-full hidden group-hover:flex items-center justify-center bg-black bg-opacity-50"
          >
            <CameraIcon class="opacity-80 text-white w-5 h-5"></CameraIcon>
          </div>
          <user-avatar :user="user" class="min-w-full min-h-full"></user-avatar>
        </div>
      </slot>
    </div>
    <overlay
      v-model="show"
      :blockScroll="blockScroll"
      :xsFullHeight="false"
      :xsSelfEnd="true"
    >
      <box
        class="w-full xs:m-3 self-end xs:self-center rounded-t-lg xs:rounded-md text-center"
      >
        <snackbar v-model="err.show" :closeBtn="true">{{ err.msg }}</snackbar>
        <div
          class="w-full p-4 border-b text-black-lighter dark:text-white border-secondary-300 dark:border-secondary-700 text-2xl font-medium"
        >
          Change Avatar
        </div>
        <div class="font-semibold">
          <div
            class="cursor-pointer p-3 border-b border-secondary-300 dark:border-secondary-700 text-blue-500"
          >
            <file-input
              class="w-full"
              accept="image/png,image/jpeg,image/gif,image/webp"
              :check="true"
              :maxSize="maxImageSize"
              :required="true"
              @filesChange="avatarUpload"
              @error="onErr"
            >
              Upload Image
            </file-input>
          </div>
          <div
            class="cursor-pointer p-3 border-b border-secondary-300 dark:border-secondary-700 text-red-500"
            v-if="user.avatar"
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

<script>
import { defineComponent, reactive, toRefs } from "vue";
import { useStore } from "vuex";

import { User } from "../../models/user";
import { RespToError } from "../../models/responses";
import { SomethingWentWrong } from "../form/errors";

import {
  avatar as UploadAvatar,
  removeAvatar,
} from "../../api/user/update/avatar";

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
    const store = useStore(),
      updateStoreAvatar = (url) => store.commit("updateAvatar", url),
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
      updateStoreAvatar,
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
    avatarUpload(avatar) {
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
          hasErr = true;
          this.onErr(SomethingWentWrong);
        });
    },
    onErr(err) {
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
