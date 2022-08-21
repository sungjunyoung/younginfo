##
## Build
##
FROM golang:alpine3.15 AS build

WORKDIR /app
COPY . ./
RUN go build -o /ratings

##
## Deploy
##
FROM alpine:3.15.4
RUN apk add bash
COPY --from=build /ratings /ratings
RUN chmod +x /ratings
CMD ["/ratings"]