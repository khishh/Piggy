FROM golang:1.19 AS build

WORKDIR /opt/src
COPY ./backend .
WORKDIR /opt/src/
RUN go mod tidy
RUN go build -o server server.go

FROM gcr.io/distroless/base-debian10

COPY --from=build /opt/src/server /
COPY --from=build /opt/src/.env /

# RUN apt-get update && apt-get install -y tail

EXPOSE 8080
ENTRYPOINT [ "/server" ]

