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
Run the project in detached mode:

```
go run main.go -d
```

Add default prompt (helpful while running in detached mode):

```
go run main.go -d -p "Your default prompt"
```