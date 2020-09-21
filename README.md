# [![kicksware-api logo][]][Kicksware url]

<p align="center">
	<a href="https://kicksware.com">
		<img src="https://img.shields.io/website?label=Visit%20website&down_message=unavailable&up_color=teal&up_message=kicksware.com%20%7C%20online&url=https%3A%2F%2Fkicksware.com">
	</a>
</p>

[![golang badge]](https://golang.org)&nbsp;
[![lines][lines counter]](https://github.com/timoth-y/kicksware-api)&nbsp;
[![github commit activity][commit activity badge]][repo commit activity]&nbsp;
[![kubernetes badge]](https://kubernetes.io)&nbsp;
[![architecture badge]][microservice article]&nbsp;
[![license badge]](https://www.gnu.org/licenses/agpl-3.0)

[![gitlab badge]](https://ci.kicksware.com/kicksware/kicksware-api)&nbsp;
[![api pipeline]](https://ci.kicksware.com/kicksware/api/-/commits/master)&nbsp;
[![maintainability][maintainability badge]][maintainability source]

## Overview

**Kicksware API** provides both RESTful and gRPC interfaces to deliver access, control, and management of the Kicksware sneaker resale platform.

## Endpoints

### REST

All API's endpoints are divided into 10 base resources which mostly correspond to their source microservices as follows:

| Service            | Server URL        | Endpoint base resources         |
|--------------------|-------------------|---------------------------------|
| users-service      | api.kicksware.com | /users, /auth, /mail, /interact |
| references-service | api.kicksware.com | /references/sneakers            |
| products-service   | api.kicksware.com | /products/sneakers              |
| search-service     | api.kicksware.com | /search                         |
| orders-service     | api.kicksware.com | /orders                         |
| cdn-service        | cdn.kicksware.com | /                               |
| \*                 | \*.kicksware.com   | /health                        |

RESTful API uses `api.kicksware.com` subdomain as it's base server URL.

An exception is the Content Deliver Network (CDN) service,
whitch as its's own subdomain `cdn.kicksware.com`.

API is accessible from both `:443` port via HTTPS and `:80` via plain HTTP.

The full API specification is available on [**Swagger**][swagger] and [**Readme.io**][readme.io].

### gRPC

As it is cloud-native and microservice-based application, Kicksware also provides RPC API using [gRPC framework][grpc] and [Protocol Buffers][protobuf] language.

Like many RPC systems, gRPC is based on the concept of defining a service in terms of functions (methods) that can be called remotely. This approach especially useful for distributed loose-coupled systems, as it provides a mechanism to write API specification ones as a set of `.proto` files and then generate API implementation on any language.

All `.proto` specification files are available [here][proto files].

gRPC API logical division is also based on source microservice entries as follows:

| Service            | Server URL        | Proto endpoints                                                                    |
|--------------------|-------------------|------------------------------------------------------------------------------------|
| users-service      | rpc.kicksware.com | /proto.UserService, /proto.AuthService, /proto.MailService, /proto.InteractService |
| references-service | rpc.kicksware.com | /proto.ReferenceService                                                            |
| products-service   | rpc.kicksware.com | /proto.ProductService                                                              |
| search-service     | rpc.kicksware.com | /proto.SearchReferenceService, /proto.SearchProductService                         |
| orders-service     | rpc.kicksware.com | /proto.OrdersService                                                               |
| cdn-service        | cdn.kicksware.com | /proto.CDNService                                                                  |

## Architecture

## Authentication

[JSON Web Token (JWT)][jwt auth] is used to authenticate and authorize all REST requests.

For accessing gRPC based API both secure TLS connection and [token interceptors][grpc interceptor] are required.

## License

Licensed under the [GNU AGPLv3][license file].

[kicksware-api logo]: https://ci.kicksware.com/kicksware/api/-/raw/master/assets/kicksware-api-logo.png
[kicksware url]: https://kicksware.com

[Website badge]: https://img.shields.io/website?label=Visit%20website&down_message=unavailable&up_color=teal&up_message=kicksware.com%20%7C%20online&url=https%3A%2F%2Fkicksware.com
[golang badge]: https://img.shields.io/badge/Code-Golang-informational?style=flat&logo=go&logoColor=white&color=6AD7E5
[commit activity badge]: https://img.shields.io/github/commit-activity/m/timoth-y/kicksware-api?label=Commit%20activity&color=teal
[repo commit activity]: https://github.com/timoth-y/kicksware-api/graphs/commit-activity
[lines counter]: https://img.shields.io/tokei/lines/github/timoth-y/kicksware-api?color=teal&label=Lines
[license badge]: https://img.shields.io/badge/License-AGPL%20v3-blue.svg?color=teal
[architecture badge]: https://img.shields.io/badge/Architecture-Microservices-informational?style=flat&logo=opslevel&logoColor=white&color=teal
[kubernetes badge]: https://img.shields.io/badge/DevOps-Kubernetes-informational?style=flat&logo=kubernetes&logoColor=white&color=316DE6
[gitlab badge]: https://img.shields.io/badge/CI-Gitlab_CE-informational?style=flat&logo=gitlab&logoColor=white&color=FCA326
[api pipeline]: https://ci.kicksware.com/kicksware/api/badges/master/pipeline.svg?key_text=API%20|%20pipeline&key_width=85
[maintainability badge]: https://api.codeclimate.com/v1/badges/367c3a861b61cc78d24c/maintainability
[maintainability source]: https://codeclimate.com/github/timoth-y/kicksware-api/maintainability

[microservice article]: https://martinfowler.com/articles/microservices.html

[jwt auth]: https://jwt.io/introduction
[grpc interceptor]: https://github.com/grpc/grpc-go/tree/master/examples/features/interceptor

[swagger]: https://app.swaggerhub.com/apis/timoth-y/kicksware-api/1.0.0
[readme.io]: https://kicksware-api.readme.io/reference
[grpc]: https://grpc.io
[protobuf]: https://developers.google.com/protocol-buffers
[proto files]: https://github.com/timoth-y/kicksware-api/tree/master/service-protos

[license file]: https://github.com/timoth-y/kicksware-platform/blob/master/LICENSE