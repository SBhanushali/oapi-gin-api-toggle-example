package spec

//go:generate oapi-codegen -generate embedded-spec -package spec -o spec.gen.go hello.v1.yaml
//go:generate oapi-codegen -generate gin-server,models,strict-server,models -package server -o ../server/server.gen.go hello.v1.yaml
//go:generate oapi-codegen -config oapi-client-cfg.yaml hello.v1.yaml
