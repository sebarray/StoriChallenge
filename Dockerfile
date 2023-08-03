FROM golang:alpine AS build


ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src/api
COPY . .

RUN GOOS=linux go build -o /go/bin/api cmd/main.go

EXPOSE 8080

FROM alpine
COPY --from=build /go/bin/api /go/bin/api
COPY --from=build /go/src/api/.env /go/bin/.env
ENTRYPOINT ["go/bin/api"]