import { createStore } from 'vuex'
import VuexPersistence from 'vuex-persist'
import { UnixTime, MinToSec, DaysToSec } from "./pkg/time";

import { status } from "./api/account/auth/status"
import { keepAlive } from "./api/account/auth/keepAlive"
import { logout } from "./api/account/logout"

import { GEN_User } from "./models/go_structs_gen"

const vuexLocal = new VuexPersistence({
  key: "a10storage",
  storage: window.localStorage,
  reducer: (state) => {
    var storage = {
      ui: {
        theme: state.ui.theme
      }
    }
    if (state.session.expires > UnixTime()) {
      if (state.session.fresh == true) {
        state.session.fresh = false
      }
      storage.session = state.session
    }
    return storage
  },
})

function sessionDefault() {
  return {
    loggedIn: false,
    user: new GEN_User(),
    fresh: true,
    created: UnixTime(),
    expires: UnixTime(MinToSec(10))
  }
}

export default createStore({
  state: {
    ui: {
      theme: 'dark',
      // TODO: https://github.com/vuejs/vue-router/issues/974
      viewKey: 0,
    },
    session: sessionDefault()
  },
  mutations: {
    newSession(state, payload) {
      state.session = payload.session
    },
    usedSession(state) {
      if (state.session.expires < UnixTime(MinToSec(15))) {
        state.session.expires = UnixTime(MinToSec(15))
      }
    },
    toggleTheme(state) {
      state.ui.theme = (state.ui.theme == 'dark') ? 'light' : 'dark'
    },
    updateViewKey(state) {
      state.ui.viewKey++
    },
    updateAvatar(state, url) {
      if (url?.length > 0) {
        state.session.user.avatar = true;
        state.session.user.avatarUrl = url
      } else {
        state.session.user.avatar = false;
        state.session.user.avatarUrl = null
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
    }
  },
  actions: {
    newLoginSession({ commit }, payload) {
      var expires
      if (payload.authToken) {
        expires = DaysToSec(30)
      } else {
        expires = MinToSec(15)
      }
      expires = UnixTime(expires)

      commit({
        type: 'newSession',
        session: {
          loggedIn: true,
          user: payload.user,
          fresh: true,
          created: UnixTime(),
          expires: expires
        }
      })
    },
    sessionCheck({ commit, state }) {
      if (state.session.loggedIn || state.session.fresh) {
        status().then(st => {
          if (st.loggedIn == false) {
            commit({
              type: 'newSession',
              session: sessionDefault()
            })
          } else {
            commit({
              type: 'newSession',
              session: {
                loggedIn: true,
                user: st.user,
                fresh: state.session.fresh,
                created: state.session.created,
                expires: state.session.expires
              }
            })
            commit('usedSession')
          }
        }).catch(e => {
          console.error(e);
        })
      }
    },
    sessionKeepAlive({ commit }) {
      keepAlive().then(() => {
        commit('usedSession')
      })
    },
    sessionLogout({ commit, state }) {
      if (state.session.loggedIn) {
        logout()
      }
      commit({
        type: 'newSession',
        session: sessionDefault()
      })
    }
  },
  getters: {
    loggedIn(state, getters) {
      if (state.session.expires < UnixTime()) {
        return false
      }
      return state.session.loggedIn
    },
    user(state) {
      return new GEN_User(state.session.user)
    },
    theme(state) {
      return state.ui.theme
    },
    viewKey(state) {
      return state.ui.viewKey
    }
  },
  plugins: [vuexLocal.plugin]
})