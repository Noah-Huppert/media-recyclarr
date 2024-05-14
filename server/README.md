# Server
Media Recyclarr Server.

# Table Of Contents
- [Configuration](#configuration)
- [Development](#development)

# Configuration
Configuration is set via environment variables:

- `MEDIA_RECYCLARR_EMBY_URL`
- `MEDIA_RECYCLARR_EMBY_API_KEY`
- `MEDIA_RECYCLARR_JELLYSEERR_URL`
- `MEDIA_RECYCLARR_JELLYSEERR_API_KEY`

# Development
## Code Generation
Some API clients are not distributed and instead we rely on code generation. The [OpenAPI Generator tool](https://openapi-generator.tech/) is used for this.

To clean generated files run:

```
./scripts/clean-oapi-generated.sh embyclient|jellyseerrclient
```

## Emby Go library
The [`embyclient/openapi/openapi.json`](./embyclient/openapi/openapi.json) is downloaded from [Emby's Swagger API browser](https://swagger.emby.media/openapi.json).

Then run:

```
./scripts/oapi-generate.sh embyclient
sed -i 's/ModelModel/Model/g' embyclient/*.go
```

The `sed` replacement is required to fix a bug where the `modelNamePrefix` option for openapi-generator is duplicated sometimes.

## Jellyseerr Go Library
The [`jellyseerrclient/openapi/overseerr-api.yml`](./jellyseerrclient/openapi/overseerr-api.yml) is downloaded from [Jellyseerr's Git repository](https://github.com/Fallenbagel/jellyseerr/blob/develop/overseerr-api.yml).

```
./scripts/oapi-generate.sh jellyseerrclient
```