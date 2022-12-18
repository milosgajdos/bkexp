# bkexp

[![Build Status](https://github.com/milosgajdos/bkexp/workflows/CI/badge.svg)](https://github.com/milosgajdos/bkexp/actions?query=workflow%3ACI)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/milosgajdos/bkexp)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Docker Buildkit experiment that lets you trigger build in remote builders from the cli.

# HOWTO

Build the binary

```shell
$ go build

$ ./bkexp -help
Usage of ./bkexp:
  -builder string
    	buildx builder (default "default")
  -dockerfile string
    	Path to Dockerfile (default "./Dockerfile")
  -image string
    	Image reference
```

Run a build on the preconfigured remote builder
```shell
$ ./bkexp -image "foo/bar" -builder remote02 https://github.com/milosgajdos/bkexp.git
[+] Building 10.6s (14/15)
 => [internal] load git source https://github.com/milosgajdos/bkexp.git                                                                                                                                                                                                                                                                                                                  2.6s
 => [internal] load metadata for docker.io/library/golang:1.19-alpine3.16                                                                                                                                                                                                                                                                                                                1.6s
 => ERROR importing cache manifest from foo/bar                                                                                                                                                                                                                                                                                                                                          0.9s
 => [build 1/8] FROM docker.io/library/golang:1.19-alpine3.16@sha256:4b4f7127b01b372115ed9054abc6de0a0b3fdea224561b354af7bb6e946acaa9                                                                                                                                                                                                                                                    0.0s
 => => resolve docker.io/library/golang:1.19-alpine3.16@sha256:4b4f7127b01b372115ed9054abc6de0a0b3fdea224561b354af7bb6e946acaa9                                                                                                                                                                                                                                                          0.0s
 => CACHED [build 2/8] RUN apk add --no-cache git build-base ca-certificates                                                                                                                                                                                                                                                                                                             0.0s
 => CACHED [build 3/8] WORKDIR /app                                                                                                                                                                                                                                                                                                                                                      0.0s
 => CACHED [build 4/8] COPY go.mod ./                                                                                                                                                                                                                                                                                                                                                    0.0s
 => CACHED [build 5/8] COPY go.sum ./                                                                                                                                                                                                                                                                                                                                                    0.0s
 => CACHED [build 6/8] RUN go mod download                                                                                                                                                                                                                                                                                                                                               0.0s
 => [build 7/8] COPY . ./                                                                                                                                                                                                                                                                                                                                                                0.0s
 => [build 8/8] RUN --mount=type=cache,target=/root/.cache/go-build     GOOS=linux GOARCH=amd64     go build -ldflags "-X github.com/milosgajdos/bkexp/version.version=${GIT_VERSION}" -o /app/bkexp                                                                                                                                                                                     2.4s
 => CACHED [binary 1/2] COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/                                                                                                                                                                                                                                                                                             0.0s
 => [binary 2/2] COPY --from=build /app/bkexp /bkexp                                                                                                                                                                                                                                                                                                                                     0.1s
 => exporting to oci image format                                                                                                                                                                                                                                                                                                                                                        2.7s
 => => exporting layers                                                                                                                                                                                                                                                                                                                                                                  1.6s
 => => exporting manifest sha256:9fc67bb5f4b790ec271cbfaca5801d44f072f3470ef3bc8449c9f4faef04804f                                                                                                                                                                                                                                                                                        0.0s
 => => exporting config sha256:58273a2b74c73869f65b38f04466e0892058332dc85450aa1962730d6a2fdff9                                                                                                                                                                                                                                                                                          0.0s
 => => sending tarball                                                                                                                                                                                                                                                                                                                                                                   1.1s
 => importing to docker
```
