import React, { useEffect } from "react";
import * as d3 from "d3";
import { ProvideGraph } from "../api/models/ProvideGraph";

interface ProvideChartProps {
  data: ProvideGraph;
}

function renderChartFn(graphData: ProvideGraph) {
  const margin = { top: 30, right: 30, bottom: 30, left: 120 };
  const width = 800 - margin.left - margin.right;
  const height = 500 - margin.top - margin.bottom;

  const conns = graphData.connections;
  const findNodes = graphData.findNodes;
  const addProviders = graphData.addProviders;

  const canvas = d3
    .select("#provide-chart")
    .append("svg")
    .attr("width", width + margin.left + margin.right)
    .attr("height", height + margin.top + margin.bottom)
    .append("g")
    .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

  const uniq = [
    ...new Set(graphData.peers.filter((peer) => peer.firstInteractedAt).map((peer) => peer.peerId.slice(0, 16))),
  ];

  const xScale = d3
    .scaleLinear()
    .domain([0, (new Date(graphData.endedAt).getTime() - new Date(graphData.startedAt).getTime()) / 1000])
    .range([0, width]);
  const xAxisTop = canvas.append("g").attr("class", "xAxis").call(d3.axisTop(xScale));
  const xAxisBottom = canvas
    .append("g")
    .attr("transform", "translate(0," + height + ")")
    .call(d3.axisBottom(xScale));

  // Add X axis label:
  canvas.append("text").attr("text-anchor", "end").attr("x", width).attr("y", -18).text("Duration in s");

  const yScale = d3.scaleBand().domain(uniq).range([0, height]);
  const yAxis = canvas
    .append("g")
    .attr("class", "yAxis")
    .call(d3.axisLeft(yScale))
    .selectAll("text")
    .attr("font-family", "monospace");

  const clip = canvas
    .append("defs")
    .append("svg:clipPath")
    .attr("id", "clip")
    .append("svg:rect")
    .attr("width", width)
    .attr("height", height)
    .attr("x", 0)
    .attr("y", 0);

  // Add brushing
  const brush = d3
    .brushX() // Add the brush feature using the d3.brush function
    .extent([
      [0, 0],
      [width, height],
    ]) // initialise the brush area: start at 0,0 and finishes at width,height: it means I select the whole graph area
    .on("end", updateChart); // Each time the brush selection changes, trigger the 'updateChart' function

  const connsSpans = canvas.append("g").attr("clip-path", "url(#clip)");

  connsSpans
    .selectAll("rect")
    .data(conns)
    .enter()
    .append("rect")
    .attr("x", (val) => xScale((new Date(val.startedAt).getTime() - new Date(graphData.startedAt).getTime()) / 1000))
    .attr("y", (val) => yScale(val.remoteId.slice(0, 16)) + 3)
    .attr("width", (val) => xScale(val.durationInS))
    .attr("height", height / uniq.length - 6)
    .attr("rx", 5)
    .attr("fill", d3[`schemeCategory10`][0] + "aa");

  const findNodesSpans = canvas.append("g").attr("clip-path", "url(#clip)");

  findNodesSpans
    .selectAll("rect")
    .data(findNodes)
    .enter()
    .append("rect")
    .attr("x", (val) => xScale((new Date(val.startedAt).getTime() - new Date(graphData.startedAt).getTime()) / 1000))
    .attr("y", (val) => yScale(val.remoteId.slice(0, 16)) + 3)
    .attr("width", (val) => xScale(val.durationInS))
    .attr("height", height / uniq.length - 6)
    .attr("rx", 5)
    .attr("fill", d3[`schemeCategory10`][1] + "aa");
  const addProvidersSpans = canvas.append("g").attr("clip-path", "url(#clip)");

  addProvidersSpans
    .selectAll("rect")
    .data(addProviders)
    .enter()
    .append("rect")
    .attr("x", (val) => xScale((new Date(val.startedAt).getTime() - new Date(graphData.startedAt).getTime()) / 1000))
    .attr("y", (val) => yScale(val.remoteId.slice(0, 16)) + 3)
    .attr("width", (val) => xScale(val.durationInS))
    .attr("height", height / uniq.length - 6)
    .attr("rx", 5)
    .attr("fill", d3[`schemeCategory10`][2] + "aa");

  canvas.append("g").attr("class", "brush").call(brush);

  canvas
    .selectAll("g.xAxis g.tick")
    .append("line")
    .attr("class", "gridline")
    .attr("x1", 0)
    .attr("y1", 0)
    .attr("x2", 0)
    .attr("y2", height)
    .attr("stroke", "lightgrey")
    .attr("stroke-width", 1);

  canvas
    .selectAll("g.yAxis g.tick")
    .append("line")
    .attr("class", "gridline")
    .attr("x1", 0)
    .attr("y1", 0)
    .attr("x2", width)
    .attr("y2", 0)
    .attr("stroke", "lightgrey")
    .attr("stroke-width", 1);

  const crossHair = canvas.append("g").attr("class", "crosshair");
  crossHair
    .append("line")
    .attr("id", "h_crosshair") // horizontal cross hair
    .attr("x1", 0)
    .attr("y1", 0)
    .attr("x2", 0)
    .attr("y2", 0)
    .style("stroke", "gray")
    .style("stroke-width", "1px")
    .style("stroke-dasharray", "5,5")
    .style("display", "none");

  crossHair
    .append("line")
    .attr("id", "v_crosshair") // vertical cross hair
    .attr("x1", 0)
    .attr("y1", 0)
    .attr("x2", 0)
    .attr("y2", 0)
    .style("stroke", "gray")
    .style("stroke-width", "1px")
    .style("stroke-dasharray", "5,5")
    .style("display", "none");

  crossHair
    .append("text")
    .attr("id", "crosshair_text") // text label for cross hair
    .style("font-size", "10px")
    .style("stroke", "gray")
    .style("stroke-width", "0.5px");

  canvas.on("mousemove", function (event) {
    const coords = d3.pointer(event);
    addCrossHair(coords[0], coords[1]);
  });

  function addCrossHair(xCoord, yCoord) {
    // Update horizontal cross hair
    d3.select("#h_crosshair")
      .attr("x1", 0)
      .attr("y1", yCoord)
      .attr("x2", width)
      .attr("y2", yCoord)
      .style("display", "block");
    // Update vertical cross hair
    d3.select("#v_crosshair")
      .attr("x1", xCoord)
      .attr("y1", 0)
      .attr("x2", xCoord)
      .attr("y2", height)
      .style("display", "block");
    // Update text label
    d3.select("#crosshair_text")
      .attr("transform", "translate(" + (xCoord + 10) + "," + (yCoord - 10) + ")")
      .text(xScale.invert(xCoord).toFixed(3) + "s");
  }
  // A function that set idleTimeOut to null
  let idleTimeout;
  function idled() {
    idleTimeout = null;
  }

  function updateChart(event) {
    const extent = event.selection;

    // If no selection, back to initial coordinate. Otherwise, update X axis domain
    if (!extent) {
      if (!idleTimeout) return (idleTimeout = setTimeout(idled, 350)); // This allows to wait a little bit
      xScale.domain([0, (new Date(graphData.endedAt).getTime() - new Date(graphData.startedAt).getTime()) / 1000]);
    } else {
      xScale.domain([xScale.invert(extent[0]), xScale.invert(extent[1])]);
      canvas.select(".brush").call(brush.move, null); // This remove the grey brush area as soon as the selection has been done
    }

    // Update axis and circle position
    xAxisTop.transition().duration(500).call(d3.axisTop(xScale));
    xAxisBottom.transition().duration(500).call(d3.axisBottom(xScale));

    connsSpans
      .selectAll("rect")
      .transition()
      .duration(500)
      .attr("x", (val) => xScale((new Date(val.startedAt).getTime() - new Date(graphData.startedAt).getTime()) / 1000))
      .attr("width", (val) => xScale(xScale.domain()[0] + val.durationInS));

    findNodesSpans
      .selectAll("rect")
      .transition()
      .duration(500)
      .attr("x", (val) => xScale((new Date(val.startedAt).getTime() - new Date(graphData.startedAt).getTime()) / 1000))
      .attr("width", (val) => xScale(xScale.domain()[0] + val.durationInS));
    addProvidersSpans
      .selectAll("rect")
      .transition()
      .duration(500)
      .attr("x", (val) => xScale((new Date(val.startedAt).getTime() - new Date(graphData.startedAt).getTime()) / 1000))
      .attr("width", (val) => xScale(xScale.domain()[0] + val.durationInS));
  }
}

const ProvideChart: React.FC<ProvideChartProps> = ({ data }) => {
  useEffect(() => {
    renderChartFn(data);
  }, [data]);
  return <div id="provide-chart" />;
};

export default ProvideChart;
