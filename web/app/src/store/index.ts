import { VuexPersistence } from "vuex-persist";
import { createStore, useStore as baseUseStore, Store } from "vuex";
import { getters } from "./getters";

import { DaysToSec, MinToSec, UnixTime } from "../pkg/time";

import { keepAlive } from "../api/account/auth/keepAlive";
import { logout } from "../api/account/logout";
import { status } from "../api/account/auth/status";

import { User } from "../models/user";
import { InjectionKey } from "vue";

const vuexLocal = new VuexPersistence<State>({
  key: "sessionStore",
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

// define injection key
export const key: InjectionKey<Store<State>> = Symbol();

export class Session {
  loggedIn: boolean;
  user: User;
  fresh: boolean;
  created: number;
  expires: number;

  constructor(source: Partial<Session> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.loggedIn = source["loggedIn"] || false;
    this.user = new User(source["user"]);
    this.fresh = source["fresh"] || true;
    this.created = source["created"] || UnixTime();
    this.expires = source["expires"] || UnixTime(MinToSec(10));
  }
}

export const store = createStore<State>({
  state: {
    ui: {
      theme: "dark",
      // TODO: https://github.com/vuejs/vue-router/issues/974
      viewKey: 0,
    },
    session: new Session(),
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
        state.session.user.avatar = url;
      } else {
        state.session.user.avatar = "";
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
        session: new Session({
          loggedIn: true,
          user: payload.user,
          fresh: true,
          created: UnixTime(),
          expires: UnixTime(payload.authToken ? DaysToSec(30) : MinToSec(15)),
        }),
      });
    },
    sessionCheck({ commit, state }) {
      status()
        .then((st) => {
          if (!st.loggedIn) {
            if (state.session.loggedIn) {
              commit({
                type: "newSession",
                session: new Session(),
              });
            }
          } else {
            if (!state.session.loggedIn) {
              commit({
                type: "newSession",
                session: new Session({
                  loggedIn: true,
                  user: st.user,
                  fresh: state.session.fresh,
                  created: state.session.created,
                  expires: state.session.expires,
                }),
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
        session: new Session(),
      });
    },
  },
  getters,
  plugins: [vuexLocal.plugin],
});

export function useStore() {
  return baseUseStore(key);
}
