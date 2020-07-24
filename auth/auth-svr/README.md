# AuthSvr Service

This is the AuthSvr service

Generated with

```
micro new --namespace=go.micro.chat.auth --type=service auth/auth-svr
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.chat.auth.service.auth-svr
- Type: service
- Alias: auth-svr

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
./auth-svr-service
```

Build a docker image
```
make docker
```