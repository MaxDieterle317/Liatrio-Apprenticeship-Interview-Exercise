FROM golang:1.27.1 
WORKDIR /liatrio_app 

COPY . .

RUN go build -o liatrio-app .

CMD ["./liatrio-app"]

