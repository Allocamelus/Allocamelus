import { GetterTree } from "vuex";
import { User } from "../models/user";
import { UnixTime } from "../pkg/time";
import { State } from "./index";

export type Getters = {
  loggedIn(state: State): boolean;
  user(state: State): User;
  theme(state: State): string;
  viewKey(state: State): number;
};

// skipcq: JS-0323
export const getters = <GetterTree<State, any>>{
  loggedIn(state: State) {
    if (state.session.expires < UnixTime()) {
      return false;
    }
    return state.session.loggedIn;
  },
  user(state: State) {
    return new User(state.session.user);
  },
  theme(state: State) {
    return state.ui.theme;
  },
  viewKey(state: State) {
    return state.ui.viewKey;
  },
};
