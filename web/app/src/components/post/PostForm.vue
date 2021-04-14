<template>
  <div class="flex-grow flex flex-col">
    <snackbar v-model="err.show" :closeBtn="true">{{ err.msg }}</snackbar>
    <div class="flex flex-col">
      <div class="flex flex-row">
        <div
          v-if="sanitize(richText).length == 0 && !focused"
          class="absolute select-none cursor-text text-lg opacity-90 p-1.5"
          @click="editor.focus()"
        >
          The Text...
        </div>
        <div
          ref="editor-div"
          class="flex-grow text-lg p-1.5 outline-none"
        ></div>
      </div>
      <div class="flex mt-1 flex-wrap">
        <div
          v-for="(url, key) in imageUrls"
          :key="key"
          class="group relative flex"
          :class="[
            images.length == 1 ? 'w-full' : '',
            images.length == 2 ? [key == 0 || key == 1 ? 'w-1/2' : ''] : '',
            images.length == 3
              ? [key == 0 ? 'w-full' : '', key == 1 || key == 2 ? 'w-1/2' : '']
              : '',
            images.length == 4 ? 'w-1/2' : '',
          ]"
        >
          <div
            class="absolute w-full h-full hidden group-hover:flex flex-col p-2 bg-black bg-opacity-50"
          >
            <circle-bg class="hover:bg-white w-6 h-6 self-end" @click="removeImage(key)">
              <XIcon class="text-white"></XIcon>
            </circle-bg>
          </div>
          <img :src="url" class="w-full object-cover" />
        </div>
      </div>
    </div>
    <div
      class="sticky bottom-3 flex justify-between mt-2 bg-warm-gray-200 dark:bg-black-lighter p-1.5 rounded"
    >
      <div class="flex items-center">
        <circle-bg
          v-for="(isActive, key, index) in active"
          :key="index"
          @click="btnClick(key)"
          class="hover:bg-rose-800"
          :class="[
            isActive ? 'bg-secondary-700 text-warm-gray-200' : '',
            index != 0 ? 'ml-1.5' : '',
          ]"
        >
          <component
            :is="`radix-${key != 'underline' ? 'font-' : ''}${key}`"
            class="w-5 h-5"
          />
        </circle-bg>
        <circle-bg class="ml-1.5 hover:bg-rose-800">
          <file-input
            accept="image/png,image/jpeg,image/gif,image/webp"
            :check="true"
            :maxSize="10485760 /* 10MB */"
            :maxFiles="4"
            :multiple="true"
            :fileCount="images.length"
            @filesChange="imagesUpload"
            @error="onErr"
          >
            <radix-image class="w-5 h-5" />
          </file-input>
        </circle-bg>
      </div>
      <div class="flex items-center">
        <basic-btn class="text-secondary-700 dark:text-rose-600 p-1.5">
          Post
        </basic-btn>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";

import sanitize from "../../pkg/sanitize";
import Squire from "squire-rte";

import RadixFontBold from "../icons/RadixFontBold.vue";
import RadixFontItalic from "../icons/RadixFontItalic.vue";
import RadixUnderline from "../icons/RadixUnderline.vue";
import FileInput from "../form/FileInput.vue";
import RadixImage from "../icons/RadixImage.vue";
import CircleBg from "../button/CircleBg.vue";
import BasicBtn from "../button/BasicBtn.vue";
import Snackbar from "../box/Snackbar.vue";
import XIcon from "@heroicons/vue/solid/XIcon";

function getValidator(str) {
  return new RegExp(`>${str}\\b`);
}

Squire.prototype.hasActionSelection = function (name, action, format) {
  var path = this.getPath(),
    test = getValidator(format).test(path) | this.hasFormat(format);
  if (name == action && test) {
    return true;
  } else {
    return false;
  }
};
Squire.prototype.thePath = function () {
  return this.getPath();
};

export default defineComponent({
  setup() {
    const data = reactive({
      editor: null,
      richText: "",
      focused: false,
      active: {
        bold: false,
        italic: false,
        underline: false,
      },
      images: [],
      imageUrls: [],
      err: {
        msg: "",
        show: false,
      },
    });
    return {
      ...toRefs(data),
      sanitize,
    };
  },
  methods: {
    btnClick(action) {
      var vm = this;
      console.log(vm.editor);
      var test = {
        value: action,
        testBold: vm.editor.hasActionSelection("bold", action, "B"),
        testItalic: vm.editor.hasActionSelection("italic", action, "I"),
        testUnderline: vm.editor.hasActionSelection("underline", action, "U"),
        testOrderedList: vm.editor.hasActionSelection(
          "makeOrderedList",
          action,
          "OL"
        ),
        testLink: vm.editor.hasActionSelection("makeLink", action, "A"),
        testQuote: vm.editor.hasActionSelection(
          "increaseQuoteLevel",
          action,
          "blockquote"
        ),
        isNotValue: function (a) {
          return a == action && this.value !== "";
        },
      };

      vm.editor.alignRight = function () {
        vm.editor.setTextAlignment("right");
      };
      vm.editor.alignCenter = function () {
        vm.editor.setTextAlignment("center");
      };
      vm.editor.alignLeft = function () {
        vm.editor.setTextAlignment("left");
      };
      vm.editor.alignJustify = function () {
        vm.editor.setTextAlignment("justify");
      };
      vm.editor.makeHeading = function () {
        vm.editor.setFontSize("2em");
        vm.editor.bold();
      };

      if (
        test.testBold |
        test.testItalic |
        test.testUnderline |
        test.testOrderedList |
        test.testLink |
        test.testQuote
      ) {
        if (test.testBold) {
          vm.editor.removeBold();
          vm.active.bold = false;
        }
        if (test.testItalic) {
          vm.editor.removeItalic();
          vm.active.italic = false;
        }
        if (test.testUnderline) {
          vm.editor.removeUnderline();
          vm.active.underline = false;
        }
        if (test.testLink) vm.editor.removeLink();
        if (test.testOrderedList) vm.editor.removeList();
        if (test.testQuote) vm.editor.decreaseQuoteLevel();
      } else if (
        test.isNotValue("makeLink") |
        test.isNotValue("insertImage") |
        test.isNotValue("selectFont")
      ) {
        // do nothing these are dropdowns.
      } else {
        vm.active[action] = true;
        vm.editor[action]();
        vm.editor.focus();
      }
      console.log(vm.editor.thePath());
      console.log(vm.editor);
    },
    onInput() {
      this.richText = this.editor.getHTML();
    },
    imagesUpload(images) {
      for (let i = 0; i < images.length; i++) {
        this.images.push(images[i]);
      }
      this.imagesToUrl();
    },
    onErr(err) {
      this.err.msg = "";
      if (err.length > 0) {
        this.err.msg = err;
        this.err.show = true;
      }
    },
    removeImage(key) {
      this.images.splice(key, 1);
      this.imagesToUrl()
    },
    imagesToUrl() {
      this.imageUrls = [];
      for (let i = 0; i < this.images.length; i++) {
        this.imageUrls.push(URL.createObjectURL(this.images[i]));
      }
    },
  },
  mounted() {
    this.editor = new Squire(this.$refs["editor-div"]);
    this.editor.addEventListener("input", this.onInput);
    this.editor.addEventListener("focus", () => (this.focused = true));
    this.editor.addEventListener("blur", () => (this.focused = false));
    // TODO: finish addEventListener s
  },
  beforeUnmount() {
    this.editor.destroy();
  },
  components: {
    RadixFontBold,
    RadixFontItalic,
    RadixUnderline,
    FileInput,
    RadixImage,
    CircleBg,
    BasicBtn,
    Snackbar,
    XIcon,
  },
});
</script>