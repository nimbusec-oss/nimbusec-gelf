# --- BUILD
FROM golang:1.9

WORKDIR /go/src/github.com/nimbusec-oss/nimbusec-gelf/
COPY . .
RUN go install

# --- RUN
FROM debian:9
LABEL maintainer "Thomas Kastner <thomas@nimbusec.com>"

# "run config"
EXPOSE 8080
ENV PORT=8080

WORKDIR /app
CMD ["/app/nimbusec-gelf"]

# add artifacts
COPY --from=0 /go/bin/nimbusec-gelf /app/
