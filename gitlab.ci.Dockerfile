FROM alpine:3.17
ENV GIN_MODE=release
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /go/src/podcats/db
WORKDIR /go/src/podcats
COPY backend/config/*.yaml config/
COPY backend/build .
COPY backend/dist/ dist/
RUN addgroup --gid 2000 podcats-group && adduser --uid 1000 -D "podcats-user" "podcats-group"
RUN chown -R podcats-user:podcats-group  /go/src/podcats
USER podcats-user:podcats-group
EXPOSE 8000
CMD ["./podcats","-e","prod"]