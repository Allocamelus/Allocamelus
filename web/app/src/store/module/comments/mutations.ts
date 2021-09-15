import { MutationTree } from "vuex";
import { API_Comment } from "../../../api/post/comment";
import { API_Comments, Ordered_API_Comments } from "../../../api/post/comments/get";
import { GEN_User } from "../../../models/go_structs_gen";
import { Comment, CommentFromPath, CommentPath } from "./getters";
import { State } from "./state";

export type AddChildrenParams = {
  parentId: number
  children: Ordered_API_Comments
}

export type Mutations = {
  populate(state: State, c: API_Comments): void
  addComment(state: State, comment: API_Comment): void
  updateComment(state: State, comment: API_Comment): void
  addUser(state: State, user: GEN_User): void
  remove(state: State, id: number): void
}

export const mutations = <MutationTree<State>>{
  // populate state comments
  populate(state, c: API_Comments) {
    state.comments = c
  },
  // addComment to parentId
  addComment(state: State, comment: API_Comment) {
    if (comment.parentId === 0) {
      state.comments.appendComment(comment)
      return
    }
    let path = CommentPath(state)(comment.parentId)
    let parent = CommentFromPath(state.comments, path)
    parent.appendChild(comment)
    path.pop()

    // Add reply count to all parents
    while (path.length > 0) {
      parent = CommentFromPath(state.comments, path)
      parent.replies++
      path.pop()
    }
  },
  /**
   * updateComment
   *
   * @param {State} state Store state
   * @param {API_Comment} comment Only id, updated, and content are used here
   */
  updateComment(state: State, comment: API_Comment) {
    let c = Comment(state)(comment.id)
    c.updated = comment.updated
    c.content = comment.content
  },
  // addUser to state
  // TODO user vuex store
  addUser(state: State, user: GEN_User) {
    state.comments.appendUser(user)
  },
  // remove comment
  remove(state: State, id: number) {
    let path = CommentPath(state)(id)
    // Get parent path and child key
    let key = path.pop()
    let parent = CommentFromPath(state.comments, path)

    // delete comment
    if (Object.hasOwnProperty.call(parent.children, key)) {
      delete parent.children[key]
    }

    // remove comment path from cache
    delete state.comPathCache[id]

    // Remove reply count from all parents
    while (path.length > 0) {
      parent = CommentFromPath(state.comments, path)
      parent.replies--
      path.pop()
    }
  },
}