# CScp

A cheating script made using go. Make your semester easy and focus on the real stuf!!! Checkout the project plan here on [excalidraw](https://excalidraw.com/#room=f09c0fc3888ea1381682,k8t9iyn9PCE7cjbc-YU0JQ)

## Get Started

Clone the repository and run :

```
go mod download
```

Run the project and view log files:

```
go run main.go > app.log 2>&1
```

to view your logs:

```
tail -f app.log
```

To start rabbitmq server:

```
docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```
