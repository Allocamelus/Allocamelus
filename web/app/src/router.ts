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
  if (Array.isArray(redirect) || redirect == null) {
    redirect = "";
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
    component: () => import("./views/About.vue"),
    meta: { footer: true },
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("./views/Login.vue"),
    props: (route) => ({ redirect: route.query.r }),
    meta: { footer: true },
  },
  {
    path: "/signup",
    name: "Signup",
    component: () => import("./views/SignUp.vue"),
    props: (route) => ({ redirect: route.query.r }),
    meta: { footer: true },
  },
  {
    path: "/logout",
    redirect: (to) => {
      const session = useSessionStore();
      session.logout();
      return { path: redirectUrl(to.query.r), query: { ref: "logout" } };
    },
  },
  {
    path: "/account/verify-email",
    name: "Account Verify Email",
    component: () => import("./views/account/VerifyEmail.vue"),
    props: (route) => ({
      selector: route.query.selector,
      token: route.query.token,
    }),
    meta: { footer: true },
  },
  {
    path: "/post/:id(\\d+)",
    component: () => import("./views/Post.vue"),
    props: true,
    meta: { footer: true },
  },
  {
    path: "/post/new",
    name: "New Post",
    component: () => import("./views/post/New.vue"),
    meta: { footer: true },
  },
  {
    path: "/",
    name: "Landing",
    component: () => import("./views/Landing.vue"),
    meta: { footer: true },
  },
  {
    path: "/home",
    name: "Home",
    component: () => import("./views/Home.vue"),
  },
  {
    path: "/u/:userName(.*)",
    name: "User",
    component: () => import("./views/User.vue"),
    props: true,
  },
  {
    path: "/:pathMatch(.*)*",
    name: "Error 404",
    component: () => import("./views/errors/404.vue"),
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
      case "Account Verify Email":
        return redirectUrl(to.query.r);
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
