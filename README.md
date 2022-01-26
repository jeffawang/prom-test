prom-test
===

This repo is a simple example of go + prometheus in docker.

Run the example with:

```
docker-compose up
```

Once it's up, the go http server will listen on the host at `:8080`, and you can visit prometheus at `:9090`.

Some other stuff
---

```
# run the containers in the background
docker-compose up -d

# see what containers are running
docker-compose ps

# Clean up everything we made (containers, images, volumes, networks)
docker-compose down --rmi local -v
```
