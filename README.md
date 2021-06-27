# pitch_extractor
A service for extracting pitch from audio blobs sent via HTTP.

## Setup
run `docker build .` in the root directory of this project. This will create the docker image and install `praat`.

## Praat
Praat is the tool that is used to perform pitch extraction.
There is a script called `getPitchTier.praat` that takes in the name of a wav file (without the extension)
and produces a .csv file of the same name. The .csv file has two columns, the first column is the timestamp
in the file. The second column is the pitch at that timestamp. 
