# password-generator

##  About
This is a very small microservice that generates passwords using the standard Golang password generator utility.

## Configuration Variables
**RELEASE_MODE**: 
Determines whether the app should run in "prod" mode or "debug" mode.  Choose one of: prod, debug.  Defaults to "debug"

**BIND_HOST**: 
Sets the IP Address to bind to.  Defaults to "0.0.0.0"

**BIND_PORT**: 
Sets the port to bind to.  Defaults to "8080"

**DISABLE_LOGGING_COLOR**:
Disables logging color when set to "true"
    
## Administrative URIs

**/liveness** : Liveness probe endpoint

**/about** : Server build information

**/swagger/index.html** : Swagger sandbox