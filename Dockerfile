# syntax=docker/dockerfile:1.0.0-experimental

# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM gracepoint/buffalo:0.15.5-1 as go-builder

RUN apk add openssh-client && \
    mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group && \
    apk add --no-cache ca-certificates

# Before we add the whole project in, download just the dependencies, so we can cache that layer and rebuild fast on app changes
COPY go.mod go.sum ./
# We copy the .netrc that needs to be injected for github auth here
COPY .netrc /root/.netrc
RUN go mod download

# since we re-enabled node stuff in buffalo
RUN apk add -u nodejs yarn nghttp2

COPY . .

RUN yarn && buffalo build --ldflags '-s -w -extldflags "-static"' -o /bin/app

FROM alpine

COPY --from=go-builder /user/group /user/passwd /etc/
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /bin/app /app

USER nobody:nobody

ENV GO_ENV=production ADDR=0.0.0.0
EXPOSE 3000

# Uncomment to run the migrations before running the binary:
CMD /app migrate; /app
