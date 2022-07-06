import { writable } from 'svelte/store';

const PATH_STORAGE_KEY = "path"
const REPO_STORAGE_KEY = "repo"

function createPath(key) {

    const path = localStorage.getItem(key)
	const { subscribe, set, update } = writable(path ? path : "");

	return {
		subscribe,
		savePath: (path) => {
            set(path)
            localStorage.setItem(key, path)
        }
	};
}

export const sessionPath = createPath(PATH_STORAGE_KEY)
export const sessionRepo = createPath(REPO_STORAGE_KEY)