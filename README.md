# CX to Cytoscape.js Converter

![](docs/cytoscape-flat-logo-orange.png)

## Introduction
This is a simple web service converting CX stream into JSON in Cytoscape.js format.

## Warning!
This service is experimental because it depends on development version of NDEx.

## Build and Deploy

### Requirments
Of course, you can use bare metal machines, but we recommend to use Docker for quick and easy deployment.

The service is tested on the following environment:

* Docker Engine v 1.10.x
* curl (Optional, but usuful to test API)
* jq (Optional. Great tool to generate human-friendly JSON)


### Run in a Docker Container

Make sure you are running latest version of docker engine.

```
~ ❯❯❯ docker -v
Docker version 1.10.0, build 590d5108
```

#### Quick Start
 
1. ```git clone https://github.com/cytoscape-ci/service-cxtool.git```
1. ```cd service-cxtool```
1. ```docker build -t cytoscape-ci/service-cxtool .```
1. ```docker run -p 3000:3000 cytoscape-ci/service-cxtool```

Then access root of the server for testing.  It will display basic service information:

```
❯❯❯ curl -v 192.168.99.100:3000 | jq .                                                                           ⏎
* Rebuilt URL to: 192.168.99.100:3000/
*   Trying 192.168.99.100...
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* Connected to 192.168.99.100 (127.0.0.1) port 3000 (#0)
> GET / HTTP/1.1
> Host: 192.168.99.100:3000
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Vary: Origin
< Date: Thu, 11 Feb 2016 23:35:54 GMT
< Content-Length: 172
< Content-Type: text/plain; charset=utf-8
<
{ [172 bytes data]
100   172  100   172    0     0    927      0 --:--:-- --:--:-- --:--:--   924
* Connection #0 to host 192.168.99.100 left intact
{
  "name": "CXTOOL service",
  "version": "v1",
  "description": "Converts CX format into Cytoscape.js compatible JSON.",
  "documents": "https://github.com/cytoscape-ci/service-cxtool"
}
``` 

Now you are ready to use cxtool service.

## API Reference

### ```/```

#### Supported methods 
* **GET**

#### Description
Returns basic information of the service.

#### Sample Output
```json
{
  "name": "CXTOOL service",
  "version": "v1",
  "description": "Converts CX format into Cytoscape.js compatible JSON.",
  "documents": "https://github.com/cytoscape-ci/service-cxtool"
}
```

----

### ```/cx2cyjs```

#### Supported methods 
* **POST**

#### Description
Convert CX JSON into Cytoscape.js compatible format.  You can simply _POST_ complete CX data as the body of request.

#### Sample Client Code
**Python**

```python
import json, requests

# URL of your service endpoint
SERVICE_URL = "http://192.168.99.100:3000/cx2cyjs"

# Load CX file into memory
with open('my_cx_file.cx') as f:    
    cx = json.load(f)

# POST it to service
res = requests.post(SERVICE_URL, json=cx)
print( json.dumps(res.json(), indent=4))
```

#### Sample Output
Output is a Cytoscape.js compatible JSON.  Note that the result may include _style_ object in addition to network potion of JSON if original CX data contains style aspect.
 
```json
{
    "style": [
        {
            "selector": "node",
            "css": {
                "background-opacity": 255,
                "border-opacity": 255,
                "width": 70,
                "height": 30,
                "font-family": "HelveticaNeue-UltraLight",
                "border-width": 0,
                "background-color": "#33FFFF",
                "color": "#333333",
                "font-size": 12,
                "shape": "roundrectangle",
                "text-opacity": 255
            }
        },...
}
```

----

### ```/ndex2cyjs/:ndex_network_id```

#### WARNING
**This API is experimental!!**
Since this function depends on development version of external API (development version of NDEx), it may not work until they publish the final version.

#### Supported methods 
* **GET**

#### Description
Get an NDEx network in Cytoscape.js format.  Next version of [NDEx]() supports API to generate CX from network data sets.  This is an utility function to directly convert their CX data stream into Cytoscape.js compatible JSON.  You can feed the output of this API to draw NDEx network with Cytoscape.js.

#### Parameters

##### ```ndex_network_id```
Globally unique ID of NDEx network.  For example, this network:

```http://dev2.ndexbio.org/#/network/b7eb8b32-ce84-11e5-83ca-0251251672f9```

has ID ```b7eb8b32-ce84-11e5-83ca-0251251672f9```.  If you want this network in Cytoscape.js format, call the following:

```/ndex2cyjs/b7eb8b32-ce84-11e5-83ca-0251251672f9```

#### Sample Output

```json
{
  "data": {},
  "elements": {
    "nodes": [
      {
        "data": {
          "Degree": 0,
          "KAM_COMPILE_DATE": "1340385379000",
          "KAM_NAME": "Large Corpus",
          "KAM_NODE_FUNCTION": "PROTEIN_ABUNDANCE",
          "KAM_NODE_ID": "NAAAABAAAKIw",
          "KAM_NODE_LABEL": "p(MGI:Rbl2)",
          "WSDL_URL": "http://demo.openbel.org/openbel-ws/belframework.wsdl",
          "canonicalName": "p(MGI:Rbl2)",
          "id": "3509",
          "name": "p(MGI:Rbl2)",
          "selected": false,
          "shared name": "p(MGI:Rbl2)",
          "vizmap:Network 0 NODE_LABEL": "p(MGI:Rbl2)",
          "vizmap:Network 0 NODE_SIZE": "40.0"
        },
        "position": {
          "x": 0,
          "y": 0
        },
        "selected": false
      },...
```

## Questions?
Please send your questions to _kono at ucsd edu_.

## License
MIT

----
&copy; 2016 The Cytoscape Consortium

Keiichiro Ono