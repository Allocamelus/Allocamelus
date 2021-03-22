<template>
  <div>
    <div class="post-head mw-100">
      <div class="flex mw-100">
        <user-name :user="user" class="mw-100"></user-name>
        <fmt-time
          :time="post.published"
          class="dot-before post-time"
          :type="Fmt_Short_Time"
        ></fmt-time>
      </div>
      <div class="ml-3">t</div>
    </div>
    <div
      @click="textClick"
      :class="isLink ? 'pointer' : ''"
      v-html="post.content"
    ></div>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import UserName from "../user/Name.vue";
import FmtTime, { Fmt_Short_Time } from "../FmtTime.vue";

export default defineComponent({
  name: "post-box",
  props: {
    post: {
      type: Object,
    },
    user: {
      type: Object,
    },
    isLink: {
      type: Boolean,
      default: false,
    },
  },
  setup() {
    return {
      Fmt_Short_Time,
    };
  },
  computed: {
    link() {
      return `/post/${this.post.id}`;
    },
  },
  methods: {
    textClick() {
      if (this.isLink) {
        this.$router.push(this.link);
      }
    },
  },
  components: { FmtTime, UserName },
});
</script>

<style lang="scss" scoped>
@import "@/scss/vars";
.post-head {
  color: $light-text-8;
  &,
  .dot-before {
    display: flex;
    align-items: center;
  }
}
.post-time {
  white-space: nowrap;
}
.dot-before::before {
  font-weight: 400;
  font-size: 10px;
  margin: 0 4px;
}
.dark-theme {
  .post-head {
    color: $dark-text-26;
  }
}
</style>