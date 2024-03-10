<script>
    import { onMount, afterUpdate } from "svelte";
    import * as d3 from "d3";
    import { scaleTime, scaleLinear } from "d3-scale";
    import { extent, min, max } from "d3-array";
    import { axisBottom, axisRight } from "d3-axis";
    import { select } from "d3-selection";

    export let data = [];
    export let width = 600;
    export let height = 300;

    let slice_window = 1440;
    let chartWidth = 1500;
    let chartHeight = 300;
    let paddingLeft = 50;
    let paddingRight = 25;
    let paddingTop = 50;
    let paddingBottom = 25;
    $: sliced_data = data.slice(-slice_window);

    $: console.log("data", data);

    $: xScale = scaleTime()
        .domain(extent(sliced_data, (d) => d.time))
        .range([paddingLeft, chartWidth - paddingRight]);
    
    // Define the yScale for power consumption as before
    $: yScalePower = scaleLinear()
        .domain([
            min(sliced_data, (d) => d.powerConsumption),
            max(sliced_data, (d) => d.powerConsumption),
        ])
        .range([chartHeight - paddingBottom, paddingTop]);

    // Define a new yScale for temperature
    $: yScaleTemp = scaleLinear()
        .domain([
            min(sliced_data, (d) => d.temperature),
            max(sliced_data, (d) => d.temperature),
        ])
        .range([chartHeight - paddingBottom, paddingTop]);

    // Define the line generator function for power consumption
    $: lineGeneratorPower = d3.line()
        .x(d => xScale(d.time))
        .y(d => yScalePower(d.powerConsumption))
        .curve(d3.curveMonotoneX);

    // Define the line generator function for temperature
    $: lineGeneratorTemp = d3.line()
        .x(d => xScale(d.time))
        .y(d => yScaleTemp(d.temperature))
        .curve(d3.curveMonotoneX);

    let xAxisPower, yAxisPower, xAxisTemp, yAxisTemp;

    $: if (xAxisPower && yAxisPower) {
        select(xAxisPower).call(axisBottom(xScale))
            .call(g => g.select(".domain").remove());
        select(yAxisPower)
            .attr("transform", `translate(${chartWidth - paddingRight}, 0)`) // Move y-axis to the right
            .call(axisRight(yScalePower)) // Use axisRight to place the ticks and labels on the right
            .call(g => g.select(".domain").remove());
    }

    $: if (xAxisTemp && yAxisTemp) {
        select(xAxisTemp).call(axisBottom(xScale))
            .call(g => g.select(".domain").remove());
        select(yAxisTemp)
            .attr("transform", `translate(${chartWidth - paddingRight}, 0)`)
            .call(axisRight(yScaleTemp))
            .call(g => g.select(".domain").remove());
    }


    // $: console.log("YSCALE", yScale(sliced_data[sliced_data.length-1].powerConsumption))
</script>

<!-- Power Consumption Line Chart -->
<h5>Power</h5>
<svg class="historgram" width={chartWidth} height={chartHeight}>
    <g
        bind:this={xAxisPower}
        transform={`translate(0, ${chartHeight - paddingBottom})`}
        class="x-Axis"
    ></g>
    <g bind:this={yAxisPower} class="y-Axis"></g>

    <!-- Draw the line for power consumption -->
    <path
        d={lineGeneratorPower(sliced_data)}
        fill="none"
        stroke="orange"
        stroke-width="1"
    />
</svg>

<!-- Temperature Line Chart -->
<h5>Temperature</h5>
<svg class="historgram" width={chartWidth} height={chartHeight} style="margin-top: 20px;">
    <g
        bind:this={xAxisTemp}
        transform={`translate(0, ${chartHeight - paddingBottom})`}
        class="x-Axis"
    ></g>
    <g bind:this={yAxisTemp} class="y-Axis"></g>

    <!-- Draw the line for temperature -->
    <path
        d={lineGeneratorTemp(sliced_data)}
        fill="none"
        stroke="cyan"
        stroke-width="1"
    />
</svg>