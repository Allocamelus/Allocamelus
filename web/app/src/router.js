import { createRouter, createWebHistory } from 'vue-router'
import store from "./store"

export function redirectUrl(redirect = "") {
  if (redirect?.length > 0) {
    return redirect;
  }
  return "/";
}

const routes = [
  {
    path: "/about",
    name: "About",
    component: () => import('./views/About.vue'),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import('./views/Login.vue'),
    props: route => ({ redirect: route.query.r })
  },
  {
    path: "/signup",
    name: "Signup",
    component: () => import('./views/SignUp.vue'),
    props: route => ({ redirect: route.query.r })
  },
  {
    path: "/logout",
    redirect: to => {
      store.dispatch("sessionLogout")
      return { path: redirectUrl(to.query.r), query: { ref: "logout" } }
    },
  },
  {
    path: "/account/verify-email",
    name: "Account Verify Email",
    component: () => import('./views/account/VerifyEmail.vue'),
    props: route => ({ selector: route.query.selector, token: route.query.token }),
  },
  {
    path: "/post/:id(\\d+)",
    component: () => import('./views/Post.vue'),
    props: true,
  },
  {
    path: "/post/new",
    name: "New Post",
    component: () => import('./views/post/New.vue'),
  },
  {
    path: "/",
    name: "Home",
    component: () => import('./views/Home.vue'),
  },
  {
    path: '/u/:userName(.*)*',
    name: 'User',
    component: () => import('./views/User.vue'),
    props: true
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'Error 404',
    component: () => import('./views/errors/404.vue')
  },
];

const router = createRouter({
  // Provide the history implementation to use. We are using the hash history for simplicity here.
  history: createWebHistory(),
  routes, // short for `routes: routes`
  scrollBehavior(_to, _from, savedPosition) {
    return savedPosition ? savedPosition : { top: 0 }
  },
})

router.beforeEach(async (to, _from) => {
  // canUserAccess() returns `true` or `false`
  if (store.getters.loggedIn) {
    switch (to.name) {
      case "Login":
      case "Signup":
      case "Account Verify Email":
        return redirectUrl(to.query.r)
    }
  } else {
    switch (to.name) {
      case "New Post":
        return "/login"
    }
  }
})

export default router