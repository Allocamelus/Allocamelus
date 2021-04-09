<template>
  <overlay v-model="visable">
    <box
      class="w-full xs-max:h-full xs:m-3 rounded-none xs:rounded-md shadow-lg bg-secondary-800 focus:outline-none overflow-hidden flex flex-col"
    >
      <snackbar v-model="err.snackbar.show" :closeBtn="true">
        {{ err.snackbar.msg }}
      </snackbar>
      <div class="w-full p-3 border-b border-secondary-600 flex items-end">
        <div class="flex-1 flex justify-start">
          <basic-btn @click="visable = false">
            <XIcon class="w-5 h-5"></XIcon>
          </basic-btn>
        </div>
        <div class="flex-1 flex justify-center">
          <div class="font-medium text-base leading-4">Edit Profile</div>
        </div>
        <div class="flex-1 flex justify-end">
          <basic-btn @click="submit">Save</basic-btn>
        </div>
      </div>
      <div class="flex-grow flex">
        <div class="flex flex-grow flex-col py-6 px-6 xs:px-8">
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
          <div class="mt-3">
            <input-label for="name" :err="err.name"> Name </input-label>
            <text-input
              v-model="name"
              name="name"
              :watchModel="true"
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
              :watchModel="true"
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
import ApiResp from "../../models/responses";
import { InvalidCharacters, SomethingWentWrong } from "../form/errors";

import { bio as UpdateBio } from "../../api/user/update/bio";
import { name as UpdateName } from "../../api/user/update/name";

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
      updateStoreName = (name) => store.commit("updateName", name);
    const data = reactive({
      visable: props.show,
      name: props.user.name,
      bio: props.user.bio,
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
    };
  },
  watch: {
    show(newValue, old) {
      this.visable = newValue;
    },
    visable(newValue, old) {
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
      var vm = this;
      vm.err.snackbar.msg = "";
      if (!vm.noErrs()) {
        return;
      }

      (async () => {
        if (vm.name != vm.user.name) {
          console.log(1);
          UpdateName(vm.user.userName, vm.name)
            .then((r) => {
              if (r.success) {
                vm.updateStoreName(vm.name);
              } else {
                switch (r.error) {
                  case ApiResp.User.Validate.Name.Length:
                    vm.err.name = "Invalid Length";
                    break;
                  case ApiResp.User.Validate.Invalid:
                    vm.err.name = InvalidCharacters;
                    break;
                  default:
                    vm.snackbarErr(SomethingWentWrong);
                    break;
                }
              }
            })
            .catch((e) => {
              vm.snackbarErr(SomethingWentWrong);
            });
        }
        if (vm.bio != vm.user.bio) {
          UpdateBio(vm.user.userName, vm.bio)
            .then((r) => {
              if (r.success) {
                vm.updateStoreBio(vm.bio);
              } else {
                switch (r.error) {
                  case ApiResp.User.Validate.Bio.Length:
                    vm.err.bio = "Invalid Length";
                    break;
                  case ApiResp.User.Validate.Invalid:
                    vm.err.bio = InvalidCharacters;
                    break;
                  default:
                    vm.snackbarErr(SomethingWentWrong);
                    break;
                }
              }
            })
            .catch((e) => {
              console.log(e);
              vm.snackbarErr(SomethingWentWrong);
            });
        }
      })();
      
      if (vm.noErrs()) {
        vm.visable = false;
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
  },
});
</script>