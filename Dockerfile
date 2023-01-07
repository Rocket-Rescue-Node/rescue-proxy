FROM golang:1.19-buster AS build

RUN apt-get update; apt-get install -y make protobuf-compiler gcc libc-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

COPY . /src

WORKDIR /src
RUN make

FROM debian:buster
COPY --from=build /src/rescue-proxy /bin/rescue-proxy
ENTRYPOINT ["/bin/rescue-proxy"]
