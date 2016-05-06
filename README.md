# CX to Cytoscape.js Converter

![](docs/cytoscape-flat-logo-orange.png)


## Introduction
This is a simple web API converting CX stream into Cytoscape.js JSON.

## Status

* 5/6/2016Updated for new cxtool


## Build and Deploy

### Requirments
Of course, you can use bare metal machines, but we recommend to use Docker for quick and easy deployment.

The service is tested on the following environment:

* Docker Engine v 1.11.x
* curl (Optional, but usuful to test API)
* jq (Optional. Great tool to generate human-friendly JSON)


### Run in a Docker Container

Make sure you are running latest version of docker engine:

```
~ ❯❯❯ docker -v                                                                    ✱ ◼
Docker version 1.11.1, build 5604cbe
```

#### Quick Start
 
1. ```git clone https://github.com/cyService/service-cxtool.git```
1. ```cd service-cxtool```
1. ```docker build -t cytoscape-ci/service-cxtool .```
1. ```docker run -p 3000:3000 cytoscape-ci/service-cxtool```

Then access root of the server for testing.  It will display basic service information:

```
❯❯❯ curl -v http://localhost:3000/ | jq .                                        ✱ ◼
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 3000 (#0)
> GET / HTTP/1.1
> Host: localhost:3000
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Vary: Origin
< Date: Fri, 06 May 2016 22:50:09 GMT
< Content-Length: 190
< Content-Type: text/plain; charset=utf-8
<
{ [190 bytes data]
100   190  100   190    0     0  27846      0 --:--:-- --:--:-- --:--:-- 31666
* Connection #0 to host localhost left intact
{
  "name": "Cxtool service",
  "version": "v1",
  "build": "05-06-2016",
  "description": "Converts CX format into Cytoscape.js compatible JSON.",
  "documents": "https://github.com/cyService/service-cxtool"
}

``` 

Now you are ready to use cxtool service.

## API Reference

### ```/```

#### Supported methods 
* **GET**

#### Description
Returns basic information of the API

#### Sample Output
```json
{
  "name": "Cxtool service",
  "version": "v1",
  "build": "05-06-2016",
  "description": "Converts CX format into Cytoscape.js compatible JSON.",
  "documents": "https://github.com/cyService/service-cxtool"
}
```


### ```/cx2cyjs```

#### Supported methods 
* **POST**

#### Description
Convert CX JSON into Cytoscape.js compatible format.  You can simply _POST_ complete CX data as the body of request.

#### Sample Client Code

**curl**

```bash
curl -X POST -v -H "content-type:application/json" --data-binary "@my_network.cx" http://localhost:3000/cx2cyjs | jq .
```

**Python**

```python
import json, requests

# URL of your service endpoint
SERVICE_URL = "http://localhost:3000/cx2cyjs"

# Load CX file into memory
with open('my_network.cx') as f:    
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
        }, . . .
    ]
}
```


### ```/ndex2cyjs/:ndex_network_id```

#### Supported methods 
* **GET**

#### Description
Get an NDEx network in Cytoscape.js format.  Next version of [NDEx](http://www.ndexbio.org/) supports API to generate CX from network data sets.  This is an utility function to directly convert their CX data stream into Cytoscape.js compatible JSON.  You can feed the output of this API to draw NDEx network with Cytoscape.js.

#### Parameters

##### ```ndex_network_id```
Globally unique ID of NDEx network.  For example, this network:

```http://dev2.ndexbio.org/#/network/b7eb8b32-ce84-11e5-83ca-0251251672f9```

has ID ```b7eb8b32-ce84-11e5-83ca-0251251672f9```.  If you want this network in Cytoscape.js format, call the following:

```/ndex2cyjs/b7eb8b32-ce84-11e5-83ca-0251251672f9```


#### Query Parameters

##### ```server```
Server type of NDEx.  Default is **public**.  If you want to use some other official NDEx servers, you can pass this parameter.  For example, if you want to use _dev2_ server, the URL will be something like this:

```
http://localhost:3000/ndex2cyjs/a54acf93-1300-11e6-9191-0660b7976219?server=dev2
```

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
    ]
  }
}
```

## Questions?
Please send your questions to _kono at ucsd edu_.

## License
MIT

----
&copy; 2016 The Cytoscape Consortium

Keiichiro Ono