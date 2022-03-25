import { DBSchema, IDBPDatabase, openDB } from "idb";

interface A9sDBV1 extends DBSchema {
  keyValueStore: { key: string; value: string };
}
interface A9sDBV2 extends DBSchema {
  state: { key: string; value: string };
  session: { key: string; value: string };
}

export const A9sDatabase = (async () => {
  let dbName = "allocamelus";
  let version = 2;
  const db = await openDB<A9sDBV2>(dbName, version, {
    async upgrade(db, oldVersion) {
      const v1db = db as unknown as IDBPDatabase<A9sDBV1>;
      console.log(oldVersion, db.objectStoreNames);

      if (oldVersion < 1) {
        v1db.createObjectStore("keyValueStore");
      }
      console.log(db.objectStoreNames);

      if (oldVersion < 2) {
        v1db.deleteObjectStore("keyValueStore");

        db.createObjectStore("state");
        db.createObjectStore("session");

        console.log(db.objectStoreNames);
      }
    },
  });
  return db;
})();
