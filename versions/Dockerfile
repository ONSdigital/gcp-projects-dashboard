FROM golang:alpine AS builder
RUN apk update && apk add --no-cache make
WORKDIR $GOPATH/src/github.com/ONSdigital/gcp-projects-dashboard/versions/
COPY . .
RUN make

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/ONSdigital/gcp-projects-dashboard/versions/build/linux-amd64/bin/gkeversions /gkeversions

ENTRYPOINT [ "/gkeversions" ]