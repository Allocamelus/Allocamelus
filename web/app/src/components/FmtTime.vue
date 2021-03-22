<template>
  <div>
    <router-link v-if="isLink" class="time" :to="link">
      {{ timeFmt }}
    </router-link>
    <div v-else class="time">
      {{ timeFmt }}
    </div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import { UnixTime } from "../pkg/time";
import FmtTime from "../pkg/fmtTime";
import FmtShort from "../pkg/fmtTime/sort";

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
  },
});
</script>

<style lang="scss" scoped>
@import "@/scss/vars";

.time {
  font-size: 15px;
  font-weight: 400;
  color: $light-text-12;
}

.dark-theme {
  .time {
    color: $dark-text-26;
  }
}
</style>