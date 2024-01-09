ARG app=webapp-playdemo

FROM golang:1.21 as build
ARG app
WORKDIR /build

COPY go.mod ./
RUN go mod download
COPY . .

## Copy goreleaser compiled binary
RUN { cp -v dist/${app}_linux_*/${app} dist/app; chmod +x dist/app; } || :

## Build conditonally (unless dist/${app})
RUN [ -x dist/app ] || CGO_ENABLED=0 GOOS=linux go build -o dist/app -v cmd/${app}/main.go

FROM scratch
COPY --from=build /build/dist/app /app
CMD [ "/app" ]
