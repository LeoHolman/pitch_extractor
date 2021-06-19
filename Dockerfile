FROM ubuntu:latest
RUN apt-get update && apt-get -y install praat
RUN echo 'form File input and output\n\tword hash\nendform\nRead from file: hash$ + ".wav"\nselectObject: "Sound " + hash$\nTo Manipulation: 0.01, 75, 600\nExtract pitch tier\nSave as PitchTier spreadsheet file: hash$ + ".csv"' > getPitchTier.praat
CMD bash
