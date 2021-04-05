import { createStore } from 'vuex'
import VuexPersistence from 'vuex-persist'
import { UnixTime, MinToSec, DaysToSec } from "./pkg/time";

import { status } from "./api/account/auth/status"
import { keepAlive } from "./api/account/auth/keepAlive"
import { logout } from "./api/account/logout"

const vuexLocal = new VuexPersistence({
  key: "a10storage",
  storage: window.localStorage,
  reducer: (state) => {
    var storage = {
      theme: state.theme
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
    userId: 0,
    userName: '',
    fresh: true,
    created: UnixTime(),
    expires: UnixTime(MinToSec(10))
  }
}

export default createStore({
  state: {
    theme: 'dark',
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
      state.theme = (state.theme == 'dark') ? 'light' : 'dark'
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
          userId: payload.userId,
          userName: payload.userName,
          fresh: true,
          created: UnixTime(),
          expires: expires
        }
      })
    },
    sessionCheck({ commit, state }) {
      if (state.session.loggedIn || state.session.fresh) {
        status().then(loggedIn => {
          if (loggedIn == false) {
            commit({
              type: 'newSession',
              session: sessionDefault()
            })
          } else {
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
    userName(state) {
      return state.session.userName
    },
    theme(state) {
      return state.theme
    }
  },
  plugins: [vuexLocal.plugin]
})