import { Media } from "./media";

export class Post {
  id: number;
  userId: number;
  created?: number;
  published: number;
  updated: number;
  content: string;
  media: boolean;
  mediaList?: Media[];

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.id = source["id"];
    this.userId = source["userId"];
    this.created = source["created"];
    this.published = source["published"];
    this.updated = source["updated"];
    this.content = source["content"];
    this.media = source["media"];
    if (source["mediaList"] != undefined && source["mediaList"].slice) {
      this.mediaList = (source["mediaList"] as any[]).map((m) => new Media(m));
    }
  }
}
