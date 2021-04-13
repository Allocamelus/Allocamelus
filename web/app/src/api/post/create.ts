import v1 from "../v1";
import { API_Success_Error } from "../../models/api_error";


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

export async function create(content: string, publish: boolean) {
  return v1.post("/post",
    JSON.stringify({
      publish: publish,
      content: content
    }), {
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(r => {
      return CreateResponse.createFrom(r.data)
    })
}