FROM golang:1.13
WORKDIR /go/src/app
RUN go get github.com/gin-gonic/gin \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/jinzhu/gorm
ADD ./app /go/src/app
CMD ["go", "run", "gin.go"]
