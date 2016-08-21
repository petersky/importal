# Importal
The Immortal Portal

Importal is a simple application portal which provides a web UI for REST-based applications. One or more applications may be added to a running Importal instance, all receiving their authentication and authorization information from Importal without any UI code needed in the REST application server.

The UI is created based on a specification received from the REST server, using a custom object definition in JSON.

# Authentication

Importal provides its own user authentication, or may be configured to use another service such as LDAP.

# Authorization

Importal provides its own user authorization, and may also be configured to use another service in addition to or in replacement of.

# API

## Application API

Importal expects the REST application server to respond appropriately to the following API endpoints. Endpoints may contain a prefix if defined in the application instantiation.

/status

    Request headers:
        X-IMPORTAL-VERSION
        X-IMPORTAL-PUBLIC-KEY
        X-IMPORTAL-REQUEST-TIME
        X-IMPORTAL-REQUEST-HASH
    Returns:
    {
        "version": "X.Y.Z",  // If a version change occurs, Importal may request the user to reload the web page. If configured, Importal may pass the server version along with other API calls as ?api_version=X.Y.Z (parameter name can be configured)
        "status": "status text", // Shown next to the application in the UI. If this is missing, status "Unknown" is shown. If the API call to status times out, status "Unavailable" is shown.
    }

/importalConfig

    Returns:
    json UI definition object

All other proxied API calls contain the following request headers

    X-IMPORTAL-VERSION
    X-IMPORTAL-PUBLIC-KEY
    X-IMPORTAL-REQUEST-TIME
    X-IMPORTAL-REQUEST-HASH
    X-IMPORTAL-AUTHENTICATED-USER
    X-IMPORTAL-AUTHORIZED-ROLES
    
    The X-IMPORTAL-REQUEST-HASH is created as md5 of
    
    X-IMPORTAL-PUBLIC-KEY+";"+X-IMPORTAL-REQUEST-TIME+";"+X-IMPORTAL-AUTHENTICATED-USER+";"+X-IMPORTAL-AUTHORIZED-ROLES+md5(request-method+request-query-string)+md5(request-fields)
    
    The REST application server should validate this field to ensure the received request is authentic.

## UI JSON

When a user selects an application in Importal they are presented with a UI defined by the object received from /importalConfig. This complex object defines all the views, models, collections, and behaviors of the application. The UI is specified by what is desired. Importal will create appropriate HTML, css, and javascript methods to achieve the desired result.

For instance, instead of saying "Click button X to show screen Y", the config will state "X is a list of Y, new Y can be added, existing Y and be deleted and modified". "Y is a set of A, B, and C". When modifying Y, Importal will display appropriate widgets for modifying attributes A, B, and C. This section is very much a WIP.

An Importal application is essentially a UI for interacting with an API. As such, it considers an API as an event-driven state machine. A 

### Request Object

Optionally, the UI may send a POST instead of a GET request for the config. The POST will contain a JSON object detailing the capabilities of the UI. The API should return a config that is appropriate for the capabilities published.

### Response Object
```json
{
    "models": {
        "model01_name": {
            },
        "model02_name": {
            },
        }
    },
    "collections": {
        "collection01_name": {
            "model": "model01_name",
        }
    },
}
```