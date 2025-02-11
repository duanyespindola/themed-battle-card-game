import { defineStore, acceptHMRUpdate } from 'pinia';
type msg = [number, string, string];
export const useStore = defineStore('counter', {
  state: () => ({
    host_status : 'Not Connected',
    messages : [] as msg[],
    connection: null as WebSocket | null,
  }),

});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useStore, import.meta.hot));
}
