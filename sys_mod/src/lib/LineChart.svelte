<script>
    import { onMount, afterUpdate } from 'svelte';
    import * as d3 from 'd3';
  
    export let data = [];
    export let width = 600;
    export let height = 300;
  
    let svg;

    $: console.log("data", data)
  
    // Function to draw the chart
    
    // function drawChart() {
    //   if (!svg || data.length === 0) return; // Ensure SVG is bound
  
    //   d3.select(svg).selectAll("*").remove(); // Clear previous chart
  
    //   const margin = { top: 20, right: 20, bottom: 30, left: 50 },
    //     chartWidth = width - margin.left - margin.right,
    //     chartHeight = height - margin.top - margin.bottom;
  
    //   const x = d3.scaleLinear().domain(d3.extent(data, d => d.time)).range([0, chartWidth]);
    //   const y = d3.scaleLinear().domain([0, d3.max(data, d => d.value)]).range([chartHeight, 0]);
  
    //   const line = d3.line()
    //     .x(d => x(d.time))
    //     .y(d => y(d.value));
  
    //   const g = d3.select(svg)
    //     .append('g')
    //     .attr('transform', `translate(${margin.left},${margin.top})`);
  
    //   g.append('g')
    //     .attr('transform', `translate(0,${chartHeight})`)
    //     .call(d3.axisBottom(x));
  
    //   g.append('g')
    //     .call(d3.axisLeft(y));
  
    //   g.append('path')
    //     .datum(data)
    //     .attr('fill', 'none')
    //     .attr('stroke', 'steelblue')
    //     .attr('stroke-width', 1.5)
    //     .attr('d', line);
    // }
  
    // onMount(drawChart);
    // afterUpdate(drawChart);
  </script>
  

    <svg>

      {#each data as data }
      
      {/each}
    
    </svg>