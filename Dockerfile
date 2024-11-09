FROM golang:1.23.2-alpine3.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Building the binary
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o webui .


FROM alpine:3.20 AS main
COPY --from=builder /app/webui /bin/
COPY engine-ui/ /app/
WORKDIR /app

ENV RunInDocker=true
ENV PORT=9000

LABEL maintainer="Biswajit Das <billionto@gmail.com>, Mohan Mal <mohanmal553@gmail.com>"
LABEL org.opencontainers.image.version="0.0.3-alpha"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source=https://github.com/BiltuDas1/search-core
LABEL org.opencontainers.image.documentation=https://github.com/BiltuDas1/search-core/wiki

CMD [ "webui" ]
