# Taken from https://github.com/jacobtomlinson/go-container-action/blob/master/Dockerfile
# MIT License.

FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v -o app .

FROM alpine:latest

COPY --from=builder /app/check_yag_tmpl_syntax.json /check_yag_tmpl_syntax.json
COPY --from=builder /app/app /app

ENTRYPOINT ["/app"]