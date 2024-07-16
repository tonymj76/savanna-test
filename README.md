- To install dep `go get .`

Set the following env data

```
HOST=host
PORT=port
DB_USER=user
DB_PASSWORD=pass
DB_NAME=db
DB_HOST=localhost
DB_POST=5432
DB_SSLMODE=disable
```

To run server, use `air` for live reload 
- ensure make binary is installed
- Run the command 
`make update`
OR
- run `docker-compose up`

## Run with docker
 - have both docker and docker compose install
 - most PC's or MAC's have make installed 

use `make run` to run in container and `make log` to show logs

using entgo https://entgo.io/docs/getting-started/
