# Graphs

### API to find shortest Path Between two nodes in weighted Graph using Dijkshtra Algorithn

## Steps to Deploy Application
   
```
   git clone
   cd graphs
   docker build -t graphs .
```

## Run Application
 ```
    docker run --net=host graphs:latest
```

## Test Application(Sample API Call)
``` 
    curl -X POST 'localhost:9090/api/v1/task2'  --data '{ "graph": [ { "source": "A", "destination": "B", "weight": 2 }, { "source": "A", "destination": "D", "weight": 1 }, { "source": "A", "destination": "C", "weight": 5 }, { "source": "B", "destination": "C", "weight": 3 }, { "source": "B", "destination": "D", "weight": 2 }, { "source": "D", "destination": "E", "weight": 1 }, { "source": "D", "destination": "C", "weight": 3 }, { "source": "E", "destination": "C", "weight": 1 } , { "source": "E", "destination": "F", "weight": 2 } , { "source": "C", "destination": "F", "weight": 5 } ], "from": "A", "to": "C" }'
```

##### Graphs are Represented as array of json object containing source destination and weights, to find shortest distance pass from: , to:
