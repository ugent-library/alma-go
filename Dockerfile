# build stage
FROM golang:alpine AS build
WORKDIR /build
COPY . .
RUN go get -d -v ./...
RUN go build -o app -v cmd/alma/*

# final stage
FROM alpine:latest
WORKDIR /dist
COPY --from=build /build .
ENTRYPOINT ["/dist/app"]  
CMD ["--help"]
