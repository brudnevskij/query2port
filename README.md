# Query2Port Traefik Plugin

## Overview
`query2port` is a Traefik middleware plugin that modifies the destination port of HTTP requests based on a query parameter.

## Installation

Enable plugins in `traefik.yml`:
```yaml
experimental:
  plugins:
    query2port:
      moduleName: "github.com/brudnevskij/query2port"
      version: "v0.1.0"
```

Add the middleware:
```yaml
http:
  middlewares:
    query2port-middleware:
      plugin:
        query2port:
          queryParamName: "port"
```

Attach it to a router:
```yaml
http:
  routers:
    my-router:
      rule: "Host(`example.com`)"
      middlewares:
        - query2port-middleware
      service: my-service
```

## Configuration

| Parameter        | Description                                    |
|-----------------|------------------------------------------------|
| `queryParamName` | Query parameter that specifies the port. |

## Usage
Requests with the query parameter update the destination port dynamically:
- `GET http://example.com?port=8081` → Forwards to `example.com:8081`
- `GET http://example.com` → No change
