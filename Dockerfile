FROM golang:1.21-alpine AS build
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /sysinfo ./

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /sysinfo /usr/local/bin/sysinfo
EXPOSE 3000
ENTRYPOINT ["/usr/local/bin/sysinfo"]
