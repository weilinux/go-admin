import AuthService from "../services/auth.service";
import { Commit } from "vuex";

// const user = JSON.parse(localStorage.getItem("Token"));
const token = localStorage.getItem("Token")
const user = token ? JSON.parse(token): null;
const initialState = user
  ? { status: { loggedIn: true }, user }
  : { status: { loggedIn: false }, user: null };

export const auth = {
  namespaced: true,
  state: initialState,
  actions: {
    login({ commit}:{ commit: Commit }, user: any) {
      return AuthService.login(user).then(
        (user) => {
          commit("loginSuccess", user);
          return Promise.resolve(user);
        },
        (error) => {
          commit("loginFailure");
          return Promise.reject(error);
        }
      );
    },
    logout({ commit }: { commit: Commit }) {
      AuthService.logout();
      commit("logout");
    },
    register({ commit }: { commit: Commit }, user: any) {
      return AuthService.register(user).then(
        (response) => {
          commit("registerSuccess");
          return Promise.resolve(response.data);
        },
        (error) => {
          commit("registerFailure");
          return Promise.reject(error);
        }
      );
    },
  },
  mutations: {
    loginSuccess(state: any, user: any) {
      state.status.loggedIn = true;
      state.user = user;
    },
    loginFailure(state: any) {
      state.status.loggedIn = false;
      state.user = null;
    },
    logout(state: any) {
      state.status.loggedIn = false;
      state.user = null;
    },
    registerSuccess(state: any) {
      state.status.loggedIn = false;
    },
    registerFailure(state: any) {
      state.status.loggedIn = false;
    },
  },
};
