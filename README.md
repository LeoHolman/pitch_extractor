# pitch_extractor
A service for extracting pitch from audio blobs sent via HTTP.

## Setup
run `docker build -t pitch_extractor .` in the root directory of this project. This will create the docker image and install `praat`.

## Running
run `docker run -p 8080:8080 pitch_extractor` to start the container.
All requests to localhost:8080 will now forward to the docker container,
which expects a .wav file in an HTTP request, and will attempt to return a 
.csv of the pitch tier. 

## Praat
Praat is the tool that is used to perform pitch extraction.
There is a script called `getPitchTier.praat` that takes in the name of a wav file (without the extension)
and produces a .csv file of the same name. The .csv file has two columns, the first column is the timestamp
in the file. The second column is the pitch at that timestamp. 
