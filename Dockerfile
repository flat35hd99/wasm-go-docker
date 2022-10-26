FROM golang:1.19.2 AS build
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o server server.go

FROM gcr.io/distroless/static-debian11
COPY --link --from=build /app/server /server
CMD [ "/server" ]
