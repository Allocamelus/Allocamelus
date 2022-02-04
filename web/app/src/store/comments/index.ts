import { defineStore } from "pinia";
import { API_Comment } from "@/api/post/comment";
import { User } from "@/models/user";
import { API_Comments } from "@/api/post/comments/get";
import { GetFromPath, GetPath } from "./getters";

export interface Comments {
  postId: number;
  comments: API_Comments;
  pathCache: { [id: number]: number[] };
}

export interface State extends Comments {}

const pathErr = (path: number[]) =>
  new Error(`Error: Missing comment at path (${path})`);

export const useCommentStore = (postId: number | string) => {
  return defineStore(`post-${postId}-comments`, {
    state: (): State => {
      return {
        postId: Number(postId).valueOf(),
        comments: new API_Comments(),
        pathCache: {},
      };
    },
    getters: {
      user(state: State) {
        return (id: number): User => {
          return state.comments.user(id);
        };
      },
    },
    actions: {
      comment(id: number) {
        const path = GetPath(this, id);

        return GetFromPath(this.comments, path);
      },
      missingReplies(id: number): number {
        let c = this.comment(id);
        if (c === null) {
          return 0;
        }
        return c.missingReplies();
      },
      populate(cs: Partial<API_Comments>) {
        this.$reset();
        this.comments = new API_Comments(cs);
      },
      // addComment to parentId
      addComment(comment: API_Comment, isNew: boolean) {
        if (comment.parentId === 0) {
          this.comments.appendComment(comment);
          return;
        }

        const path = GetPath(this, comment.parentId);
        let parent = GetFromPath(this.comments, path);
        if (parent === null) {
          throw pathErr(path);
        }

        parent.appendChild(comment, isNew);

        if (isNew) {
          path.pop();

          // Add reply count to all parents
          while (path.length > 0) {
            parent = GetFromPath(this.comments, path);
            if (parent === null) {
              throw pathErr(path);
            } else {
              parent.replies++;
            }
            path.pop();
          }
        }
      },
      /**
       * updateComment
       *
       * @param {State} state Store state
       * @param {API_Comment} comment Only id, updated, and content are used here
       */
      updateComment(comment: API_Comment) {
        const c = this.comment(comment.id);
        if (c === null) {
          throw new Error(`Error: Missing comment by id (${comment.id})"`);
        }
        c.updated = comment.updated;
        c.content = comment.content;
      },
      // addUser to state
      // TODO user vuex store
      addUser(user: User) {
        this.comments.appendUser(user);
      },
      // remove comment
      remove(id: number) {
        let path = GetPath(this, id);

        // Get parent path and child key
        const key = path.pop();
        if (key == undefined) {
          return;
        }

        // Is comment top level
        if (path.length === 0) {
          this.comments.delComment(key);
          // remove comment path from cache
          delete this.pathCache[id];
          return;
        }

        let parent = GetFromPath(this.comments, path);
        if (parent === null) {
          throw pathErr(path);
        }

        // delete comment
        if (Object.hasOwnProperty.call(parent.children, key)) {
          parent.delChild(key);
          // remove comment path from cache
          delete this.pathCache[id];

          // remove current parent from path
          path.pop();

          // Remove reply count from all parents
          while (path.length > 0) {
            parent = GetFromPath(this.comments, path);
            if (parent === null) {
              throw pathErr(path);
            } else {
              parent.delDeep();
            }
            path.pop();
          }
        }
      },
    },
  })();
};
