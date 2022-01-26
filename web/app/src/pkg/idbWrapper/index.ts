import { StorageLikeAsync } from "@vueuse/core";
import { get, set, del, createStore, UseStore } from "idb-keyval";

export interface IDBStore extends StorageLikeAsync {}

const store = createStore("allocamelus", "keyValueStore");

export class IDBStore {
  prefix: string;
  store: UseStore;

  constructor(prefix: string) {
    this.prefix = prefix;
    this.store = store;
  }
  getItem(key: string) {
    return new Promise<string | null>(async (resolve) => {
      let v = await get(this.preKey(key), this.store);
      if (typeof v == "undefined") {
        resolve(null);
        return;
      }
      resolve(String(v));
      return;
    });
  }
  setItem(key: string, value: string) {
    return set(this.preKey(key), value, this.store);
  }
  removeItem(key: string) {
    return del(this.preKey(key), this.store);
  }
  preKey(key: string) {
    return `${this.prefix}-${key}`;
  }
}

export function newStore(prefix: string): IDBStore {
  return new IDBStore(prefix);
}
