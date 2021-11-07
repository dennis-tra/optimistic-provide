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
        """
        parsed_path = os.path.join(folder, ModelLoader.parsed_filename)
        if os.path.isfile(parsed_path):
            return ModelLoader.load(parsed_path)

        measurements = ModelLoader.parse(folder)
        ModelLoader.save(folder, measurements)
        return measurements

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


if __name__ == '__main__':
    measurements = ModelLoader.load("save.p")
    print(measurements)
