FROM golang:1.19-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 5252

CMD [ "/app/main" ]