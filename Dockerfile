############################
# STEP 1 build executable binary
############################
FROM golang:alpine
#FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN mkdir /app
RUN mkdir /app/src
WORKDIR /app/src
COPY . .

RUN go build -o /app/main cmd/app/main.go
RUN rm -rf /app/src
RUN rm -rf /go
ENTRYPOINT ["/app/main"]
# Profiler
#ENTRYPOINT ["go","run", "cmd/app/main.go"] 

