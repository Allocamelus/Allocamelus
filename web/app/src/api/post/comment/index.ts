import { Ordered_API_Comments } from "../comments/get";

export class API_Comment {
  id: number;
  userId: number;
  postId: number;
  parentId: number;
  created: number;
  updated: number;
  content: string;
  replies: number;
  depth: number;
  children: Ordered_API_Comments;

  static createFrom(source: Partial<API_Comment> = {}) {
    return new API_Comment(source);
  }

  constructor(source: Partial<API_Comment> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.id = source["id"];
    this.userId = source["userId"];
    this.postId = source["postId"];
    this.parentId = source["parentId"];
    this.created = source["created"];
    this.updated = source["updated"];
    this.content = source["content"];
    this.replies = source["replies"];
    this.depth = source["depth"];
    this.children = source["children"];
  }

  // Methods
  child(key: string): API_Comment {
    // Convert child to API_Comment class if not
    if (!(this.children[key] instanceof API_Comment)) {
      this.children[key] = new API_Comment(this.children[key]);
    }
    return this.children[key];
  }

  delChild(key: string) {
    if (Object.prototype.hasOwnProperty.call(this.children, key)) {
      delete this.children[key];
      this.replies--;
    }
  }

  // Get childKey by commentId
  childKey(commentId: number): string {
    for (const key in this.children) {
      if (Object.prototype.hasOwnProperty.call(this.children, key)) {
        const c = this.children[key];
        if (c.id == commentId) {
          return key;
        }
      }
    }
  }

  missingReplies(): number {
    let num = 0;
    for (const key in this.children) {
      if (Object.prototype.hasOwnProperty.call(this.children, key)) {
        num += this.child(key).replies;
        num++;
      }
    }

    num = this.replies - num;
    return num;
  }

  numDirectChildren(): number {
    if (this.children == undefined || this.children === null) {
      return 0;
    }
    return Object.keys(this.children).length;
  }

  hasChildren(): boolean {
    if (this.numDirectChildren() != 0) {
      return true;
    }
    return false;
  }

  appendChild(c: API_Comment, newChild = false) {
    c.depth = this.depth + 1;
    this.children[this.numDirectChildren()] = c;
    if (newChild) {
      this.replies++;
    }
  }
}
