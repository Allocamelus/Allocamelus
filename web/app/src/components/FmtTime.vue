<template>
  <text-small
    :is="isLink ? 'to-link' : 'div'"
    class="font-normal"
    :to="link"
    :title="MDY_HM"
  >
    {{ timeFmt }}
  </text-small>
</template>

<script>
import { defineComponent } from "vue";
import { UnixTime } from "../pkg/time";
import FmtTime from "../pkg/fmtTime";
import FmtShort, { MDY_HM } from "../pkg/fmtTime/sort";

import ToLink from "./ToLink.vue";
import TextSmall from "./text/Small.vue";

export const Raw_Time = "raw";
export const Fmt_Time = "fmt";
export const Fmt_Short_Time = "fmt-short";
export const Iso_Time = "iso";

export default defineComponent({
  props: {
    time: {
      type: Number,
    },
    type: {
      type: String,
      default: Raw_Time,
    },
    link: {
      type: String,
      default: "",
    },
  },
  computed: {
    isLink() {
      if (this.link.length > 0) {
        return true;
      }
      return false;
    },
    timeC() {
      if (this.time == undefined) {
        return UnixTime();
      }
      return this.time;
    },
    timeFmt() {
      switch (this.type) {
        case Fmt_Time:
          return FmtTime(this.timeC);
        case Fmt_Short_Time:
          return FmtShort(this.timeC);
        case Iso_Time:
          return new Date(this.timeC * 1000).toISOString();
        default:
          return String(this.timeC);
      }
    },
    MDY_HM() {
      return MDY_HM(this.timeC);
    },
  },
  components: { ToLink, TextSmall },
});
</script>
