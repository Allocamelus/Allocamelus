import v1 from "../v1";
import { API_Error, API_Success_Error } from "../../models/api_error";


export class CreateResponse extends API_Success_Error {
  id?: Number;

  static createFrom(source: any = {}) {
    return new CreateResponse(source);
  }

  constructor(source: any = {}) {
    super(source)
    if ('string' === typeof source) source = JSON.parse(source);
    this.id = source["id"];
  }
}

export class MediaFile {
  media: File;
  alt: String;

  static createFrom(source: any = {}) {
    return new MediaFile(source);
  }

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.media = source["media"];
    this.alt = source["alt"];
  }
}

export async function create(content: string, images: Array<MediaFile>, publish: boolean) {
  var formData = new FormData();
  formData.append("publish", JSON.stringify(publish))
  formData.append("content", content)
  var image;
  for (let i = 0; i < images.length; i++) {
    formData.append("images[]", images[i].media)
    formData.append("imageAlts[]", images[i].alt.toString())
  }

  return v1.post("/post", formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
    .then(r => {
      return CreateResponse.createFrom(r.data)
    })
    .catch(e => {
      return API_Success_Error.createFrom(e)
    })
}