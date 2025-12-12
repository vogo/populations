this project package future population tiff file and provide a restful api to query the population

compile and package:
```bash
apt-get update
apt-get install libgdal-dev
make package
```

run:
```bash
go run cmd/server/main.go
```

query population:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"lat": 39.9042, "lng": 116.4074}' http://localhost:8080/api/v1/population
```


