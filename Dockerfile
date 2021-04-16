FROM golang:1.14

LABEL base.name="taxdooclientwithgo"

WORKDIR /app

COPY . .

RUN go build -o application .

EXPOSE 8080

ENTRYPOINT ["./application"]
