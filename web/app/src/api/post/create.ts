import v1 from "../v1";
import { API_Success_Error } from "../../models/api_error";

export class CreateResponse extends API_Success_Error {
  id?: number;

  static createFrom(source: object | string = {}) {
    return new CreateResponse(source);
  }

  constructor(source: object | string = {}) {
    super(source);
    if (typeof source === "string") source = JSON.parse(source);
    this.id = source["id"];
  }
}

export class MediaFile {
  media: File;
  alt: string;

  static createFrom(source: object | string = {}) {
    return new MediaFile(source);
  }

  constructor(source: object | string = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.media = source["media"];
    this.alt = source["alt"];
  }
}

export async function create(
  content: string,
  images: Array<MediaFile>,
  publish: boolean
): Promise<CreateResponse | API_Success_Error> {
  const formData = new FormData();
  formData.append("publish", JSON.stringify(publish));
  formData.append("content", content);
  for (let i = 0; i < images.length; i++) {
    formData.append("images[]", images[i].media);
    formData.append("imageAlts[]", images[i].alt.toString());
  }

  return v1
    .post("/post", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    })
    .then((r) => {
      return CreateResponse.createFrom(r.data);
    })
    .catch((e) => {
      return API_Success_Error.createFrom(e);
    });
}
