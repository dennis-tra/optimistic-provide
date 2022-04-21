import re 
import sys

pattern = r".*:.*: ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\), ([\d\.]+) \(\d+\)"


out = open(sys.argv[1] + ".csv", "w")
with open(sys.argv[1], "r") as f:
	lines = f.readlines()
	for line in lines:
		match = re.search(pattern, line)
		if match is None:
			continue

		for i in range(1, 21):
			if i != 1:
				out.write(",")	
			out.write(match.group(i))
		out.write("\n")

out.close()