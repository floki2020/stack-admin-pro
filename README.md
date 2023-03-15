# go-manage
This is a base project 

use Ent orm

mkdir go-manage

cd go-manage

go mod init go-manage

goctl api new core

go mod tidy

mkdir pkg/ent

cd pkg

go run -mod=mod entgo.io/ent/cmd/ent new User Company Personal

go generate ./ent