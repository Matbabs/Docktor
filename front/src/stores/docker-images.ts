import { writable } from "svelte/store";

const baseURL = "http://localhost:4040";

function createLocalDockerImages() {
  const { subscribe, set, update } = writable([]);

  return {
    subscribe,
    get: () => {
      set([]);
      localDockerImagesIsLoading.set(true);
      return fetch(`${baseURL}/docker-images/local`)
        .then((res) => res.json())
        .then((res) => {
          set(res.dockerImages);
          localDockerImagesIsLoading.set(false);
        });
    },
  };
}

export const localDockerImages = createLocalDockerImages();
export const localDockerImagesIsLoading = writable(false);
