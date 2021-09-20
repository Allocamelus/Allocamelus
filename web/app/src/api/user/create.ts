import v1 from "../v1";
import {
  GEN_CreateA10Token,
  GEN_CreateRequest,
  GEN_CreateResp,
} from "../../models/go_structs_gen";

const a9s = "allocamelus";

export async function create(
  request: GEN_CreateRequest
): Promise<GEN_CreateResp> {
  return v1
    .post("user", JSON.stringify(request), {
      headers: {
        "Content-Type": "application/json",
      },
    })
    .then((r) => {
      return GEN_CreateResp.createFrom(r.data);
    });
}

export async function createA9s(
  token: GEN_CreateA10Token
): Promise<GEN_CreateResp> {
  return create(
    GEN_CreateRequest.createFrom({
      with: a9s,
      token: JSON.stringify(token),
    })
  );
}
