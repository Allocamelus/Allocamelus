import { state } from "./state";
import { getters } from "./getters";
import { mutations } from "./mutations";

const CommentsStore = {
  namespaced: true,
  state,
  getters,
  mutations,
};

export default CommentsStore;
