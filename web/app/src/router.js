import { createRouter, createWebHistory } from 'vue-router'

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

export default createRouter({
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