<script>
    import { onMount } from 'svelte'
    import { gpuData, connectWebSocket } from "$lib/websocket"

    onMount(() => {
        connectWebSocket()
    })
</script>

<p>GPU Temperature: {$gpuData.temperature} °C</p>
<p>GPU Power Consumption: {$gpuData.powerConsumption} W</p>
<p>Current GPU Clock State: {$gpuData.currentGPUClockState}</p>
<p>GPU VRAM Usage: {$gpuData.vramUsedPercentage}%</p>
<p>GPU Utilization: {$gpuData.gpuUtilization}%</p>

<style>
    :global(body) {
        background-color: black;
        color: white; /* Optional: Change text color to white for better readability */
    }
</style>