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
## Emby Go library
The Emby Go client library does not distribute a normal package. Instead you must download it to the project: 

1. Download the [Emby API Clients repository](https://github.com/MediaBrowser/Emby.ApiClients/tree/master) ZIP file
2. Extract the ZIP:  
   ```
   unzip Emby.ApiClients-master.zip
   ```
3. Rename the `Clients/Go/` directory from the ZIP to `embyclient`:  
   ```
   mv Emby.ApiClients-master/Clients/Go/ embyclient
   ```
4. Rename the Go package in the client source code:
   ```
   sed -i 's/embyclient-rest-go/embyclient/g' *.go
   ```
   This is due to an error with Emby's auto generated code where they try to use a package name with hyphens.
5. Clean up the download:
   ```
   rm -rf Emby.ApiClients-master
   rm Emby.ApiClients-master.zip
   ```

## Jellyseerr Go Library
The Jellyseerr Go client is auto-generated from their swagger API definition.

The API definition is found is found [here](https://github.com/Fallenbagel/jellyseerr/blob/develop/overseerr-api.yml) and downloaded to `jellyseerrclient/openapi/overseerr-api.yml`

The [OpenAPI Generator tool](https://openapi-generator.tech/) is used to generate the client. Run the following command from the `server/` directory:

```
openapi-generator generate --config ./jellyseerrclient/openapi/openapi-generator.yml
```