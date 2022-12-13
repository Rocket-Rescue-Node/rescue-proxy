FROM golang:1.18-alpine AS build

RUN apk add --update make protobuf-dev gcc libc-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

COPY . /src

WORKDIR /src
RUN make

FROM alpine:latest
COPY --from=build /src/rescue-proxy /bin/rescue-proxy
ENTRYPOINT /bin/rescue-proxy
