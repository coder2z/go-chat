# SocketSvr Service

This is the SocketSvr service

Generated with

```
micro new --namespace=go.micro.chat.socket --type=service socket/socket-svr
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.chat.socket.service.socket-svr
- Type: service
- Alias: socket-svr

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
./socket-svr-service
```

Build a docker image
```
make docker
```