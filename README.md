# Finominal

## Gateway

### Introduction
Gateway is a L7 proxy service for a microservice/distributed system. It provides routing, monitoring/observability, security. 

### Features
- HTTP L7 routing using request method, host, port and path only
+ Enforce security like SSL and CORS by default
+ Allow request, response? mutation
+ Auto translate responses
+ Cache route definition in Redis for faster reads
+ Health checking for registered routes
+ REST API, full access to the internals

### Resources
https://github.com/motiv-labs/janus
https://www.krakend.io/docs/overview
https://itnext.io/why-should-you-write-your-own-api-gateway-from-scratch-378074bfc49e

### Todo
- Load balancing
- Automatic retries
- Circuit breaking/Outlier detection
- Rate limiting 
- Request racing
- Response Aggregation
- Request shadowing
### Reference
It's based on [Go Buffalo Boilerplate](https://github.com/chsqur/boilerplate-go)

