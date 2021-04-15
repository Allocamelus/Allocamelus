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

export async function create(content: string, images: Array<File>, publish: boolean) {
  var formData = new FormData();
  formData.append("publish", JSON.stringify(publish))
  formData.append("content", content)
  var image;
  for (let i = 0; i < images.length; i++) {
    formData.append("images[]", images[i])
  }

  return v1.post("/post", formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
    .then(r => {
      return CreateResponse.createFrom(r.data)
    })
}

export async function CreatePost(content: string, images: Array<File>) {
  return create(content, images, true).then(r => {
    if (!r.success) {
      throw new API_Error(r);
    }
  })

}