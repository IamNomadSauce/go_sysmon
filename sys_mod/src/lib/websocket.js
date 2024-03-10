import { writable } from 'svelte/store'

export const gpuData = writable([]);

export function connectWebSocket() {
  const ws = new WebSocket('ws://localhost:8069/ws');

  ws.onmessage = (event) => {
    const newData = JSON.parse(event.data);
    newData.time = Date.now()
    gpuData.update(currentData => {
        console.log("GPU DATA", newData)
      return [...currentData, newData];
    });
  };

    ws.onclose = () => {
        console.log("Websocket connection closed. Attempting to reconnect...")
        setTimeout(connectWebSocket, 1000)
    }

    ws.onerror = (error) => {
        console.error("Websocket Error", error)
        ws.close()
    }
}

