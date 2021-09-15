import { GetterTree } from "vuex";
import { API_Comment } from "../../../api/post/comment";
import { API_Comments } from "../../../api/post/comments/get";
import { GEN_User } from "../../../models/go_structs_gen";
import { State } from "./state";

export type Getters = {
  comment(state: State): (id: number) => API_Comment | null
  user(state: State): (id: number) => GEN_User
}

export type GettersResp = {
  comment(id: number): API_Comment | null
  user(id: number): GEN_User
}

export const getters = <GetterTree<State, any>>{
  comment(state: State) {
    return (id: number): API_Comment | null => {
      let path = CommentPath(state)(id)
      return CommentFromPath(state.comments, path)
    }
  },
  user(state: State) {
    return (id: number): GEN_User => {
      return state.comments.user(id)
    }
  }
}

export function Comment(state: State) {
  return (id: number): API_Comment | null => {
    let path = CommentPath(state)(id)
    return CommentFromPath(state.comments, path)
  }
}

export function CommentPath(state: State) {
  return (id: number): string[] => {
    if (!Object.prototype.hasOwnProperty.call(state.comPathCache, id) ||
      Object.keys(state.comPathCache[id]).length === 0
    ) {
      state.comPathCache[id] = commentPathFromComments(state.comments, id)
    }
    return state.comPathCache[id]
  }
}

function commentPathFromComments(comments: API_Comments, id: number): string[] {
  let path: string[] | null
  for (const k in comments.comments) {
    if (Object.prototype.hasOwnProperty.call(comments.comments, k)) {
      path = getCommentPath(comments.comments[k], id, [k])
      if (path !== null) {
        return path
      }
    }
  }
}

export function CommentFromPath(comments: API_Comments | API_Comment, path: string[]): API_Comment {
  let c: API_Comment = null

  path.forEach(key => {
    if (c === null) {
      if (comments instanceof API_Comments) {
        c = comments.comment(comments.comments[key].id)
      } else {
        c = comments.child(key)
      }
    } else {
      c = c.child(key)
    }
  });

  return c
}


function getCommentPath(comment: API_Comment, id: number, path: string[]): string[] | null {
  if (comment.id == id) {
    return path
  }
  let keys: string[]
  for (const k in comment.children) {
    if (Object.hasOwnProperty.call(comment.children, k)) {
      keys = []
      keys.push(...path, k)
      let p = getCommentPath(comment.children[k], id, keys)
      if (p !== null) {
        return p
      }
    }
  }
  return null
}