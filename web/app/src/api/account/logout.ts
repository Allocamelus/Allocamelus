import v1 from "../v1";

export function logout(): void {
  v1.delete("account/logout");
}
