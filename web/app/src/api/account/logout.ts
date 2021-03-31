import v1 from "../v1";

export async function logout() {
  return v1.delete("account/logout")
}