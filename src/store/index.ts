import { createStore } from "vuex";
// add auth module
import { auth } from "@/store/auth.module";

export default createStore({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    auth,
  },
});

// 别的模块想使用这个store也可以像下面这样export store
// const store = createStore({
//   modules: {
//     auth,
//   },
// });
//
// export default store;
