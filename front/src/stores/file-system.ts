import { writable } from 'svelte/store';

const baseURL = "http://localhost:4040" 

function createFileSystem() {
	const { subscribe, set, update } = writable(undefined);

	return {
		subscribe,
		check: (path) => {
			set(undefined)
			fileSystemIsLoading.set(true)
			return fetch(`${baseURL}/file-system/check?path=${path}`)
            .then(res => res.json())
            .then(res => {
				set(res)
				fileSystemIsLoading.set(false)
			})
        }
	};
}

export const fileSystem = createFileSystem();
export const fileSystemIsLoading = writable(false)