FROM golang:latest
RUN apt-get update && apt-get -y install praat 
ADD static/getPitchTier.praat ./getPitchTier.praat
ADD server.go .
CMD go run server.go
