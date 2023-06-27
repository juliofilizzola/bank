FROM golang:1.20

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"


RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

COPY . /go/src

RUN go mod download

RUN go build bank

CMD ["tail", "-f", "/dev/null"]