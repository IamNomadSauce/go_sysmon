<script>
    import { onMount } from 'svelte'
    import { gpuData, connectWebSocket } from "$lib/websocket"
    import LineChart from '$lib/LineChart.svelte';

    onMount(() => {
        connectWebSocket()
    })

    let temp = {
        currentGPUClockState: 0,
        temperature: 0,
        gpuUtilization: 0,
        powerConsumption: 0,
        vramUsedPercentage: 0,
    }

    $: data = ($gpuData.length > 0) ? $gpuData : [temp]
    // $: console.log("LAST\n", $gpuData.length,data[data.length-1])
    // $: console.log("GPU Data", $gpuData.length, $gpuData)
</script>

<!-- <p>GPU Temperature: {$gpuData.temperature} °C</p>
<p>GPU Power Consumption: {$gpuData.powerConsumption} W</p>
<p>Current GPU Clock State: {$gpuData.currentGPUClockState}</p>
<p>GPU VRAM Usage: {$gpuData.vramUsedPercentage}%</p>
<p>GPU Utilization: {$gpuData.gpuUtilization}%</p> -->


    <h5>Temperature:{data[data.length-1].temperature}C {1 * (data[data.length-1].temperature) + 32}F</h5>
    <h5>GPU-Utilization: {data[data.length-1].gpuUtilization}%</h5>
    <h5>GPU-Clock: {data[data.length-1].currentGPUClockState}</h5>
    <h5>Avg-Power: {data[data.length-1].powerConsumption}W</h5>
    <h5>VRam%: {data[data.length-1].vramUsedPercentage}</h5>


<LineChart data={data} width="600" height="300" />

<style>
    :global(body) {
        background-color: black;
        color: white; /* Optional: Change text color to white for better readability */
    }
</style>