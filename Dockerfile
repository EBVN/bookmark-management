FROM golang:1.26-alpine AS base

RUN mkdir -p /opt/app

WORKDIR /opt/app

COPY . .

RUN go mod download

FROM base AS build

RUN go build -o main ./cmd/api/main.go

FROM base AS test-exec

ARG _outputdir="/tmp/coverage"
ARG COVERAGE_EXCLUDE

RUN mkdir -p ${_outputdir} && \
    go test ./... -coverprofile=coverage.tmp -covermode=atomic -coverpkg=./... -p 1 && \
    grep -v -E "${COVERAGE_EXCLUDE}" coverage.tmp > ${_outputdir}/coverage.out && \
    go tool cover -html=${_outputdir}/coverage.out -o ${_outputdir}/coverage.html

FROM scratch AS test

ARG _outputdir="/tmp/coverage"

COPY --from=test-exec ${_outputdir}/coverage.out /
COPY --from=test-exec ${_outputdir}/coverage.html /

FROM alpine:3.23.2 AS run

WORKDIR /app

COPY --from=build /opt/app/main .
COPY --from=build /opt/app/docs .

CMD ["/app/main"]
