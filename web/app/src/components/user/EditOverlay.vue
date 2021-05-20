<template>
  <overlay v-model="visable">
    <box
      class="w-full xs-max:h-full xs:m-3 max-h-screen rounded-none xs:rounded-md shadow-lg bg-secondary-800 focus:outline-none flex flex-col"
    >
      <snackbar v-model="err.snackbar.show" :closeBtn="true">
        {{ err.snackbar.msg }}
      </snackbar>
      <div
        class="w-full p-3 border-b border-secondary-600 flex items-end flex-shrink-0"
      >
        <div class="flex-1 flex justify-start">
          <basic-btn @click="visable = false">
            <XIcon
              class="w-5 h-5 text-black dark:text-gray-100 hover:text-gray-600 dark:hover:text-gray-300"
            ></XIcon>
          </basic-btn>
        </div>
        <div class="flex-1 flex justify-center">
          <div
            class="font-medium text-base leading-4 text-gray-900 dark:text-gray-100"
          >
            Edit Profile
          </div>
        </div>
        <div class="flex-1 flex justify-end">
          <basic-btn @click="submit" class="link">Save</basic-btn>
        </div>
      </div>
      <div class="flex-grow flex flex-col overflow-y-auto">
        <div class="flex flex-grow flex-shrink-0 flex-col py-6 px-6 xs:px-8">
          <div class="flex items-center">
            <user-avatar
              class="h-11 w-11"
              :user="user"
              :isLink="false"
            ></user-avatar>
            <change-avatar
              class="w-full ml-2"
              :user="user"
              :blockScrool="false"
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
              :minLen="0"
              :maxLen="128"
              placeholder="Mary Smith"
              :regex="/^[^<>\[\]]*$/"
              :regexMsg="InvalidCharacters"
              @error="err.name = $event"
            ></text-input>
          </div>
          <div class="mt-3">
            <input-label for="bio" :err="err.bio"> Bio </input-label>
            <text-area
              v-model="bio"
              name="bio"
              :check="true"
              :minLen="0"
              :maxLen="255"
              placeholder="Say Something?"
              :regex="/^[^<>\[\]]*$/"
              :regexMsg="InvalidCharacters"
              @error="err.bio = $event"
            ></text-area>
          </div>
        </div>
      </div>
    </box>
  </overlay>
</template>

<script>
import { defineComponent, reactive, toRefs } from "vue";

import { GEN_User } from "../../models/go_structs_gen";
import { RespToError } from "../../models/responses";
import { InvalidCharacters, SomethingWentWrong } from "../form/errors";

import { bio as UpdateBio } from "../../api/user/update/bio";
import { name as UpdateName } from "../../api/user/update/name";
import {
  type as UpdateType,
  TYPE_PRIVATE,
  TYPE_PUBLIC,
} from "../../api/user/update/type";

import XIcon from "@heroicons/vue/solid/XIcon";

import Box from "../box/Box.vue";
import BasicBtn from "../button/BasicBtn.vue";
import InputLabel from "../form/InputLabel.vue";
import TextArea from "../form/TextArea.vue";
import TextInput from "../form/TextInput.vue";
import Overlay from "../overlay/Overlay.vue";
import UserAvatar from "./Avatar.vue";
import ChangeAvatar from "./ChangeAvatar.vue";
import { useStore } from "vuex";
import Snackbar from "../box/Snackbar.vue";
import Checkbox from "../form/Checkbox.vue";

export default defineComponent({
  props: {
    user: {
      type: GEN_User,
      required: true,
    },
    show: {
      type: Boolean,
      default: false,
    },
  },
  emits: ["close"],
  setup(props) {
    const store = useStore(),
      updateStoreBio = (bio) => store.commit("updateBio", bio),
      updateStoreName = (name) => store.commit("updateName", name),
      updateStoreType = (type) => store.commit("updateType", type);
    const data = reactive({
      visable: props.show,
      name: props.user.name,
      bio: props.user.bio,
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
      updateStoreBio,
      updateStoreName,
      updateStoreType,
    };
  },
  watch: {
    show(newValue, _old) {
      this.visable = newValue;
    },
    visable(newValue, _old) {
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
            .catch((_e) => {
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
            .catch((_e) => {
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
            .catch((_e) => {
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
    snackbarErr(err) {
      this.err.snackbar.msg = "";
      if (err.length > 0) {
        this.err.snackbar.msg = err;
        this.err.snackbar.show = true;
      }
    },
  },
  components: {
    XIcon,
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