ARG GO_VERSION=1.18
ARG NODE_VERSION=18
ARG ALPINE_VERSION=3.16

FROM golang:${GO_VERSION}-alpine AS go-builder
RUN apk update && apk add build-base
RUN mkdir -p /go/src/podcats
WORKDIR /go/src/podcats
COPY backend/ .
RUN go mod download
RUN go mod verify
RUN go build -o podcats


FROM node:${NODE_VERSION}-alpine AS vue-builder
RUN mkdir -p /backend
RUN mkdir -p /vue
WORKDIR /vue
COPY frontend/ .
RUN npm ci
RUN npm run build

FROM alpine:${ALPINE_VERSION}
RUN mkdir -p /go/src/podcats
WORKDIR /go/src/podcats
COPY backend/config config/
COPY --from=go-builder /go/src/podcats/podcats .
COPY --from=vue-builder /backend/dist/ dist/

EXPOSE 8000
CMD ["./podcats","-e","prod"]