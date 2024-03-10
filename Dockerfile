FROM golang:1.22-alpine AS build

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags=-w -o /bin/esquilo-aniquilador ./cmd/api/main.go

FROM scratch AS production

COPY --from=build /bin/esquilo-aniquilador /bin/esquilo-aniquilador

ENTRYPOINT [ "/bin/esquilo-aniquilador" ]
