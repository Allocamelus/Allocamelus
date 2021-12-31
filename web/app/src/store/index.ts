import { VuexPersistence } from "vuex-persist";
import { createStore } from "vuex";
import { DaysToSec, MinToSec, UnixTime } from "../pkg/time";

import { keepAlive } from "../api/account/auth/keepAlive";
import { logout } from "../api/account/logout";
import { status } from "../api/account/auth/status";

import { User } from "../models/user";

const vuexLocal = new VuexPersistence<State>({
  key: "a10storage",
  storage: window.localStorage,
  reducer: (state) => {
    let storage = <State>{
      ui: {
        theme: state.ui.theme,
      },
    };
    if (state.session.expires > UnixTime()) {
      if (state.session.fresh === true) {
        state.session.fresh = false;
      }
      storage.session = state.session;
    }
    return storage;
  },
});

export interface State {
  ui: {
    theme: string;
    viewKey: number;
  };
  session: Session;
}

export interface Session {
  loggedIn: boolean;
  user: User;
  fresh: boolean;
  created: number;
  expires: number;
}

function sessionDefault(): Session {
  return {
    loggedIn: false,
    user: new User(),
    fresh: true,
    created: UnixTime(),
    expires: UnixTime(MinToSec(10)),
  };
}

export default createStore({
  state: <State>{
    ui: {
      theme: "dark",
      // TODO: https://github.com/vuejs/vue-router/issues/974
      viewKey: 0,
    },
    session: sessionDefault(),
  },
  mutations: {
    newSession(state, payload) {
      state.session = payload.session;
    },
    usedSession(state) {
      if (state.session.expires < UnixTime(MinToSec(15))) {
        state.session.expires = UnixTime(MinToSec(15));
      }
    },
    toggleTheme(state) {
      state.ui.theme = state.ui.theme === "dark" ? "light" : "dark";
    },
    updateViewKey(state) {
      state.ui.viewKey++;
    },
    updateAvatar(state, url) {
      if (url?.length > 0) {
        state.session.user.avatar = true;
        state.session.user.avatarUrl = url;
      } else {
        state.session.user.avatar = false;
        state.session.user.avatarUrl = undefined;
      }
    },
    updateBio(state, bio) {
      state.session.user.bio = bio;
    },
    updateName(state, name) {
      state.session.user.name = name;
    },
    updateType(state, type) {
      state.session.user.type = type;
    },
  },
  actions: {
    newLoginSession({ commit }, payload) {
      commit({
        type: "newSession",
        session: {
          loggedIn: true,
          user: payload.user,
          fresh: true,
          created: UnixTime(),
          expires: UnixTime(payload.authToken ? DaysToSec(30) : MinToSec(15)),
        },
      });
    },
    sessionCheck({ commit, state }) {
      status()
        .then((st) => {
          if (st.loggedIn === false) {
            if (state.session.loggedIn) {
              commit({
                type: "newSession",
                session: sessionDefault(),
              });
            }
          } else {
            if (!state.session.loggedIn) {
              commit({
                type: "newSession",
                session: {
                  loggedIn: true,
                  user: st.user,
                  fresh: state.session.fresh,
                  created: state.session.created,
                  expires: state.session.expires,
                },
              });
            }
          }
          commit("usedSession");
        })
        .catch(() => {
          // TODO
        });
    },
    sessionKeepAlive({ commit }) {
      keepAlive().then(() => {
        commit("usedSession");
      });
    },
    sessionLogout({ commit, state }) {
      if (state.session.loggedIn) {
        logout();
      }
      commit({
        type: "newSession",
        session: sessionDefault(),
      });
    },
  },
  getters: {
    loggedIn(state) {
      if (state.session.expires < UnixTime()) {
        return false;
      }
      return state.session.loggedIn;
    },
    user(state) {
      return new User(state.session.user);
    },
    theme(state) {
      return state.ui.theme;
    },
    viewKey(state) {
      return state.ui.viewKey;
    },
  },
  plugins: [vuexLocal.plugin],
});