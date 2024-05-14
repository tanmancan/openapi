# openapi

OpenAPI client generator for Go

## Quick start

Place spec file in `./spec/[PACKAGE_NAME]` directory, where `[PACKAGE_NAME]` will be the go package name.

For example:

```
./spec/azurecv/computer-vision.yaml
```

will generate a go module in `./azurecv` with the following package:

```go
// ./azurecv/client.go
package azurecv
...
```

Then run the following script to generate the client:

```bash
./dbuild.sh
```

## Current Client(s)

### Azure Computer Vision

- [OpenAPI Documentation](azurecv/README.md)
- [Official Specs](https://eastus.dev.cognitive.microsoft.com/docs/services/Cognitive_Services_Unified_Vision_API_2024-02-01/operations/61d65934cd35050c20f73ab6)

Notes: This projects uses a modified version of the official spec file to improve the generated go client.

## References

This project uses the [OpenAPI Generator](https://openapi-generator.tech/) Docker image to generate the go clients.
