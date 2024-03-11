<script>
    import { onMount, afterUpdate } from "svelte";
    
    import ScatterPlot from "./ScatterPlot.svelte";

    export let data = [];
    export let width = 600;
    export let height = 300;

    let slice_window = 1440;
    let chartWidth = 800;
    let chartHeight = 300;
    let paddingLeft = 50;
    let paddingRight = 25;
    let paddingTop = 50;
    let paddingBottom = 25;
    // $: sliced_data = data.slice(-slice_window);

    $: console.log("data", data);

    $: temp_data = data.map((v) => ({time: v.time, value: v.temperature}))
    $: gpuU_data = data.map((v) => ({time: v.time, value: v.gpuUtilization}))
    $: pwr_data = data.map((v) => ({time: v.time, value: v.powerConsumption}))
    $: vram_data = data.map((v) => ({time: v.time, value: v.vramUsedPercentage}))
    $: mclk_data = data.map((v) => ({time: v.time, value: v.mclk}))
    $: sclk_data = data.map((v) => ({time: v.time, value: v.sclk}))




</script>

<div class="d-flex row">
    <div class="col">
        <ScatterPlot data={temp_data} title="Temperature" color="orange" />
    </div>
    <div class="col">
        <ScatterPlot data={gpuU_data} title="GPU Utilization %" color="blue"/>
    </div>
</div>

<div class="row">
    <div class="col">
        <ScatterPlot data={pwr_data} title="Avg Power (W)" color="yellow" />
    </div>
    <div class="col">
        <ScatterPlot data={vram_data} title="VRam Utilization" color="white"/>
    </div>
</div>

<div class="row">

    <div class="col">
        <ScatterPlot data={mclk_data} title="Memory Clock (MHz)" color="purple"/>
    </div>
    <div class="col">
        <ScatterPlot data={sclk_data} title="System Clock (MHz)" color="green"/>
    </div>
</div>