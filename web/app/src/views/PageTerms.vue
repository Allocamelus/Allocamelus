<template>
  <div class="container flex flex-col py-5">
    <Box class="flex flex-col gap-6 rounded-xl p-5 text-base">
      <h1 class="h2">Terms of Service</h1>
      <p>Last Updated: November 9th, 2025</p>

      <div>
        <p>
          Welcome to Allocamelus! These Terms of Service ("Terms") govern your
          use of the Allocamelus website and services (collectively, the
          "Service"), so please read them carefully.
        </p>
      </div>

      <div>
        <h2 id="acceptance-of-terms" class="h3 mb-3">
          <a href="#acceptance-of-terms">Acceptance of Terms</a>
        </h2>
        <p>
          By using the Service, you agree to be bound by these Terms. If you do
          not agree to these Terms, you may not use the Service.
        </p>
      </div>

      <div>
        <h2 id="description-of-service" class="h3 mb-3">
          <a href="#description-of-service">Description of Service</a>
        </h2>
        <p>
          Allocamelus provides a platform for users to post text and images. We
          may update, modify, or discontinue any part of the Service at any
          time.
        </p>
      </div>

      <div>
        <h2 id="user-accounts" class="h3 mb-3">
          <a href="#user-accounts">User Accounts</a>
        </h2>
        <p>
          You are responsible for maintaining the security of your account and
          for all activities that occur under your account. You must immediately
          notify us of any unauthorized use of your account.
        </p>
      </div>

      <div>
        <h2 id="user-content" class="h3 mb-3">
          <a href="#user-content">User Content</a>
        </h2>
        <p>
          You are solely responsible for the content you post ("User Content").
          By posting User Content, you grant us a worldwide, non-exclusive,
          royalty-free license to use, reproduce, modify, and distribute your
          User Content in connection with the Service and for our marketing and
          promotional purposes.
        </p>
      </div>

      <div>
        <h2 id="prohibited-activities" class="h3 mb-3">
          <a href="#prohibited-activities">Prohibited Activities</a>
        </h2>
        <p>
          You agree not to post any User Content that is illegal or unlawful.
        </p>
      </div>

      <div>
        <h2 id="termination" class="h3 mb-3">
          <a href="#termination">Termination</a>
        </h2>
        <p>
          We may terminate or suspend your account at any time, for any reason,
          without notice.
        </p>
      </div>

      <div>
        <h2 id="disclaimer-of-warranties" class="h3 mb-3">
          <a href="#disclaimer-of-warranties">Disclaimer of Warranties</a>
        </h2>
        <p>
          The Service is provided "as is," without any warranties of any kind.
          We do not warrant that the Service will be error-free or
          uninterrupted.
        </p>
      </div>

      <div>
        <h2 id="limitation-of-liability" class="h3 mb-3">
          <a href="#limitation-of-liability">Limitation of Liability</a>
        </h2>
        <p>
          We will not be liable for any damages arising from your use of the
          Service.
        </p>
      </div>

      <div>
        <h2 id="governing-law" class="h3 mb-3">
          <a href="#governing-law">Governing Law</a>
        </h2>
        <p>
          These Terms are governed by the laws of the Commonwealth of Kentucky,
          without regard to its conflict of law provisions.
        </p>
      </div>

      <div>
        <h2 id="changes-to-terms" class="h3 mb-3">
          <a href="#changes-to-terms">Changes to Terms</a>
        </h2>
        <p>
          We may modify these Terms at any time. We will notify you of any
          changes by posting the new Terms on this page.
        </p>
      </div>

      <div>
        <h2 id="contact-us" class="h3 mb-3">
          <a href="#contact-us">Contact Us</a>
        </h2>
        <p>
          If you have any questions about these Terms, please contact us at
          {{ contactEmail }}.
        </p>
      </div>
    </Box>
  </div>
</template>

<script>
import { defineComponent, toRefs, reactive } from "vue";
import Box from "@/components/box/Box.vue";
export default defineComponent({
  setup() {
    const data = reactive({});

    document.title = `Terms of Service - ${import.meta.env.VITE_SITE_NAME}`;

    return {
      ...toRefs(data),
      contactEmail: import.meta.env.VITE_CONTACT_EMAIL,
    };
  },
  components: { Box },
});
</script>
