import { useStorageAsync } from "@vueuse/core";
import { defineStore } from "pinia";
import { IDBStore } from "@/pkg/idbWrapper";
import { A9sDatabase } from "@/pkg/idbWrapper/allocamelus";

import { keepAlive } from "@/api/account/auth/keepAlive";
import { status } from "@/api/account/auth/status";
import { logout } from "@/api/account/logout";
import { User } from "@/models/user";
import { MinToSec, UnixTime } from "@/pkg/time";

const storeName = "session";
const idbStore = new IDBStore(A9sDatabase, storeName);

export interface Session {
  loggedIn: boolean;
  user: User;
  fresh: boolean;
  created: number;
  expires: number;
}

export class Session {
  loggedIn: boolean;
  user: User;
  fresh: boolean;
  created: number;
  expires: number;

  constructor(source: Partial<Session> = {}) {
    if (typeof source === "string") source = JSON.parse(source);
    this.loggedIn = source["loggedIn"] || false;
    this.user = new User(source["user"]);
    this.fresh = source["fresh"] || true;
    this.created = source["created"] || UnixTime();
    this.expires = source["expires"] || UnixTime(MinToSec(10));
  }
}
export interface State extends Session {}

export const useSessionStore = defineStore(storeName, {
  state: () => {
    let s = new Session();
    return {
      loggedIn: useStorageAsync("loggedIn", s.loggedIn, idbStore),
      user: useStorageAsync("user", s.user, idbStore, {
        serializer: {
          read: (v: any) => new User(JSON.parse(v)),
          write: (v: User) => JSON.stringify(v),
        },
      }),
      fresh: false,
      created: useStorageAsync("created", s.created, idbStore),
      expires: useStorageAsync("expires", s.expires, idbStore),
    };
  },
  actions: {
    used() {
      if (this.expires < UnixTime(MinToSec(15))) {
        this.$patch({ expires: UnixTime(MinToSec(15)) });
      }
    },
    // Get & Update session status
    async getStatus() {
      // Get auth status from server
      let s = await status();

      // State mismatched reset state
      if (!s.loggedIn && this.loggedIn) {
        return this.reset();
      }

      if (s.loggedIn) {
        // State mismatched new session
        if (!this.loggedIn) {
          this.$patch(
            new Session({
              loggedIn: true,
              user: s.user,
            })
          );
        } else {
          // Sync local user with server's
          this.$patch({
            loggedIn: true,
            user: new User(s.user),
          });
        }
        this.used();
      }
    },
    // Keep session alive w/ server
    async keepAlive() {
      await keepAlive();
      this.used();
    },
    async logout() {
      await logout();
      this.reset();
    },
    reset() {
      return this.$patch(new Session());
    },
  },
});
