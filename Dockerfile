ARG GO_VERSION=1.18
ARG NODE_VERSION=18
ARG ALPINE_VERSION=3.17

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS go-builder
RUN apk update && apk add build-base && rm -rf /var/cache/apk/*
RUN mkdir -p /go/src/podcats
WORKDIR /go/src/podcats
COPY backend/ .
RUN go mod download
RUN go mod verify
RUN go build -o podcats


FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION} AS vue-builder
RUN mkdir -p /backend
RUN mkdir -p /vue
WORKDIR /vue
COPY frontend/ .
RUN npm ci
RUN npm run build

FROM alpine:${ALPINE_VERSION}
ENV GIN_MODE=release
RUN apk -U upgrade && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /go/src/podcats/db
WORKDIR /go/src/podcats
COPY backend/config config/
COPY --from=go-builder /go/src/podcats/podcats .
COPY --from=vue-builder /backend/dist/ dist/
RUN addgroup --gid 2000 podcats-group && adduser --uid 1000 -D "podcats-user" "podcats-group"
RUN chown -R podcats-user:podcats-group  /go/src/podcats
USER podcats-user:podcats-group
EXPOSE 8000
CMD ["./podcats","-e","prod"]