##
## Build
##
FROM golang:alpine3.15 AS build

WORKDIR /app
COPY . ./
RUN go build -o /details

##
## Deploy
##
FROM alpine:3.15.4
RUN apk add bash
COPY --from=build /details /details
RUN chmod +x /details
CMD ["/details"]