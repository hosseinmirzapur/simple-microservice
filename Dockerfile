FROM golang:1.21-alpine

WORKDIR /app

COPY . ./

RUN go build -o price .


EXPOSE 3000

CMD ["./price"]