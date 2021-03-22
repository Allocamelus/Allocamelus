import { createStore } from 'vuex'
import VuexPersistence from 'vuex-persist'
import { UnixTime, MinToSec, DaysToSec } from "./pkg/time";

import { status } from "./api/account/auth/status"
import { keepAlive } from "./api/account/auth/keepAlive"

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
    fresh: true,
    created: UnixTime(),
    expires: UnixTime(MinToSec(10))
  }
}

export default createStore({
  state: {
    theme: 'dark',
    session: {
      loggedIn: false,
      userId: 0,
      fresh: true,
      created: UnixTime(),
      expires: UnixTime(MinToSec(10))
    },
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
          fresh: true,
          created: UnixTime(),
          expires: expires
        }
      })
    },
    sessionCheck({ commit, state }) {
      if (state.session.loggedIn || state.session.fresh) {
        status().then(session => {
          if (state.session.loggedIn != session.loggedIn || state.session.userId != session.userId) {
            commit({
              type: 'newSession',
              session: {
                loggedIn: session.loggedIn,
                userId: session.userId,
                fresh: true,
                created: UnixTime(),
                expires: MinToSec(15)
              }
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
    }
  },
  getters: {
    loggedIn(state, getters) {
      if (state.session.expires < UnixTime()) {
        return false
      }
      return state.session.loggedIn
    },
    theme(state) {
      return state.theme
    }
  },
  plugins: [vuexLocal.plugin]
})