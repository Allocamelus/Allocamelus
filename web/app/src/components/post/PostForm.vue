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
          :total-number="images.length"
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
        v-if="editor !== null"
        class="flex w-full flex-col bg-warm-gray-200 p-1.5 dark:bg-black-lighter"
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
                'bg-secondary-700 text-warm-gray-200': editor.isActive('bold'),
              }"
              @click="editor.chain().focus().toggleBold().run()"
            >
              <RadixFontBold class="h-5 w-5" />
            </circle-bg>
            <circle-bg
              class="ml-1.5 hover:bg-rose-800"
              :class="{
                'bg-secondary-700 text-warm-gray-200':
                  editor.isActive('italic'),
              }"
              @click="editor.chain().focus().toggleItalic().run()"
            >
              <RadixFontItalic class="h-5 w-5" />
            </circle-bg>
            <circle-bg
              class="ml-1.5 hover:bg-rose-800"
              :class="{
                'bg-secondary-700 text-warm-gray-200':
                  editor.isActive('underline'),
              }"
              @click="editor.chain().focus().toggleUnderline().run()"
            >
              <RadixUnderline class="h-5 w-5" />
            </circle-bg>
            <circle-bg
              class="ml-1.5 hover:bg-rose-800"
              :class="{
                'bg-secondary-700 text-warm-gray-200': editor.isActive('link'),
              }"
              @click="editor.chain().focus().toggleLink().run()"
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

<script lang="ts">
// TODO: Drag and drop & reorder images
import { defineComponent, toRefs, reactive, computed } from "vue";
import Turndown from "turndown";

import { create as CreatePost, MediaFile } from "@/api/post/create";
import { notNull, RespToError } from "@/models/responses";
import { SomethingWentWrong } from "../form/errors";

import { textContent } from "@/pkg/sanitize";

import { Editor, EditorContent } from "@tiptap/vue-3";
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
import XIcon from "@heroicons/vue/solid/XIcon";
import ImageBox from "../box/ImageBox.vue";
import TextInput from "../form/TextInput.vue";
import InputLabel from "../form/InputLabel.vue";
import RadixLink2 from "../icons/RadixLink2.vue";

const turndownService = new Turndown().keep("u");

export default defineComponent({
  setup() {
    const altRegex = /^[^<>[\]"&]*$/u;

    const data = reactive({
      editor: null as Editor | null,
      richText: "",
      link: "",
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

    return {
      ...toRefs(data),
      altRegex,
      hasNoText: computed(() => textContent(data.richText).length == 0),
    };
  },
  mounted() {
    this.editor = new Editor({
      editorProps: {
        attributes: {
          class: "text-lg p-1.5 outline-none",
        },
      },
      onUpdate: ({ editor }) => {
        this.richText = editor.getHTML();
      },
      onSelectionUpdate: ({ editor }) => {
        if (editor.isActive("link")) {
          this.link = editor.getAttributes("link").href;
        } else {
          this.link = "";
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
  },
  beforeUnmount() {
    if (this.editor !== null) {
      this.editor.destroy();
    }
  },
  methods: {
    updateLink() {
      if (this.editor === null) {
        return;
      }
      let newLink = this.link;

      if (newLink == "") {
        this.editor.chain().focus().extendMarkRange("link").unsetLink().run();
        return;
      }

      this.editor
        .chain()
        .focus()
        .extendMarkRange("link")
        .updateAttributes("link", {
          href: newLink,
        })
        .run();
    },
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
    RadixLink2,
  },
});
</script>

<style scoped lang="scss">
:deep(.ProseMirror) p.placeholder-empty:first-child::before {
  content: attr(data-placeholder);
  @apply pointer-events-none float-left h-0 opacity-90;
}
</style>
