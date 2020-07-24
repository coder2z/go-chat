# AuthWeb Service

This is the AuthWeb service

Generated with

```
micro new --namespace=go.micro.api.auth --type=web auth/auth-web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.auth.web.auth-web
- Type: web
- Alias: auth-web

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./auth-web-web
```

Build a docker image
```
make docker
```