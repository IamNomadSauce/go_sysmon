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
    
    $: yScale = scaleLinear()
        .domain([
            min(sliced_data, (d) => d.powerConsumption),
            max(sliced_data, (d) => d.powerConsumption),
        ])
        .range([chartHeight - paddingBottom, paddingTop]); // Inverted range for y-axis

    // Define the line generator function
    $: lineGenerator = d3.line()
        .x(d => xScale(d.time))
        .y(d => yScale(d.powerConsumption))
        .curve(d3.curveMonotoneX); // This makes the line smooth
    let xAxis, yAxis;

    $: if (xAxis && yAxis) {
        select(xAxis).call(axisBottom(xScale))
            .call(g => g.select(".domain").remove()); // Remove the x-axis line
        select(yAxis)
            .attr("transform", `translate(${paddingLeft}, 0)`)
            .call(axisRight(yScale))
            .call(g => g.select(".domain").remove()); // Remove the y-axis line
    }


    $: console.log("YSCALE", yScale(sliced_data[sliced_data.length-1].powerConsumption))
</script>

<svg class="historgram" width={chartWidth} height={chartHeight}>
    <g
        bind:this={xAxis}
        transform={`translate(0, ${chartHeight - paddingBottom})`}
        class="x-Axis"
    ></g>
    <g bind:this={yAxis} class="y-Axis"></g>

    <!-- Draw the line for the line chart -->
    <path
        d={lineGenerator(sliced_data)}
        fill="none"
        stroke="orange"
        stroke-width="1"
    />

    <!-- Optionally, if you want to keep the circles on the line chart -->
    <g>
        {#each sliced_data as val, i}
            <circle
                cx={xScale(val.time)}
                cy={yScale(val.powerConsumption)}
                r="1"
                fill="gold"
            />
        {/each}
    </g>
</svg>