FROM golang:1.15 as build

WORKDIR /app
COPY ./app /app/
RUN go build -o app

FROM alpine as runtime
COPY --from=build /app/app /
CMD ./app