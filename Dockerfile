FROM golang:1.25-bookworm AS build

RUN apt-get update; apt-get install -y make protobuf-compiler gcc libc-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    make ./rescue-proxy ./api-client

FROM debian:bookworm
COPY --from=build /src/rescue-proxy /bin/rescue-proxy
COPY --from=build /src/api-client /bin/api-client
ENTRYPOINT ["/bin/rescue-proxy"]
