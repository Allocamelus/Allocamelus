import { GetterTree } from "vuex";
import { API_Comment } from "../../../api/post/comment";
import { GEN_User } from "../../../models/go_structs_gen";
import { State } from "./state";

export type Getters = {
  comment(state: State): (id: number) => API_Comment | null
  user(state: State): (id: number) => GEN_User
}

export const getters = <GetterTree<State, any>>{
  comment(state: State) {
    return (id: number): API_Comment | null => {
      for (const k in state.comments.comments) {
        if (Object.prototype.hasOwnProperty.call(state.comments.comments, k)) {
          let c = getComment(state.comments.comments[k], id)
          if (c) {
            return API_Comment.createFrom(c)
          }
        }
      }
      return null
    }
  },
  user(state: State) {
    return (id: number): GEN_User => {
      return state.comments.user(id)
    }
  }
}

export function getComment(comment: API_Comment, id: number): API_Comment {
  if (comment.id === id) {
    return comment
  }

  for (const k in comment.children) {
    if (Object.hasOwnProperty.call(comment.children, k)) {
      let c = getComment(comment.children[k], id)
      if (c) {
        return c
      }
    }
  }
}