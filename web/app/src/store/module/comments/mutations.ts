import { MutationTree } from "vuex";
import { API_Comment } from "../../../api/post/comment";
import { API_Comments } from "../../../api/post/comments/get";
import { State } from "./state";

export type Mutations = {
  update(state: State, c: API_Comments): void
  remove(state: State, id: number): void
}

export const mutations = <MutationTree<State>>{
  update(state, c: API_Comments) {
    state.comments = c
  },
  remove(state: State, id: number) {
    for (const k in state.comments.comments) {
      if (Object.prototype.hasOwnProperty.call(state.comments.comments, k)) {
        // Check if comment matches removal id
        // Delete if so 
        console.log(k, id);
        
        if (state.comments.comments[k].id === id) {
          delete state.comments.comments[k]
          return
        }
        // If removeComment returns true return 
        if (removeComment(state.comments.comments[k], id)) {
          return
        }
      }
    }
  },
}

export function removeComment(comment: API_Comment, id: number): boolean {
  for (const k in comment.children) {
    if (Object.hasOwnProperty.call(comment.children, k)) {
      // Check if child matches removal id
      // Delete if so 
      if (comment.children[k].id === id) {
        delete comment.children[k]
        comment.replies--
        return true
      }
      // Recursively call until true
      if (removeComment(comment.children[k], id)) {
        return true
      }
    }
  }
  return false
}