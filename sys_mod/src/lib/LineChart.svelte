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

<ScatterPlot data={temp_data} title="Temperature" color="orange" />
<ScatterPlot data={gpuU_data} title="GPU Utilization %" color="blue"/>
<ScatterPlot data={pwr_data} title="Avg Power (W)" color="yellow" />
<ScatterPlot data={vram_data} title="VRam Utilization" color="white"/>
<ScatterPlot data={mclk_data} title="Memory Clock (MHz)" color="purple"/>
<ScatterPlot data={sclk_data} title="System Clock (MHz)" color="green"/>





<!-- Power Consumption Line Chart -->
<!-- <h5>Power</h5>
<svg class="historgram" width={chartWidth} height={chartHeight}>
    <g
        bind:this={xAxisPower}
        transform={`translate(0, ${chartHeight - paddingBottom})`}
        class="x-Axis"
    ></g>
    <g bind:this={yAxisPower} class="y-Axis"></g>

    Draw the line for power consumption
    <path
        d={lineGeneratorPower(sliced_data)}
        fill="none"
        stroke="orange"
        stroke-width="1"
    />
</svg> -->


<!-- Temperature Line Chart
<h5>Temperature</h5>
<svg class="historgram" width={chartWidth} height={chartHeight} style="margin-top: 20px;">
    <g
        bind:this={xAxisTemp}
        transform={`translate(0, ${chartHeight - paddingBottom})`}
        class="x-Axis"
    ></g>
    <g bind:this={yAxisTemp} class="y-Axis"></g>

    Draw the line for temperature
    <path
        d={lineGeneratorTemp(sliced_data)}
        fill="none"
        stroke="cyan"
        stroke-width="1"
    />
</svg> -->