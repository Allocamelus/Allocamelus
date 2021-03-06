<template>
  <div class="flex-grow flex flex-col">
    <snackbar v-model="err.show" :closeBtn="true">{{ err.msg }}</snackbar>
    <div class="flex flex-col">
      <div class="flex flex-row">
        <div
          v-if="hasNoText && !focused"
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
      <div class="flex mt-1 flex-wrap rounded-lg overflow-hidden">
        <image-box
          v-for="(url, key) in imageUrls"
          :key="key"
          :index="key"
          :url="url"
          :totalNumber="images.length"
        >
          <div
            class="absolute w-full h-full hidden group-hover:flex flex-col justify-between p-2 bg-black bg-opacity-50 text-white"
          >
            <circle-bg
              class="hover:bg-white w-6 h-6 self-end"
              @click="removeImage(key)"
            >
              <XIcon></XIcon>
            </circle-bg>
            <div class="flex flex-col">
              <input-label :label="`imageAlt${key}`" :err="imageAltErrs[key]">
                Alt/Description
              </input-label>
              <text-input
                v-model="images[key].alt"
                :name="`imageAlt${key}`"
                :check="true"
                :maxLen="512"
                :regex="altRegex"
                regexMsg="Some Characters will be escaped"
                @error="imageAltErrs[key] = $event"
              ></text-input>
            </div>
          </div>
        </image-box>
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
        <basic-btn
          class="text-secondary-700 dark:text-rose-600 p-1.5"
          @click="onPost"
          :disabled="submited"
        >
          Post
        </basic-btn>
      </div>
    </div>
  </div>
</template>

<script>
// TODO: Drag and drop & reorder images
import { defineComponent, toRefs, reactive } from "vue";
import Turndown from "turndown";

import { create as CreatePost, MediaFile } from "../../api/post/create";

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
import ImageBox from "../box/ImageBox.vue";
import TextInput from "../form/TextInput.vue";
import InputLabel from "../form/InputLabel.vue";

function getValidator(str) {
  return new RegExp(`>${str}\\b`, 'u');
}

Squire.prototype.hasActionSelection = function (name, action, format) {
  var path = this.getPath(),
    test = getValidator(format).test(path) | this.hasFormat(format);
  return name == action && test ? true : false;
};
Squire.prototype.thePath = () => {
  return this.getPath();
};

const turndownService = new Turndown().keep("u");

export default defineComponent({
  setup() {
    const altRegex = /^[^<>\[\]"&]*$/u;
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
      imageAltErrs: [],
      imageUrls: [],
      submited: false,
      err: {
        msg: "",
        show: false,
      },
    });
    return {
      ...toRefs(data),
      altRegex,
    };
  },
  computed: {
    hasNoText() {
      var sanitized = sanitize(this.richText);
      return sanitized.length == 0;
    },
  },
  methods: {
    btnClick(action) {
      var test = {
        value: action,
        testBold: this.editor.hasActionSelection("bold", action, "B"),
        testItalic: this.editor.hasActionSelection("italic", action, "I"),
        testUnderline: this.editor.hasActionSelection("underline", action, "U"),
        testOrderedList: this.editor.hasActionSelection(
          "makeOrderedList",
          action,
          "OL"
        ),
        testLink: this.editor.hasActionSelection("makeLink", action, "A"),
        testQuote: this.editor.hasActionSelection(
          "increaseQuoteLevel",
          action,
          "blockquote"
        ),
        isNotValue: (a) => {
          return a == action && this.value !== "";
        },
      };

      this.editor.alignRight = () => {
        this.editor.setTextAlignment("right");
      };
      this.editor.alignCenter = () => {
        this.editor.setTextAlignment("center");
      };
      this.editor.alignLeft = () => {
        this.editor.setTextAlignment("left");
      };
      this.editor.alignJustify = () => {
        this.editor.setTextAlignment("justify");
      };
      this.editor.makeHeading = () => {
        this.editor.setFontSize("2em");
        this.editor.bold();
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
          this.editor.removeBold();
          this.active.bold = false;
        }
        if (test.testItalic) {
          this.editor.removeItalic();
          this.active.italic = false;
        }
        if (test.testUnderline) {
          this.editor.removeUnderline();
          this.active.underline = false;
        }
        if (test.testLink) this.editor.removeLink();
        if (test.testOrderedList) this.editor.removeList();
        if (test.testQuote) this.editor.decreaseQuoteLevel();
      } else if (
        test.isNotValue("makeLink") |
        test.isNotValue("insertImage") |
        test.isNotValue("selectFont")
      ) {
        // do nothing these are dropdowns.
      } else {
        this.active[action] = true;
        this.editor[action]();
        this.editor.focus();
      }
    },
    onInput() {
      this.richText = this.editor.getHTML();
    },
    imagesUpload(images) {
      for (let i = 0; i < images.length; i++) {
        this.images.push(MediaFile.createFrom({ media: images[i], alt: "" }));
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
      this.imagesToUrl();
    },
    imagesToUrl() {
      this.imageUrls = [];
      for (let i = 0; i < this.images.length; i++) {
        this.imageUrls.push(URL.createObjectURL(this.images[i].media));
      }
    },
    onPost() {
      if (this.submited) {
        return this.onErr("Loading...");
      }
      if (this.hasNoText && this.images.length == 0) {
        return this.onErr("Text or Image(s) Required");
      }
      this.submited = true;
      CreatePost(turndownService.turndown(this.richText), this.images, true)
        .then((r) => {
          if (r.success) {
            return this.$router.push(`/post/${r.id}`);
          }
          this.onPostErr(r.error);
        })
        .catch((e) => {
          this.onPostErr(e);
        });
    },
    onPostErr(e) {
      this.submited = false;
      var errText = RespToError(e);
      if (errText.length > 0) {
        this.onErr(errText);
      } else {
        this.onErr(SomethingWentWrong);
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
    ImageBox,
    TextInput,
    InputLabel,
  },
});
</script>