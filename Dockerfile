FROM node as web-builder
ENV REACT_APP_API_URL=/
RUN mkdir -p /src
WORKDIR /src
RUN git clone https://github.com/long2ice/longurl-web.git
RUN cd longurl-web && npm i && npm run build

FROM golang AS builder
MAINTAINER long2ice "long2ice@gmail.com"
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=$GOARCH
ENV CGO_ENABLED=0
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY --from=web-builder /src/longurl-web/build /build/static
RUN go build -o app ./

FROM scratch
RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /build/app /app/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENTRYPOINT ["/app/app"]