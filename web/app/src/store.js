import { createStore, mapActions, mapGetters } from 'vuex'
import VuexPersistence from 'vuex-persist'
import { UnixTime, MinToSec } from "./models/time";
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
  }

})

export default createStore({
  state: {
    session: {
      loggedIn: false,
      userId: 0,
      fresh: true,
      expires: UnixTime(MinToSec(10))
    }
  },
  mutations: {
    increment(state) {
      state.session.userId++
    }
  },
  actions: {
    increment(context) {
      context.commit('increment')
    }
  },
  getters: {
    loggedIn (state, getters) {
      return state.session.loggedIn
    }
  },
  methods: {
    ...mapActions([
      'increment' // map `this.increment()` to `this.$store.dispatch('increment')`
    ]),
    ...mapGetters({
      loggedIn: 'loggedIn'
    })
  },
  plugins: [vuexLocal.plugin]
})