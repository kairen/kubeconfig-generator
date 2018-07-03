# Building stage
FROM golang:1.10-alpine AS build-env
LABEL maintainer="Kyle Bai <kyle.b@inwinstack.com>"

ENV GOPATH "/go"
ENV PROJECT_PATH "$GOPATH/src/github.com/inwinstack/kubeconfig-generator"

RUN apk add --no-cache git make g++ && \
  go get -u github.com/golang/dep/cmd/dep

COPY . $PROJECT_PATH
RUN cd $PROJECT_PATH && \
  dep ensure && \
  make out/kg && \
  mv out/kg /tmp/kg

# Running stage
FROM alpine:3.7
COPY --from=build-env /tmp/kg /bin/kg
ENTRYPOINT ["kg"]
