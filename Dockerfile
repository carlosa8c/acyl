FROM golang:1.20-alpine

COPY . /go/src/github.com/Pluto-tv/acyl
RUN cd /go/src/github.com/Pluto-tv/acyl && \
CGO_ENABLED=0 go install github.com/Pluto-tv/acyl

FROM alpine:3.18

RUN mkdir -p /go/bin/ /opt/integration /opt/html /opt/migrations && \
apk --no-cache add ca-certificates && apk --no-cache upgrade
COPY --from=0 /go/bin/acyl /go/bin/acyl
COPY --from=0 /go/src/github.com/Pluto-tv/acyl/testdata/integration/* /opt/integration/
COPY --from=0 /go/src/github.com/Pluto-tv/acyl/data/words.json.gz /opt/
COPY --from=0 /go/src/github.com/Pluto-tv/acyl/migrations/* /opt/migrations/
COPY --from=0 /go/src/github.com/Pluto-tv/acyl/ui/ /opt/ui/

ENV MIGRATIONS_PATH /opt/migrations

CMD ["/go/bin/acyl"]
