# Frontend
Media Recyclarr frontend.

# Table Of Contents
- [Development](#development)

# Development
1. Install dependencies: `npm install`
2. Run development mode: `npm run dev`

## HTTP API Client Code Generation
To generate a new version of the HTTP API server API client:

Run: `npm run api-client:generate`

To download a new version of the OpenAPI spec:

1. Make sure the HTTP API server is running  
   (If running on a non-default port set `MEDIA_RECYCLARR_HTTP_API_URI` to the full URI of the server)
2. Run `npm run api-client:download-spec`

To clean up any generated files run `npm run api-client:clean`.  

To perform all these steps run `npm run api-client:dev`