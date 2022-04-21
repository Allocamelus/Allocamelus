import { StorageLikeAsync } from "@vueuse/core";
import { DBSchema, IDBPDatabase, StoreNames, StoreValue } from "idb";

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
  async getItem(key: string) {
    const v = await (await this.db).get(this.store, key as Key);
    if (typeof v == "undefined") {
      return null;
    }
    return String(v).toString();
  }
  async setItem(key: string, value: string) {
    await (await this.db).put(this.store, value as Value, key as Key);
    return;
  }
  async removeItem(key: string) {
    await (await this.db).delete(this.store, key as Key);
    return;
  }
}
