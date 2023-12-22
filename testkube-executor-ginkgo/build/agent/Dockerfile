# syntax=docker/dockerfile:1
FROM golang:1.21.5-alpine3.19
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN cd cmd/agent;go build -o /runner -mod mod -a .
RUN go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo@v2.13.2
RUN cp `which ginkgo` ginkgo

FROM golang:1.21.5-alpine3.19
RUN apk update;
RUN apk --no-cache add gcc musl-dev gcompat
RUN apk --no-cache add ca-certificates git
WORKDIR /root/
COPY --from=0 /runner /bin/runner
COPY --from=0 /build/ginkgo /bin/ginkgo
RUN go version
RUN ginkgo version
ENTRYPOINT ["/bin/runner"]
