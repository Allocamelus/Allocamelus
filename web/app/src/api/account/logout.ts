import v1 from "../v1";

export async function logout() {
  v1.delete("account/logout")
}