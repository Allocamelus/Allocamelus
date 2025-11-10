<template>
  <div class="flex grow flex-col">
    <div class="flex flex-col">
      <div class="flex flex-row">
        <editor-content :editor="editor" class="grow" />
      </div>
      <div class="mt-1 flex flex-wrap overflow-hidden rounded-lg">
        <image-box
          v-for="(url, key) in imageUrls"
          :key="key"
          :index="key"
          :url="url"
          :total-number="images.length"
        >
          <div
            class="bg-opacity-50 absolute hidden h-full w-full flex-col justify-between bg-black p-2 text-white group-hover:flex"
          >
            <circle-bg
              class="h-6 w-6 self-end hover:bg-white"
              @click="removeImage(key)"
            >
              <XMarkIcon></XMarkIcon>
            </circle-bg>
            <div class="flex flex-col">
              <input-label :label="`imageAlt${key}`" :err="imageAltErrs[key]">
                Alt/Description
              </input-label>
              <text-input
                v-model="images[key].alt"
                :name="`imageAlt${key}`"
                :check="true"
                :max-len="512"
                :regex="altRegex"
                regex-msg="Some Characters will be escaped"
                @error="imageAltErrs[key] = $event"
              ></text-input>
            </div>
          </div>
        </image-box>
      </div>
    </div>
    <div class="sticky bottom-3 mt-2 overflow-hidden rounded">
      <div
        v-if="editor !== undefined"
        class="flex w-full flex-col bg-stone-200 p-1.5 dark:bg-black-lighter"
      >
        <div v-if="editor.isActive('link')" class="mb-1.5">
          <text-input
            v-model="link"
            :watch-model="true"
            type="url"
            name="link"
            placeholder="https://www.allocamelus.com"
          >
            <div class="mr-1.5 flex items-center">
              <basic-btn
                class="link p-1"
                title="Update Link"
                @click="updateLink"
              >
                Update
              </basic-btn>
            </div>
          </text-input>
        </div>
        <div class="flex justify-between">
          <div class="flex items-center">
            <circle-bg
              class="hover:bg-rose-800"
              :class="{
                'bg-secondary-700 text-stone-200': editor.isActive('bold'),
              }"
              @click="editor.commands.toggleBold()"
            >
              <RadixFontBold class="h-5 w-5" />
            </circle-bg>
            <circle-bg
              class="ml-1.5 hover:bg-rose-800"
              :class="{
                'bg-secondary-700 text-stone-200': editor.isActive('italic'),
              }"
              @click="editor.commands.toggleItalic()"
            >
              <RadixFontItalic class="h-5 w-5" />
            </circle-bg>
            <circle-bg
              class="ml-1.5 hover:bg-rose-800"
              :class="{
                'bg-secondary-700 text-stone-200': editor.isActive('underline'),
              }"
              @click="editor.commands.toggleUnderline()"
            >
              <RadixUnderline class="h-5 w-5" />
            </circle-bg>
            <circle-bg
              class="ml-1.5 hover:bg-rose-800"
              :class="{
                'bg-secondary-700 text-stone-200': editor.isActive('link'),
              }"
              @click="editor.commands.toggleLink()"
            >
              <RadixLink2 class="h-5 w-5" />
            </circle-bg>
            <circle-bg class="ml-1.5 hover:bg-rose-800">
              <file-input
                accept="image/png,image/jpeg,image/gif,image/webp"
                :check="true"
                :max-size="10485760 /* 10MB */"
                :max-files="4"
                :multiple="true"
                :file-count="images.length"
                @files-change="imagesUpload"
                @error="onErr"
              >
                <radix-image class="h-5 w-5" />
              </file-input>
            </circle-bg>
          </div>
          <div class="flex items-center">
            <basic-btn
              class="p-1.5 text-secondary-700 dark:text-rose-600"
              :disabled="submitted"
              @click="onPost"
            >
              Post
            </basic-btn>
          </div>
        </div>
      </div>
    </div>
    <snackbar v-model="err.show" :close-btn="true">{{ err.msg }}</snackbar>
  </div>
</template>

<script lang="ts" setup>
// TODO: Drag and drop & reorder images
import { defineComponent, toRefs, reactive, computed, ref } from "vue";
import Turndown from "turndown";

import { create as CreatePost, MediaFile } from "@/api/post/create";
import { notNull, RespToError } from "@/models/responses";
import { SomethingWentWrong } from "../form/errors";

import { textContent } from "@/pkg/sanitize";

import { useEditor, EditorContent } from "@tiptap/vue-3";
import Bold from "@tiptap/extension-bold";
import Document from "@tiptap/extension-document";
import History from "@tiptap/extension-history";
import Italic from "@tiptap/extension-italic";
import Link from "@tiptap/extension-link";
import Paragraph from "@tiptap/extension-paragraph";
import Placeholder from "@tiptap/extension-placeholder";
import Text from "@tiptap/extension-text";
import Underline from "@tiptap/extension-underline";

import RadixFontBold from "../icons/RadixFontBold.vue";
import RadixFontItalic from "../icons/RadixFontItalic.vue";
import RadixUnderline from "../icons/RadixUnderline.vue";
import FileInput from "../form/FileInput.vue";
import RadixImage from "../icons/RadixImage.vue";
import CircleBg from "../button/CircleBg.vue";
import BasicBtn from "../button/BasicBtn.vue";
import Snackbar from "../box/Snackbar.vue";
import { XMarkIcon } from "@heroicons/vue/20/solid";
import ImageBox from "../box/ImageBox.vue";
import TextInput from "../form/TextInput.vue";
import InputLabel from "../form/InputLabel.vue";
import RadixLink2 from "../icons/RadixLink2.vue";
import router from "@/router";

const turndownService = new Turndown().keep("u");
turndownService.addRule("pNewLine", {
  filter: "p",
  replacement: function (content) {
    return "\n\n" + content + "\n\n";
  },
});
const altRegex = /^[^<>[\]"&]*$/u;
const richText = ref("");
const link = ref("");
const focused = ref(false);
const images = reactive([] as MediaFile[]);
const imageAltErrs = reactive([] as string[]);
const imageUrls = reactive([] as string[]);
const submitted = ref(false);
const err = reactive({
  msg: "",
  show: false,
});
const hasNoText = computed(() => textContent(richText.value).length == 0);

const editor = useEditor({
  editorProps: {
    attributes: {
      class: "text-lg p-1.5 outline-none",
    },
  },
  onUpdate: ({ editor }) => {
    richText.value = editor.getHTML();
  },
  onSelectionUpdate: ({ editor }) => {
    if (editor.isActive("link")) {
      link.value = editor.getAttributes("link").href;
    } else {
      link.value = "";
    }
  },
  extensions: [
    Bold,
    Document,
    History,
    Italic,
    Link.configure({
      openOnClick: false,
      HTMLAttributes: {
        class: "link cursor-auto",
      },
    }),
    Paragraph.extend({
      name: "p",
    }),
    Placeholder.configure({
      emptyEditorClass: "placeholder-empty",
      placeholder: "The Text...",
    }),
    Text,
    Underline,
  ],
});

function updateLink() {
  if (editor.value === undefined) {
    return;
  }
  let newLink = link.value;

  if (newLink == "") {
    editor.value.chain().focus().extendMarkRange("link").unsetLink().run();
    return;
  }

  editor.value
    .chain()
    .focus()
    .extendMarkRange("link")
    .updateAttributes("link", {
      href: newLink,
    })
    .run();
}
function imagesUpload(imgs: File[]) {
  for (let i = 0; i < imgs.length; i++) {
    if (Object.hasOwnProperty.call(imgs, i)) {
      images.push(MediaFile.createFrom({ media: imgs[i], alt: "" }));
    }
  }
  imagesToUrl();
}
function onErr(errIn: string) {
  err.msg = "";
  if (errIn.length > 0) {
    err.msg = errIn;
    err.show = true;
  }
}
function removeImage(key: number) {
  images.splice(key, 1);
  imagesToUrl();
}
function imagesToUrl() {
  imageUrls.splice(0);
  for (let i = 0; i < images.length; i++) {
    imageUrls.push(URL.createObjectURL(images[i].media));
  }
}
function onPost() {
  if (submitted.value) {
    return onErr("Loading...");
  }
  if (hasNoText.value && images.length == 0) {
    return onErr("Text or Image(s) Required");
  }
  submitted.value = true;
  // TODO Limit content in browser
  CreatePost(turndownService.turndown(richText.value), images, true)
    .then((r) => {
      if ("id" in r) {
        return router.push(`/post/${r.id}`);
      }
      onPostErr(r.error);
    })
    .catch((e) => {
      onPostErr(e);
    });
}
function onPostErr(e?: string | any) {
  submitted.value = false;
  if (notNull(e)) {
    let errText = RespToError(e);
    onErr(errText);
  } else {
    onErr(SomethingWentWrong);
  }
}
</script>
