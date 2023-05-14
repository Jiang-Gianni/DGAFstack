#Build Stage
FROM golang:1.18-alpine3.16 AS BuildStage

WORKDIR /app

COPY go.mod ./

COPY . ./

RUN go build -o /test

EXPOSE 4716


# Deploy Stage
FROM alpine

WORKDIR /

COPY --from=BuildStage /test /test

EXPOSE 4716

ENTRYPOINT ["/test"]