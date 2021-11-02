import v1 from "../v1";

export function logout() {
  v1.delete("account/logout");
}
