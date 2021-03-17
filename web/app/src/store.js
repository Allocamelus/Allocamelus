import { createStore } from 'vuex'
import VuexPersistence from 'vuex-persist'
import { UnixTime, MinToSec, DaysToSec } from "./models/time";
const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
  reducer: (state) => {
    if (state.session.expires > UnixTime()) {
      if (state.session.fresh == true) {
        state.session.fresh = false
      }
      return {
        session: state.session
      }
    }
    return {}
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
    session: {
      loggedIn: false,
      userId: 0,
      fresh: true,
      created: UnixTime(),
      expires: UnixTime(MinToSec(10))
    }
  },
  mutations: {
    newSession(state, payload) {
      state.session = payload.session
    },
    usedSession(state) {
      state.session.expires = UnixTime(MinToSec(15))
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
    usedSession({ commit, state }) {
      if (state.session.expires < UnixTime(MinToSec(15))) {
        commit('usedSession')
      }
    },
  },
  getters: {
    loggedIn(state, getters) {
      if (state.session.expires < UnixTime()) {
        return false
      }
      return state.session.loggedIn
    },
  },
  plugins: [vuexLocal.plugin]
})