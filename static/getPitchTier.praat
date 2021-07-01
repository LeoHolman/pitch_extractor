form File input and output
    word hash
endform
Read from file: hash$ + ".wav"
selectObject: "Sound " + hash$
To Manipulation: 0.01, 75, 600
Extract pitch tier
Save as PitchTier spreadsheet file: hash$ + ".csv"