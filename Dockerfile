FROM golang
LABEL MAINTAINER=wei1109942647@qq.com
COPY . /go/src/app
WORKDIR /go/src/app
EXPOSE 50001
RUN go build main.go
CMD ["/go/src/app/main"]