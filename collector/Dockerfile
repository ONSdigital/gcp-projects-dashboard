FROM golang:alpine AS builder
RUN apk update && apk add --no-cache make
WORKDIR $GOPATH/src/github.com/ONSdigital/gcp-projects-dashboard/collector/
COPY . .
RUN make

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/ONSdigital/gcp-projects-dashboard/collector/build/linux-amd64/bin/dashboardcollector /dashboardcollector

ENTRYPOINT [ "/dashboardcollector" ]