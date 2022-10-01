import { writable } from "svelte/store";

const MAX_MEMORY = 50;
const LOCAL_STORAGE_KEY = "reports-history";

function createHistory() {
  const storage = JSON.parse(localStorage.getItem(LOCAL_STORAGE_KEY));
  const { subscribe, set, update } = writable(storage ? storage : []);

  return {
    subscribe,
    saveScan: (name, type) => {
      update((storage) => {
        storage.unshift({
          name,
          type,
          date: new Date().toDateString(),
        });
        if (storage.length > MAX_MEMORY) {
          storage = storage.slice(0, MAX_MEMORY);
        }
        localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(storage));
        return storage;
      });
    },
  };
}

export const history = createHistory();
