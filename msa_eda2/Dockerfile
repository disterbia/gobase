# Dockerfile
FROM golang:1.17-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch
COPY --from=build /bin/app /bin/app
EXPOSE 8080
ENTRYPOINT ["/bin/app"]