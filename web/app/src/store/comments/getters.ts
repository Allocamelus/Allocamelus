import { API_Comment } from "@/api/post/comment";
import { API_Comments } from "@/api/post/comments/get";
import { State } from ".";

export function GetFromPath(
  comments: API_Comments | API_Comment,
  path: number[]
): API_Comment | null {
  let c: API_Comment | null = null;

  path.forEach((key) => {
    if (c === null) {
      if (comments instanceof API_Comments) {
        c = comments.comment(comments.comments[key].id);
      } else {
        c = comments.child(key);
      }
    } else {
      c = c.child(key);
    }
  });

  return c;
}
/**
 * getDeepPath - Deep search of comment's children to find Comment ID
 *
 * @param comment
 * @param id Comment ID
 * @param path Current path
 * @returns Comment Path
 */
function getDeepPath(
  comment: API_Comment,
  id: number,
  path: number[]
): number[] | null {
  // Return path when comment is found
  if (comment.id == id) {
    return path;
  }
  let keys: number[] = [];
  // Iterate over comment's children
  for (const k in comment.children) {
    if (Object.hasOwnProperty.call(comment.children, k)) {
      keys = [];
      keys.push(...path, Number(k).valueOf());
      // Recursively search for comment
      const p = getDeepPath(comment.children[k], id, keys);
      if (p !== null) {
        return p;
      }
    }
  }
  return null;
}

/**
 * getPathFromComments - Iterate over top level of API_Comments for deep search
 *
 * Deep search is done by getDeepPath
 *
 * @param comments List of Comments
 * @param id Comment ID
 * @returns Path to Comment
 */
function getPathFromComments(comments: API_Comments, id: number): number[] {
  let path: number[] | null = [];

  // Iterate over top comments
  for (const k in comments.comments) {
    const key = Number(k).valueOf();
    // Check that key exists
    if (Object.prototype.hasOwnProperty.call(comments.comments, key)) {
      // Deep search
      path = getDeepPath(comments.comments[k], id, [key]);
      if (path !== null) {
        return path;
      }
    }
  }
  return [];
}

export function GetPath(state: State, id: number) {
  // If pathCache Does Not have path to id || pathCache[id] has no path
  if (
    !Object.prototype.hasOwnProperty.call(state.pathCache, id) ||
    Object.keys(state.pathCache[id]).length === 0
  ) {
    state.pathCache[id] = getPathFromComments(state.comments, id);
  }
  return state.pathCache[id];
}
