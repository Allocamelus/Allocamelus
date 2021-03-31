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
  }
];

const router = createRouter({
  // Provide the history implementation to use. We are using the hash history for simplicity here.
  history: createWebHistory(),
  routes, // short for `routes: routes`
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  },
})

router.beforeEach(async (to, from) => {
  // canUserAccess() returns `true` or `false`
  if (store.getters.loggedIn) {
    switch (to.name) {
      case "Login":
      case "Signup":
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