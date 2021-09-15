import { MutationTree } from "vuex";
import { API_Comment } from "../../../api/post/comment";
import { API_Comments, Ordered_API_Comments } from "../../../api/post/comments/get";
import { Comment, CommentFromPath, CommentPath } from "./getters";
import { State } from "./state";

export type AddChildrenParams = {
  parentId: number
  children: Ordered_API_Comments
}

export type Mutations = {
  populate(state: State, c: API_Comments): void
  updateComment(state: State, comment: API_Comment): void
  remove(state: State, id: number): void
}

export const mutations = <MutationTree<State>>{
  // populate state comments
  populate(state, c: API_Comments) {
    state.comments = c
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
  // remove comment
  remove(state: State, id: number) {
    let path = CommentPath(state)(id)
    // Get parent path and child key
    let key = path.pop()
    let parent = CommentFromPath(state.comments, path)
    // delete comment
    if (Object.hasOwnProperty.call(parent.children, key)) {
      console.log(parent.replies);
      
      delete parent.children[key]
      parent.replies--
      console.log(parent.replies);
    }
    // remove comment path from cache
    delete state.comPathCache[id]
  },
}