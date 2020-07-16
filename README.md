# Home Automation (Under Active Development - This code base might change drastically)

Distributed home automation system. Largely a learning opportunity rather than a production-ready system.

This home automation system is made up of separate microservices which run on various devices that are distributed physically around my home.
Most of the core services run within a Kubernetes cluster, however there are some periphery services that sit outside of the cluster.
The hardware used is a combination of Raspberry Pis and a Synology NAS.

## Will this work for me?

This is not designed as a general-purpose home automation system. It is pretty specific to my use cases. If you’re looking for something generic, check out [Home Assistant](https://www.home-assistant.io) or [openHAB](https://www.openhab.org). If, however, it does work for you, feel free to use it but don’t expect any support. I also don’t plan to take feature requests. If you would like to make any changes then I suggest forking the repository.

## Getting started

There are various tools that can be installed to aid development.

```shell
./tools/install
```

## Project structure

- `docs/`
  - Bare bone documentation about the system
- `libraries/`
  - Library code shared between all services
- `private/`
  - A git submodule containing mostly private configuration
- `service.x`
  - A backend microservice
- `tools/`
  - Useful tools for working with the system
- `web.x`
  - A web-based application
