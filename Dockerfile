FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app
COPY . .

RUN go get -u -d github.com/golang-migrate/migrate/v4

RUN cd ..
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

#RUN migrate -path sql/migrations -database "postgresql://postgres:123456@db:5432/bank_server?sslmode=disable" -verbose up

CMD [ "top" ]

#gerar binario do go, nome do binario é server
#-ldflags="-w -s" remover informacoes de profile e debug, retirar em producao
#CGO_ENABLED=0 remover recursos do c no go para producao, caso nao haver dependencia de bibliotecas em c
#imagem scratch nao possui CGO
#RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd

#FROM scratch
#COPY --from=builder /app/server .
#CMD ["./server"]