FROM kairen/golang:1.11-alpine AS build-env
LABEL maintainer="Kyle Bai <k2r2.bai@gmail.com>"

ENV GOPATH "/go"
ENV PROJECT_PATH "$GOPATH/src/github.com/kubedev/kubeconfig-generator"

COPY . $PROJECT_PATH
RUN cd $PROJECT_PATH && \
  dep ensure && \
  make out/kg && \
  mv out/kg /tmp/kg

# Running stage
FROM alpine:3.7
COPY --from=build-env /tmp/kg /bin/kg
ENTRYPOINT ["kg"]
