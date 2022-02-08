FROM golang:1.16 as cache
WORKDIR /app
COPY . .
RUN go build -o bin/cache ./cache
CMD ./bin/cache

FROM golang:1.16 as main
WORKDIR /app
COPY . .
RUN go build -o bin/main .
CMD ./bin/main
