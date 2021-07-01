FROM golang:latest
RUN apt-get update && apt-get -y install praat 
WORKDIR /root
RUN mkdir src
ADD static/getPitchTier.praat /root/src/getPitchTier.praat
ADD server.go /root/src
CMD go run src/server.go
