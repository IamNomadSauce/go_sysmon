<script>
    import * as d3 from "d3";
    import { scaleTime, scaleLinear } from "d3-scale";
    import { extent, min, max } from "d3-array";
    import { axisBottom, axisRight } from "d3-axis";
    import { select } from "d3-selection";

    export let data = [];
    export let title = "Chart"
    export let color = "orange"
    export let width = 600;
    export let height = 300;

    let slice_window = 1440;
    let chartWidth = 1200;
    let chartHeight = 300;
    let paddingLeft = 50;
    let paddingRight = 25;
    let paddingTop = 50;
    let paddingBottom = 25;

    $: sliced_data = data.slice(-slice_window);

    $: xScale = scaleTime()
        .domain(extent(sliced_data, (d) => d.time))
        .range([paddingLeft, chartWidth - paddingRight]);
    
    // Define the yScale for power consumption as before
    $: yScale = scaleLinear()
        .domain([
            min(sliced_data, (d) => d.value),
            max(sliced_data, (d) => d.value),
        ])
        .range([chartHeight - paddingBottom, paddingTop]);

    $: lineGeneratorPower = d3.line()
        .x(d => xScale(d.time))
        .y(d => yScale(d.value))
        .curve(d3.curveMonotoneX);

    let xAxisPower, yAxisPower, xAxisTemp, yAxisTemp;

    $: if (xAxisPower && yAxisPower) {
        select(xAxisPower).call(axisBottom(xScale))
            .call(g => g.select(".domain").remove());
        select(yAxisPower)
            .attr("transform", `translate(${chartWidth - paddingRight}, 0)`) // Move y-axis to the right
            .call(axisRight(yScale)) // Use axisRight to place the ticks and labels on the right
            .call(g => g.select(".domain").remove());
    }

</script>

<h5>{title}</h5>
<!-- Add Scatter dots later -->
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
        stroke={color}
        stroke-width="1"
    />
</svg>