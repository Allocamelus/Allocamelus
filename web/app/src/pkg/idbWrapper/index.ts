import { StorageLikeAsync } from "@vueuse/core";
import { DBSchema, IDBPDatabase, StoreNames, StoreValue } from "idb";
import { A9sDatabase } from "./allocamelus";

export interface IDBStore<
  DBType extends DBSchema | unknown,
  Key extends StoreValue<DBType, StoreNames<DBType>>,
  Value extends StoreValue<DBType, StoreNames<DBType>>
> extends StorageLikeAsync {}

export class IDBStore<
  DBType extends DBSchema | unknown,
  Key extends StoreValue<DBType, StoreNames<DBType>>,
  Value extends StoreValue<DBType, StoreNames<DBType>>
> {
  store: StoreNames<DBType>;
  db: Promise<IDBPDatabase<DBType>>;

  constructor(db: Promise<IDBPDatabase<DBType>>, store: StoreNames<DBType>) {
    this.store = store;
    this.db = db;
  }
  getItem(key: string) {
    return new Promise<string | null>(async (resolve) => {
      let v = await (await this.db).get(this.store, key as Key);
      if (typeof v == "undefined") {
        return resolve(null);
      }
      return resolve(String(v).toString());
    });
  }
  setItem(key: string, value: string) {
    return new Promise<void>(async (resolve) => {
      await (await this.db).put(this.store, value as Value, key as Key);
      return resolve();
    });
  }
  removeItem(key: string) {
    return new Promise<void>(async (resolve) => {
      await (await this.db).delete(this.store, key as Key);
      return resolve();
    });
  }
}
