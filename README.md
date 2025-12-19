this project package future population tiff file and provide a restful api to query the population

1. download futurepop ref [doc/future_pop.md]

2. compile and package:
```bash
apt-get update
apt-get install libgdal-dev
make package
```

3. run

```bash
export FUTURE_POP_TIFF_PATH=/absolute/path/to/your_futurepop.tif
export SERVER_PORT=:8080
go run cmd/server/main.go
```

4. query population
```bash
curl -X POST -H "Content-Type: application/json" -d '{"lat": 39.9042, "lng": 116.4074}' http://localhost:8080/api/v1/population
```

