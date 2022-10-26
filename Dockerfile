FROM golang:1.19.2 AS build
COPY . /app
WORKDIR /app
RUN go build -o server server.go

FROM gcr.io/distroless/static-debian11
COPY --link --from=build /app/server /server
CMD [ "/server" ]
