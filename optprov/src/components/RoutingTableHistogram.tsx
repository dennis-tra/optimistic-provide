import { useEffect } from "react";
import Plot from "react-plotly.js";

interface RoutingTableHistogramProps {
  bucketLevels: number[];
  onBucketSelect: (bucket: number) => void;
}

const RoutingTableHistogram: React.FC<RoutingTableHistogramProps> = (props) => {
  useEffect(() => {}, []);

  return (
    <Plot
      config={{
        displayModeBar: false,
      }}
      onClick={(event) => {
        const bucket = event.points[0].x;
        props.onBucketSelect(bucket as number);
      }}
      data={[
        {
          x: props.bucketLevels.map((_, index) => index),
          y: props.bucketLevels,
          type: "bar",
        },
      ]}
      layout={{
        xaxis: {
          range: [-1, 16],
        },
        yaxis: {
          range: [0, 20],
        },
      }}
    />
  );
};

export default RoutingTableHistogram;
