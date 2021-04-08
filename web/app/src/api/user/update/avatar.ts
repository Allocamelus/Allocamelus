import v1 from "../../v1";

export async function avatar(userName: string, file: File) {
  var formData = new FormData();
  formData.append("avatar", file)
  return v1.post(`/user/${userName}/update/avatar`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}