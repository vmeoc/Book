FROM golang

ENV db_type="mysql"
ENV db_username="root"
ENV db_password="PASSWORD"
ENV db_ip="172.18.12.219"
ENV db_name="Test"
ENV tito_ip="172.18.12.219:1234"

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/go-sql-driver/mysql
RUN go build -o binary ./src/booking.go ./src/check.go ./src/data.go ./src/display.go ./src/get.go ./src/main.go ./src/picture.go
WORKDIR /app/src
EXPOSE 8080
ENTRYPOINT /app/binary
