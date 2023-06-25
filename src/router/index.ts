import { createRouter, createWebHistory } from "vue-router";
import Home from '@/views/Home.vue';
import Main from '@/views/home/Main.vue';
import SayHi from "@/views/SayHi.vue";
import Users from "@/views/user/UserList.vue";
import UserEdit from "@/views/user/UserEdit.vue";
import UserNew from "@/views/user/UserNew.vue";
import User from "@/views/user/User.vue";
import Login from "@/views/user/UserLogin.vue";
import Logout from "@/views/user/UserLogout.vue";
import Register from "@/views/user/UserRegister.vue";
import Hosts from "@/views/hosts/HostList.vue";
import Host from "@/views/hosts/Host.vue";
import HostNew from "@/views/hosts/HostNew.vue";
import HostEdit from "@/views/hosts/HostEdit.vue";
import HostAuth from "@/views/hosts/HostAuth.vue";
import Database from "@/views/middleware/Database.vue";
import NotFound from "@/views/NotFound.vue";

// lazy-loaded
const Profile = () => import("../views/user/Profile.vue");

import Log from "@/views/log/Log.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    children: [
      {
        path: "/users",
        name: "Users",
        component: Users,
      },
      {
        path: "/404",
        name: "NotFound",
        component: NotFound,
      },
      {
        path: "/profile",
        name: "profile",
        // lazy-loaded
        component: Profile,
      },
      {
        path: "/user/:id",
        name: "User",
        component: User,
      },
      {
        path: "/user/edit/:id",
        name: "UserEdit",
        component: UserEdit,
      },
      {
        path: "/users/add",
        name: "UserNew",
        component: UserNew,
      },
      {
        path: "/hosts",
        name: "Hosts",
        meta: {
          requiresAuth: true,
        },
        component: Hosts,
      },
      {
        path: "/host/:id",
        name: "Host",
        component: Host,
      },
      {
        path: "/host/add",
        name: "HostNew",
        component: HostNew,
      },
      {
        path: "/host/edit/:id",
        name: "HostEdit",
        component: HostEdit,
      },
      {
        path: "/host/auth",
        name: "HostAuth",
        component: HostAuth,
      },
      {
        path: "/log",
        name: "Log",
        component: Log,
      },
      {
        path: "/main",
        name: "Main",
        component: Main,
      },
      {
        path: "/database",
        name: "Database",
        component: Database,
      },
      {
        path: "/about",
        name: "About",
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () =>
          import(/* webpackChunkName: "about" */ "@/views/About.vue"),
      },
    ],
  },
  {
    path: "/say_hi",
    name: "SayHi",
    component: SayHi,
  },
  {
    path: "/login",
    component: Login,
  },
  {
    path: "/logout",
    name: "Logout",
    component: Logout,
  },
  {
    path: "/register",
    name: "Register",
    component: Register,
  },
  {
    path: "/:pathMatch(.*)*",
    name: "404",
    redirect: "/404",
    // component: NotFound,
    // component: () => import("@/views/404.vue"),
  },

];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

// export default router;

// 判断用户是否已经登录
// router.beforeEach((to, from, next) => {
//   let token = localStorage.getItem("Token")
//
//   if (!token) {
//     next({
//       path: "/login",
//       query: {redirect: to.fullPath},
//     })
//   } else {
//     next()
//   }
//
// })

export default router;
