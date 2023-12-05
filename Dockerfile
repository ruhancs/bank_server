FROM golang:latest as builder
WORKDIR /app
COPY . .

#gerar binario do go, nome do binario é server
#-ldflags="-w -s" remover informacoes de profile e debug, retirar em producao
#CGO_ENABLED=0 remover recursos do c no go para producao, caso nao haver dependencia de bibliotecas em c
#imagem scratch nao possui CGO
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]