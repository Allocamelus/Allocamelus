<template>
  <overlay v-model="visable">
    <box
      class="w-full xs-max:h-full xs:m-3 rounded-none xs:rounded-md shadow-lg bg-secondary-800 focus:outline-none overflow-hidden flex flex-col"
    >
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
          <basic-btn>Save</basic-btn>
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
import { InvalidCharacters } from "../form/errors";

import XIcon from "@heroicons/vue/solid/XIcon";

import Box from "../box/Box.vue";
import BasicBtn from "../button/BasicBtn.vue";
import InputLabel from "../form/InputLabel.vue";
import TextArea from "../form/TextArea.vue";
import TextInput from "../form/TextInput.vue";
import Overlay from "../overlay/Overlay.vue";
import UserAvatar from "./Avatar.vue";
import ChangeAvatar from "./ChangeAvatar.vue";

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
    const data = reactive({
      visable: props.show,
      name: props.user.name,
      bio: props.user.bio,
      err: {
        name: "",
        bio: "",
      },
    });

    return {
      ...toRefs(data),
      InvalidCharacters,
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
  },
});
</script>