FROM golang:1.9 as builder

ADD . /go/src/github.com/jnewmano/public-notices

RUN CGO_ENABLED=0 go install github.com/jnewmano/public-notices


FROM alpine

RUN apk --no-cache add ca-certificates
RUN apk update
RUN apk add poppler-utils

COPY --from=builder /go/bin/public-notices .

CMD ["./public-notices"]
