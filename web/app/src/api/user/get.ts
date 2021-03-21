import v1 from "../v1";
import { User } from '../../models/user_gen'
// TODO Cache
export async function get(uniqueName: string) {
  return v1.get("/user/" + uniqueName)
    .then(r => {
      if (r.data.error == undefined) {
        return User.createFrom(r.data)
      }
    })
}