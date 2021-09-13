export class API_Comment {
  id: number;
  userId: number;
  postId: number;
  parentId: number;
  created: number;
  updated: number;
  content: string;
  replies: number;
  depth: number
  children: { [key: number]: API_Comment };

  static createFrom(source: any = {}) { // skipcq: JS-0323, JS-0306
    return new API_Comment(source);
  }

  constructor(source: any = {}) { // skipcq: JS-0323
    if ('string' === typeof source) source = JSON.parse(source);
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

  // Method
  child(commentId: number): API_Comment {
    return API_Comment.createFrom(this.children[commentId]);
  }

  numNotHad(): number {
    let num = 0
    for (const key in this.children) {
      if (Object.prototype.hasOwnProperty.call(this.children, key)) {
        const element = API_Comment.createFrom(this.children[key]);
        num += element.replies
        num++
      }
    }

    num = this.replies - num
    return num
  }

  numDirectChildren(): number {
    if (this.children == undefined || this.children === null) {
      return 0
    }
    return Object.keys(this.children).length
  }

  hasChildren(): boolean {
    if (this.numDirectChildren() != 0) {
      return true
    }
    return false
  }

  appendChild(c: API_Comment) {
    c.depth = this.depth + 1
    this.children[this.numDirectChildren()] = c
  }
}