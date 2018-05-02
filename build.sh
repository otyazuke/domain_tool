#!/bin/bash
echo mainをビルドします。。。
go build -o ./lib/main main.go
echo synonymsをビルドします。。。
go build -o ./lib/synonyms synonyms.go
echo availableをビルドします。。。
go build -o ./lib/available available.go
echo sprinkleをビルドします。。。
go build -o ./lib/sprinkle sprinkle.go
echo coolifyをビルドします。。。
go build -o ./lib/coolify coolify.go
echo domainifyをビルドします。。。
go build -o ./lib/domainify domainify.go
