# Finominal

## Gateway

### Introduction
Gateway is a L7 proxy service for a microservice/distributed system. It provides routing, monitoring/observability, security. 

### Features
- HTTP L7 routing: Match the corresponding service by setting parameters such as location, query, header, host and method
- Health checking
- Load balancing
- Automatic retries
- Circuit breaking/Outlier detection
- Rate limiting 
- Request racing
- REST API, full access to the internals
- CORS Filter,
- Auth: JWT, Basic Auth
- Security
- Response Aggregation
- Request|Response Mutation: rewriting of 'scheme', 'URI', 'host', and adding or deleting the value of the request header of the forwarding request
- Request shadowing

### Resources
https://github.com/motiv-labs/janus
https://www.krakend.io/docs/overview
https://itnext.io/why-should-you-write-your-own-api-gateway-from-scratch-378074bfc49e

### Reference
It's based on [Go Buffalo Boilerplate](https://github.com/chsqur/boilerplate-go)