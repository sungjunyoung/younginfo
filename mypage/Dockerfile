##
## Build
##
FROM golang:alpine3.15 AS build

WORKDIR /app
COPY . ./
RUN go build -o /mypage

##
## Deploy
##
FROM alpine:3.15.4
RUN apk add bash
COPY --from=build /mypage /mypage
RUN chmod +x /mypage
CMD ["/mypage"]