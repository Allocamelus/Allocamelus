import v1 from "../../v1";

export function keepAlive(): void {
  v1.post("account/auth/keep-alive");
}
