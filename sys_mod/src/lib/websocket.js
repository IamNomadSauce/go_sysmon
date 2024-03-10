import { writable } from 'svelte/store'

export const gpuData = writable({
    temperature: 0,
    powerConsumption: 0,
    gpuClockState: 0,
    vramPercent: 0,
    gpuPercent: 0,
})

export function connectWebSocket() {
    const ws = new WebSocket('ws://localhost:8069/ws')

    ws.onmessage = (event) => {
        const data = JSON.parse(event.data)
        gpuData.set(data)
    }

    ws.onclose = () => {
        console.log("Websocket connection closed. Attempting to reconnect...")
        setTimeout(connectWebSocket, 1000)
    }

    ws.onerror = (error) => {
        console.error("Websocket Error", error)
        ws.close()
    }
}

