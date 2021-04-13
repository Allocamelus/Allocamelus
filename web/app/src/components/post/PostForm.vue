<template>
  <div class="flex-grow flex flex-col">
    <div class="flex flex-row">
      <div
        v-if="sanitize(richText).length == 0 && !focused"
        class="absolute select-none cursor-text text-lg opacity-90 p-1.5"
        @click="editor.focus()"
      >
        The Text...
      </div>
      <div ref="editor-div" class="flex-grow text-lg p-1.5 outline-none"></div>
    </div>
    <div class="sticky bottom-3 flex justify-between mt-2 bg-warm-gray-200 dark:bg-black-lighter p-1.5 rounded">
      <div class="flex items-center">
        <circle-bg
          v-for="(isActive, key, index) in active"
          :key="index"
          @click="btnClick(key)"
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
        <circle-bg class="ml-1.5">
          <file-input accept="image/png,image/jpeg,image/gif,image/webp">
            <radix-image class="w-5 h-5" />
          </file-input>
        </circle-bg>
      </div>
      <div class="flex items-center">
        <basic-btn class="text-secondary-700 dark:text-rose-600 p-1.5">Post</basic-btn>
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

function getValidator(str) {
  return new RegExp(`>${str}\\b`)
}

Squire.prototype.hasActionSelection = function (
  name,
  action,
  format
) {
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
        testBold: vm.editor.hasActionSelection("bold", action, "B", />B\b/),
        testItalic: vm.editor.hasActionSelection("italic", action, "I", />I\b/),
        testUnderline: vm.editor.hasActionSelection(
          "underline",
          action,
          "U"
        ),
        testOrderedList: vm.editor.hasActionSelection(
          "makeOrderedList",
          action,
          "OL"
        ),
        testLink: vm.editor.hasActionSelection("makeLink", action, "A", />A\b/),
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
  },
  mounted() {
    console.log("mount");
    this.editor = new Squire(this.$refs["editor-div"]);
    this.editor.addEventListener("input", this.onInput);
    this.editor.addEventListener("focus", () => (this.focused = true));
    this.editor.addEventListener("blur", () => (this.focused = false));
  },
  beforeUnmount() {
    console.log("unmount");
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
  },
});
</script>
<style scoped>
span {
  cursor: pointer;
  text-decoration: underline;
}
p {
  margin: 5px 0;
}
#editor {
  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  min-height: 200px;
  border: 1px solid #888;
  padding: 1em;
  background: transparent;
  color: #2b2b2b;
  font: 13px/1.35 Helvetica, arial, sans-serif;
  cursor: text;
}
a {
  text-decoration: underline;
}
h2 {
  font-size: 123.1%;
}
h3 {
  font-size: 108%;
}
h1,
h2,
h3,
p {
  margin: 1em 0;
}
h4,
h5,
h6 {
  margin: 0;
}
ul,
ol {
  margin: 0 1em;
  padding: 0 1em;
}
blockquote {
  border-left: 2px solid blue;
  margin: 0;
  padding: 0 10px;
}
pre {
  white-space: pre-wrap;
  word-wrap: break-word;
  overflow-wrap: break-word;
  border-radius: 3px;
  border: 1px solid #ccc;
  padding: 7px 10px;
  background: #f6f6f6;
  font-family: menlo, consolas, monospace;
  font-size: 90%;
}
code {
  border-radius: 3px;
  border: 1px solid #ccc;
  padding: 1px 3px;
  background: #f6f6f6;
  font-family: menlo, consolas, monospace;
  font-size: 90%;
}
</style>