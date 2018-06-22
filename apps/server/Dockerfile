# Building stage
FROM golang:1.10-alpine AS build-env
LABEL maintainer="Kyle Bai <kyle.b@inwinstack.com>"

ENV GOPATH "/go"
ENV PROJECT_URL "https://github.com/inwinstack/kubeconfig-generator.git"
ENV PROJECT_PATH "$GOPATH/src/github.com/inwinstack/kubeconfig-generator"

RUN apk add --no-cache git make g++ && \
  go get -u github.com/golang/dep/cmd/dep

RUN git clone -b nfu $PROJECT_URL $PROJECT_PATH && \
  cd $PROJECT_PATH && \
  dep ensure && \
  make out/kg && \
  mv out/kg /tmp/kg

# Running stage
FROM alpine:3.7
COPY --from=build-env /tmp/kg /bin/kg
ENTRYPOINT ["kg"]
