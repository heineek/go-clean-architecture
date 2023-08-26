//go:build proto
// +build proto

package proto

//go:generate protoc --proto_path=.      --go_opt=paths=source_relative      --go_out=. contact.proto
//go:generate protoc --proto_path=. --go-grpc_opt=paths=source_relative --go-grpc_out=. contact.proto
