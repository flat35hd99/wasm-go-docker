FROM golang:1.19.2 AS buildbase
RUN curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v0.26.0/tinygo_0.26.0_amd64.deb \
    && dpkg -i tinygo_0.26.0_amd64.deb

FROM buildbase as build
COPY . /app
WORKDIR /app
RUN tinygo build -o server.wasm -target wasi server.go

FROM scratch
COPY --link --from=build /app/server.wasm /server.wasm
ENTRYPOINT [ "server.wasm" ]
