# Server
Media Recyclarr Server.

# Table Of Contents
- [Configuration](#configuration)
- [Operations](#operations)

# Configuration
Configuration is set via environment variables:

- `MEDIA_RECYCLARR_EMBY_URL`
- `MEDIA_RECYCLARR_EMBY_API_KEY`
- `MEDIA_RECYCLARR_JELLYSEERR_URL`
- `MEDIA_RECYCLARR_JELLYSEERR_API_KEY`
- `MEDIA_RECYCLARR_POSTGRES_URI`

# Operations
## Database Migrations
Migrate the database to setup the schema. 

To set up the schema run the files in [`migrations/up/`](./migrations/up/) in numerical order.

To tear down the schema run the files in [`migrations/down/`](./migrations/down/) in reverse numerical order.