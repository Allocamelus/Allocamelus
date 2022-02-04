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
    this.id = source["id"] || 0;
    this.userId = source["userId"] || 0;
    this.postId = source["postId"] || 0;
    this.parentId = source["parentId"] || 0;
    this.created = source["created"] || 0;
    this.updated = source["updated"] || 0;
    this.content = source["content"] || "";
    this.replies = source["replies"] || 0;
    this.depth = source["depth"] || 0;
    this.children = source["children"] || [];
  }

  // Methods
  child(key: number): API_Comment | null {
    if (!Object.hasOwnProperty.call(this.children, key)) {
      return null;
    }
    // Convert child to API_Comment class if not
    if (!(this.children[key] instanceof API_Comment)) {
      this.children[key] = new API_Comment(this.children[key]);
    }
    return this.children[key];
  }

  delChild(key: number): void {
    if (Object.prototype.hasOwnProperty.call(this.children, key)) {
      delete this.children[key];
      this.delDeep();
    }
  }

  delDeep() {
    if (this.replies > 0) {
      this.replies--;
    }
  }

  missingReplies(): number {
    let num = 0;
    for (const k in this.children) {
      const key = Number(k).valueOf();
      if (Object.prototype.hasOwnProperty.call(this.children, key)) {
        num += this.child(key)!.replies;
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

  appendChild(c: API_Comment, newChild = false): void {
    c.depth = this.depth + 1;
    let key = 0;
    if (this.numDirectChildren() != 0) {
      key = Number(Object.keys(this.children).slice(-1)[0]).valueOf() + 1;
    }
    this.children[key] = c;
    if (newChild) {
      this.replies++;
    }
  }
}
