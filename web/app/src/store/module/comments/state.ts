import { API_Comments } from "../../../api/post/comments/get";

export interface State {
  comments: API_Comments;
  comPathCache: { [id: number]: string[] };
}

export const state: State | (() => State) = () => ({
  comments: new API_Comments(),
  comPathCache: {},
});
