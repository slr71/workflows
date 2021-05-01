FROM golang:1.16-alpine

RUN apk add --no-cache make
RUN apk add --no-cache git
RUN go get -u github.com/jstemmer/go-junit-report

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/cyverse-de/workflows
COPY . .
RUN make

FROM scratch

WORKDIR /app

COPY --from=0 /go/src/github.com/cyverse-de/workflows/workflows /bin/workflows
COPY --from=0 /go/src/github.com/cyverse-de/workflows/swagger.json swagger.json

ENTRYPOINT ["workflows"]

EXPOSE 8080
