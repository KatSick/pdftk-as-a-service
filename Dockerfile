FROM golang:1.12.6 as builder

RUN go get -u github.com/desertbit/fillpdf
RUN go get -u github.com/gin-gonic/gin

ARG version=0.0.0

WORKDIR /usr/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s -X main.VersionString=${version}" -o /go/bin/gopdftkservice

FROM alpine:3.8
RUN apk update && apk upgrade && apk add pdftk
COPY --from=builder /go/bin/gopdftkservice /go/bin/gopdftkservice

EXPOSE 80

ENTRYPOINT ["/go/bin/gopdftkservice"]