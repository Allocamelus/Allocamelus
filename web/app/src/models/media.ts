export class Meta {
  alt: string;
  width: number;
  height: number;

  constructor(source: Partial<Meta> = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.alt = source["alt"] || "";
    this.width = source["width"] || 0;
    this.height = source["height"] || 0;
  }
}

export class Media {
  fileType?: number;
  meta: Meta;
  url: string;

  constructor(source: Partial<Media> = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.fileType = source["fileType"];
    this.meta = new Meta(source["meta"]);
    this.url = source["url"] || "";
  }
}
