import { useStore } from "src/stores/store";
const store = useStore();

export function connect() {
    store.connection = new WebSocket("ws://localhost:8080/ws");
    store.connection.addEventListener('open', onConnect);
    store.connection.addEventListener('close', onDisconnect);
    store.connection.addEventListener('message', onMessage);
    store.connection.addEventListener('error', onError);
}

export function disconnect() {
    store.connection?.close();
}

export function sendMessage(message: string) {
    console.log('Sending message:', message);
    if (store.connection) {
        store.connection.send(message);
    }
    store.messages = [...store.messages, [store.messages.length +1, "Me", message]]

}

function onConnect() {
    store.host_status = 'Connected';
    store.messages = [];
}

function onDisconnect() {
    store.host_status = 'Not Connected';
}

function onMessage(event: MessageEvent) {
    const id = store.messages.length +1
    const msg : [number, string, string] = [id, "Server", event.data]
    store.messages = [...store.messages, msg]
}

function onError(error: Event) {
    console.error(error);
    store.host_status = 'Not Connected';
}