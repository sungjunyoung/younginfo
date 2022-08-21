##
## Build
##
FROM golang:alpine3.15 AS build

WORKDIR /app
COPY . ./
RUN go build -o /reviews

##
## Deploy
##
FROM alpine:3.15.4
RUN apk add bash
COPY --from=build /reviews /reviews
RUN chmod +x /reviews
CMD ["/reviews"]