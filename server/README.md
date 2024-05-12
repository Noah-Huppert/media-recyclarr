# Server
Media Recyclarr Server.

# Table Of Contents
- [Development](#development)

# Development
## Emby Go library
The Emby Go client library does not distribute a normal package. Instead you must download it to the project: 

1. Download the [Emby API Clients repository](https://github.com/MediaBrowser/Emby.ApiClients/tree/master) ZIP file
2. Extract the ZIP:  
   ```
   unzip Emby.ApiClients-master.zip
   ```
3. Rename the `Clients/Go/` directory from the ZIP to `embyclient-rest-go`:  
   ```
   mv Emby.ApiClients-master/Clients/Go/ embyclient-rest-go
   ```
4. Clean up the download:
   ```
   rm -rf Emby.ApiClients-master
   rm Emby.ApiClients-master.zip
   ```