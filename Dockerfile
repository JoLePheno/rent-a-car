FROM golang:latest

WORKDIR /home/stoik/app

RUN apt-get update

RUN apt install -y -qq --no-install-recommends \
        apt-transport-https \
        apt-utils \
        ca-certificates \
        curl

COPY . ./

RUN go mod vendor

RUN chmod 755 init.sh
RUN chmod 755 cmd/rentalctl/migrations_up.sh

EXPOSE 3000

CMD ["./init.sh"]