package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"architecture_go/pkg/store/postgres"
	deliveryGrpc "architecture_go/services/contact/internal/delivery/grpc"
	deliveryHttp "architecture_go/services/contact/internal/delivery/http"
	// repositoryContact "architecture_go/services/contact/internal/repository/contact/postgres"
	// repositoryGroup "architecture_go/services/contact/internal/repository/group/postgres"
	repositoryStorage "architecture_go/services/contact/internal/repository/storage/postgres"
	useCaseContact "architecture_go/services/contact/internal/useCase/contact"
	useCaseGroup "architecture_go/services/contact/internal/useCase/group"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	// repoContact, err:= repositoryContact.New(conn.Pool, repositoryContact.Options{})
	// if err != nil {
	// 	panic(err)
	// }
	// repoGroup, err:= repositoryGroup.New(conn.Pool, repoContact, repositoryGroup.Options{})
	// if err != nil {
	// 	panic(err)
	// }

	repoStorage, err := repositoryStorage.New(conn.Pool, repositoryStorage.Options{})
	if err != nil {
		panic(err)
	}
	var (
		ucContact = useCaseContact.New(repoStorage, useCaseContact.Options{})
		// ucGroup      = useCaseGroup.New(repoGroup, useCaseGroup.Options{})
		ucGroup      = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		_            = deliveryGrpc.New(ucContact, ucGroup, deliveryGrpc.Options{})
		listenerHttp = deliveryHttp.New(ucContact, ucGroup, deliveryHttp.Options{})
	)

	go func() {
		fmt.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))
		if err = listenerHttp.Run(); err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

}
