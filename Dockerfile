FROM golang:1.22-alpine as builder
WORKDIR /
COPY . ./
RUN go mod download


RUN go build -o /service-meurig


FROM alpine
COPY --from=builder /service-meurig .


EXPOSE 7777
CMD [ "/service-meurig" ]