FROM golang:1.16 as base
WORKDIR /app
COPY . .
RUN go get

FROM base as cache
RUN go build -o bin/cache ./cache
CMD ./bin/cache

FROM base as main
RUN go build -o bin/main .
CMD ./bin/main
