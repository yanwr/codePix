package cmd

import (
	"codePix/application/grpc"
	db "codePix/config"
	"codePix/env"
	"github.com/spf13/cobra"
	"os"
)

var portNumber int

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC Server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv(env.CURRENT_ENV))
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gRPC Server port")
}
