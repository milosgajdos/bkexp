# build stage
FROM --platform=${BUILDPLATFORM} golang:1.19-alpine3.17 AS build

RUN apk add --no-cache git build-base ca-certificates

ENV CGO_ENABLED=0
ENV GO111MODULE=auto

ARG GIT_VERSION
ARG TARGETOS
ARG TARGETARCH
ARG PKG=github.com/milosgajdos/bkexp

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags "-X ${PKG}/version.version=${GIT_VERSION}" -o /app/bkexp

# runtime stage
FROM scratch as binary

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/bkexp /bkexp

ENTRYPOINT ["/bkexp"]
CMD ["--help"]
