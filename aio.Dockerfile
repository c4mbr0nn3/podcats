FROM golang:1.21-alpine3.17 AS go-builder
RUN apk update && apk add build-base && rm -rf /var/cache/apk/*
RUN mkdir -p /go/src/podcats
WORKDIR /go/src/podcats
COPY backend/ .
RUN go mod download
RUN go mod verify
RUN go build -o podcats


FROM node:21-alpine3.17 AS vue-builder
RUN mkdir -p /vue
WORKDIR /vue
COPY frontend/ .
RUN npm i -g pnpm
RUN pnpm install
RUN pnpm build

FROM alpine:3.18
ENV GIN_MODE=release
RUN apk -U upgrade && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /go/src/podcats/db
WORKDIR /go/src/podcats
COPY backend/config/*.yaml config/
COPY --from=go-builder /go/src/podcats/podcats .
COPY --from=vue-builder /vue/dist dist/
EXPOSE 8000
CMD ["./podcats","-e","prod"]