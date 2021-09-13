import { MutationTree } from "vuex";
import { API_Comment } from "../../../api/post/comment";
import { API_Comments } from "../../../api/post/comments/get";
import { State } from "./state";

export type Mutations = {
  update(state: State, c: API_Comments): void
  updateComment(state: State, comment: API_Comment): void
  remove(state: State, id: number): void
}

export const mutations = <MutationTree<State>>{
  // update state comments
  update(state, c: API_Comments) {
    state.comments = c
  },
  /**
   * updateComment
   *
   * @param {State} state Store state
   * @param {API_Comment} comment Only id, updated, and content are used here
   */
  updateComment(state: State, comment: API_Comment) {
    for (const k in state.comments.comments) {
      if (Object.prototype.hasOwnProperty.call(state.comments.comments, k)) {
        // Check if comment matches removal id
        // Delete if so 
        if (state.comments.comments[k].id === comment.id) {
          state.comments.comments[k].updated = comment.updated
          state.comments.comments[k].content = comment.content
          return
        }
        // If removeComment returns true return 
        if (upComment(state.comments.comments[k], comment)) {
          return
        }
      }
    }
  },
  // remove comment
  remove(state: State, id: number) {
    for (const k in state.comments.comments) {
      if (Object.prototype.hasOwnProperty.call(state.comments.comments, k)) {
        // Check if comment matches removal id
        // Delete if so 
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

function upComment(comment: API_Comment, newComment: API_Comment): boolean {
  for (const k in comment.children) {
    if (Object.hasOwnProperty.call(comment.children, k)) {
      // Check if child matches removal id
      // Update if so 
      if (comment.children[k].id === newComment.id) {
        comment.children[k].updated = comment.updated
        comment.children[k].content = comment.content
        return true
      }
      // Recursively call until true
      if (upComment(comment.children[k], newComment)) {
        return true
      }
    }
    return false
  }
}

function removeComment(comment: API_Comment, id: number): boolean {
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
