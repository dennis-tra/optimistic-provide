import pickle
import os
from typing import List

from model_measurement import Measurement


class ModelLoader:
    parsed_filename = "parsed.p"

    @staticmethod
    def open(folder) -> List[Measurement]:
        """
        open parses the given folder of measurement JSONs and saves
        a dump of the resulting data structure in that folder for
        future fast loading. If that "cached" pickeled version
        already exists it won't parse the folder and just load that.
        It also filters the loaded measurements for measurement instances
        that have 20 ADD_PROVIDER RPCs. See __filter_measurmenets for
        a little more information.
        """
        parsed_path = os.path.join(folder, ModelLoader.parsed_filename)
        if os.path.isfile(parsed_path):
            return ModelLoader.__filter_measurements(ModelLoader.load(parsed_path))

        measurements = ModelLoader.parse(folder)
        ModelLoader.save(folder, measurements)

        return ModelLoader.__filter_measurements(measurements)

    @staticmethod
    def __filter_measurements(measurements: List[Measurement]) -> List[Measurement]:
        """
        I observed some errors in the measurement runs and that 8 out of the >1k measurement runs
        didn't run 20 RPCs to store a provider record but less. I'm excluding them from the final
        measurements as I assume (only assume!) that this is related to these errors.
        """
        filtered_measurements = []
        for measurement in measurements:
            count = 0
            for span in measurement.provider.spans:
                if span.type == "ADD_PROVIDER":
                    count += 1
            if count == 20:
                filtered_measurements += [measurement]
        print(f"Filtered {len(measurements) - len(filtered_measurements)} measurements")
        return filtered_measurements


    @staticmethod
    def load(filename) -> List[Measurement]:
        """
        load unpickles a measurements dump that was created via save.
        """
        with open(filename, "rb") as f:
            print(f"Opening {filename}")
            return pickle.load(f)

    @staticmethod
    def save(folder: str, measurements: List[Measurement]):
        """
        save persists a pickle dump in the given folder of the given measurements.
        It uses a predefined constant file name for the pickle dump.
        """
        filename = os.path.join(folder, ModelLoader.parsed_filename)
        with open(filename, "wb") as f:
            print(f"Saving {len(measurements)} to {filename}")
            pickle.dump(measurements, f)

    @staticmethod
    def parse(folder: str) -> List[Measurement]:
        """
        parse parses a directory of measurement JSON files and returns a list
        of Measurement objects for further processing.
        """
        measurements: List[Measurement] = []
        for fn in os.listdir(folder):
            filepath = os.path.join(folder, fn)
            if not os.path.isfile(filepath) or not fn.endswith("json"):
                continue

            try:
                print("Loading " + filepath)
                measurements += [Measurement.from_file(filepath)]
            except:
                print("Error loading " + filepath)

        return measurements
