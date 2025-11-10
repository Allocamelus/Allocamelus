import {
  createRouter,
  createWebHistory,
  LocationQueryValue,
  RouteRecordRaw,
} from "vue-router";
import { useSessionStore } from "@/store/session";

export function redirectUrl(
  redirect: LocationQueryValue | LocationQueryValue[]
) {
  if (Array.isArray(redirect)) {
    redirect = redirect[0];
  }
  if (redirect === null) {
    redirect = "/";
  }
  if (redirect.length > 0) {
    return redirect;
  }
  return "/";
}

const routes: Array<RouteRecordRaw> = [
  {
    path: "/about",
    name: "About",
    component: () => import("./views/PageAbout.vue"),
    meta: { footer: true },
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("./views/PageLogin.vue"),
    props: (route) => ({ redirect: route.query.r }),
    meta: { footer: true },
  },
  {
    path: "/signup",
    name: "Signup",
    component: () => import("./views/PageSignUp.vue"),
    props: (route) => ({ redirect: route.query.r }),
    meta: { footer: true },
  },
  {
    path: "/logout",
    redirect: (to) => {
      const session = useSessionStore();
      session.logout();
      return { path: redirectUrl(to.query.r || "/"), query: { ref: "logout" } };
    },
  },
  {
    path: "/account/verify-email",
    name: "Account Verify Email",
    component: () => import("./views/account/AccountVerifyEmail.vue"),
    props: (route) => ({
      selector: route.query.selector,
      token: route.query.token,
    }),
    meta: { footer: true },
  },
  {
    path: "/post/:id(\\d+)",
    component: () => import("./views/PagePost.vue"),
    props: true,
    meta: { footer: true },
  },
  {
    path: "/post/new",
    name: "New Post",
    component: () => import("./views/post/PostNew.vue"),
    meta: { footer: true },
  },
  {
    path: "/",
    name: "Landing",
    component: () => import("./views/PageLanding.vue"),
    meta: { footer: true },
  },
  {
    path: "/home",
    name: "Home",
    component: () => import("./views/PageHome.vue"),
  },
  {
    path: "/terms",
    name: "Terms",
    component: () => import("./views/PageTerms.vue"),
    meta: { footer: true },
  },
  {
    path: "/privacy",
    name: "Privacy",
    component: () => import("./views/PagePrivacy.vue"),
    meta: { footer: true },
  },
  {
    path: "/u/:userName(.*)",
    name: "User",
    component: () => import("./views/PageUser.vue"),
    props: true,
  },
  {
    path: "/:pathMatch(.*)*",
    name: "Error 404",
    component: () => import("./views/errors/ErrorPage404.vue"),
    meta: { footer: true },
  },
];

const router = createRouter({
  // Provide the history implementation to use. We are using the hash history for simplicity here.
  history: createWebHistory(),
  routes, // short for `routes: routes`
  // skipcq: JS-0356
  scrollBehavior(_to, _from, savedPosition) {
    return savedPosition ? savedPosition : { top: 0 };
  },
});

router.beforeResolve((to) => {
  const session = useSessionStore();

  // canUserAccess() returns `true` or `false`
  if (session.loggedIn) {
    switch (to.name) {
      case "Login":
      case "Signup":
      case "Account Verify Email": {
        const rdr = redirectUrl(to.query.r);
        if (rdr === "/") {
          return "/home";
        }
        return rdr;
      }
      case "Landing":
        return "/home";
      default:
        break;
    }
  } else {
    switch (to.name) {
      case "New Post":
        return "/login";
      case "Home":
        return "/";
      default:
        break;
    }
  }
});

export default router;
