FROM golang:1.21.5

ENV CGO_ENABLED=0

WORKDIR /workspace
ADD . .
RUN go build -o .build/cfztna-tls-beacon -ldflags "-w -s" .

FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=0 /workspace/.build/* ./
ENTRYPOINT ["/app/cfztna-tls-beacon"]
