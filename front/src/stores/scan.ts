import { writable } from 'svelte/store';

const baseURL = "http://localhost:4040" 

function createScan() {
	const { subscribe, set, update } = writable(undefined);

	return {
		subscribe,
		getScanImage: (mode, image) => {
            set(undefined)
            scanIsLoading.set(true)
            return fetch(`${baseURL}/docker-images/${mode}/scan?image=${image}`)
            .then(res => res.json())
            .then(res => {
                set(res.scan)
                scanIsLoading.set(false)
            })
        },
        getScanPath: (path) => {
            set(undefined)
            scanIsLoading.set(true)
            return fetch(`${baseURL}/file-system/scan?path=${path}`)
            .then(res => res.json())
            .then(res => {
                set(res.scan)
                scanIsLoading.set(false)
            })
        },
        getScanConfig: (path) => {
            set(undefined)
            scanIsLoading.set(true)
            return fetch(`${baseURL}/config-files/scan?path=${path}`)
            .then(res => res.json())
            .then(res => {
                set(res.scan)
                scanIsLoading.set(false)
            })
        },
        getScanRemoteRepo: (uri) => {
            set(undefined)
            scanIsLoading.set(true)
            return fetch(`${baseURL}/repo-git/scan?uri=${uri}`)
            .then(res => res.json())
            .then(res => {
                set(res.scan)
                scanIsLoading.set(false)
            })
        }
	};
}

export const scan = createScan()
export const scanIsLoading = writable(false)