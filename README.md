### Informações do Projeto

Esse projeto é um exporter do prometheus para coletar informações de total de memória e memória livre

- Endpoint: /metrics



### Informações de Execução Local (Sem Dockerfile)

Iniciando o módulo

```
go mod init exporter
```

Baixando dependências

```
go mod tidy

```

Executando binário

```
./exporter
```

### Informações de Execução Local (Com Dockerfile)

```
docker build -t exporter:1.0.0 .
```

```
docker run -d --name exporter -p 7788:7788  exporter
```