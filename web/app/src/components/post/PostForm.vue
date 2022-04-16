<template>
  <div class="flex flex-grow flex-col">
    <div class="flex flex-col">
      <div class="flex flex-row">
        <editor-content :editor="editor" class="flex-grow" />
      </div>
      <div class="mt-1 flex flex-wrap overflow-hidden rounded-lg">
        <image-box
          v-for="(url, key) in imageUrls"
          :key="key"
          :index="key"
          :url="url"
          :totalNumber="images.length"
        >
          <div
            class="absolute hidden h-full w-full flex-col justify-between bg-black bg-opacity-50 p-2 text-white group-hover:flex"
          >
            <circle-bg
              class="h-6 w-6 self-end hover:bg-white"
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
    <div class="sticky bottom-3 mt-2 overflow-hidden rounded">
      <div
        class="flex w-full justify-between bg-warm-gray-200 p-1.5 dark:bg-black-lighter"
        v-if="editor !== undefined"
      >
        <div class="flex items-center">
          <circle-bg
            @click="editor.chain().focus().toggleBold().run()"
            class="hover:bg-rose-800"
            :class="{
              'bg-secondary-700 text-warm-gray-200': editor.isActive('bold'),
            }"
          >
            <RadixFontBold class="h-5 w-5" />
          </circle-bg>
          <circle-bg
            @click="editor.chain().focus().toggleItalic().run()"
            class="ml-1.5 hover:bg-rose-800"
            :class="{
              'bg-secondary-700 text-warm-gray-200': editor.isActive('italic'),
            }"
          >
            <RadixFontItalic class="h-5 w-5" />
          </circle-bg>
          <circle-bg
            @click="editor.chain().focus().toggleUnderline().run()"
            class="ml-1.5 hover:bg-rose-800"
            :class="{
              'bg-secondary-700 text-warm-gray-200':
                editor.isActive('underline'),
            }"
          >
            <RadixUnderline class="h-5 w-5" />
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
              <radix-image class="h-5 w-5" />
            </file-input>
          </circle-bg>
        </div>
        <div class="flex items-center">
          <basic-btn
            class="p-1.5 text-secondary-700 dark:text-rose-600"
            @click="onPost"
            :disabled="submitted"
          >
            Post
          </basic-btn>
        </div>
      </div>
    </div>
    <snackbar v-model="err.show" :closeBtn="true">{{ err.msg }}</snackbar>
  </div>
</template>

<script lang="ts">
// TODO: Drag and drop & reorder images
import { defineComponent, toRefs, reactive } from "vue";
import Turndown from "turndown";

import { create as CreatePost, MediaFile } from "../../api/post/create";
import { notNull, RespToError } from "../../models/responses";

import sanitize from "../../pkg/sanitize";

import { EditorContent, useEditor } from "@tiptap/vue-3";
import StarterKit from "@tiptap/starter-kit";
import Placeholder from "@tiptap/extension-placeholder";
import Underline from "@tiptap/extension-underline";

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
import { SomethingWentWrong } from "../form/errors";

const turndownService = new Turndown().keep("u");

export default defineComponent({
  setup() {
    const altRegex = /^[^<>\[\]"&]*$/u;

    const data = reactive({
      richText: "",
      focused: false,
      images: [] as MediaFile[],
      imageAltErrs: [] as string[],
      imageUrls: [] as string[],
      submitted: false,
      err: {
        msg: "",
        show: false,
      },
    });

    const editor = useEditor({
      editorProps: {
        attributes: {
          class: "text-lg p-1.5 outline-none",
        },
      },
      onUpdate: ({ editor }) => {
        data.richText = editor.getHTML();
      },
      extensions: [
        StarterKit,
        Underline,
        Placeholder.configure({
          emptyEditorClass: "placeholder-empty",
          placeholder: "The Text...",
        }),
      ],
    });

    return {
      ...toRefs(data),
      altRegex,
      editor,
    };
  },
  computed: {
    hasNoText() {
      var sanitized = sanitize(this.richText);
      return sanitized.length == 0;
    },
  },
  methods: {
    imagesUpload(images: File[]) {
      for (let i = 0; i < images.length; i++) {
        if (Object.hasOwnProperty.call(images, i)) {
          this.images.push(MediaFile.createFrom({ media: images[i], alt: "" }));
        }
      }
      this.imagesToUrl();
    },
    onErr(err: string) {
      this.err.msg = "";
      if (err.length > 0) {
        this.err.msg = err;
        this.err.show = true;
      }
    },
    removeImage(key: number) {
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
      if (this.submitted) {
        return this.onErr("Loading...");
      }
      if (this.hasNoText && this.images.length == 0) {
        return this.onErr("Text or Image(s) Required");
      }
      this.submitted = true;
      // TODO Limit content in browser
      CreatePost(turndownService.turndown(this.richText), this.images, true)
        .then((r) => {
          if ("id" in r) {
            return this.$router.push(`/post/${r.id}`);
          }
          this.onPostErr(r.error);
        })
        .catch((e) => {
          this.onPostErr(e);
        });
    },
    onPostErr(e?: string | any) {
      this.submitted = false;
      if (notNull(e)) {
        let errText = RespToError(e);
        this.onErr(errText);
      } else {
        this.onErr(SomethingWentWrong);
      }
    },
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
    EditorContent,
  },
});
</script>

<style scoped lang="scss">
:deep(.ProseMirror) p.placeholder-empty:first-child::before {
  content: attr(data-placeholder);
  @apply pointer-events-none float-left h-0 opacity-90;
}
</style>
