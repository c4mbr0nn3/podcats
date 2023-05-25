FROM alpine:3.18
ENV GIN_MODE=release
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /go/src/podcats/db
WORKDIR /go/src/podcats
COPY backend/config/*.yaml config/
COPY backend/build .
COPY backend/dist/ dist/
EXPOSE 8000
CMD ["./podcats","-e","prod"]