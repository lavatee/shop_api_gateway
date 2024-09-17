package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	gateway "github.com/lavatee/shop_api_gateway"
	"github.com/lavatee/shop_api_gateway/internal/endpoint"
	"github.com/lavatee/shop_api_gateway/internal/repository"
	"github.com/lavatee/shop_api_gateway/internal/service"
	pb "github.com/lavatee/shop_protos/gen"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("config error", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("env error", err.Error())
	}
	db, err := repository.NewPostgresDB(viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.username"), os.Getenv("DB_PASSWORD"), viper.GetString("db.dbname"), viper.GetString("db.sslmode"))
	if err != nil {
		log.Fatalf("db error", err.Error())
	}
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	productsConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", viper.GetString("products.host"), viper.GetString("products.port")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("products connection error", err.Error())
	}
	productsClient := pb.NewProductsClient(productsConn)
	reviewsConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", viper.GetString("reviews.host"), viper.GetString("reviews.port")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("reviews connection error", err.Error())
	}
	reviewsClient := pb.NewReviewsClient(reviewsConn)
	savedConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", viper.GetString("saved.host"), viper.GetString("saved.port")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("saved connection error", err.Error())
	}
	savedClient := pb.NewSavedClient(savedConn)
	ordersConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", viper.GetString("orders.host"), viper.GetString("orders.port")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("orders connection error", err.Error())
	}
	ordersClient := pb.NewOrdersClient(ordersConn)
	end := &endpoint.Endpoint{
		Services:     svc,
		ProductsConn: productsClient,
		ReviewsConn:  reviewsClient,
		SavedConn:    savedClient,
		OrdersConn:   ordersClient,
	}
	handler := end.InitRoutes()
	server := &gateway.Server{}
	go func() {
		if err := server.Run(viper.GetString("port"), handler); err != nil {
			log.Fatalf("server error", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := productsConn.Close(); err != nil {
		log.Fatalf("close error", err.Error())
	}
	if err := savedConn.Close(); err != nil {
		log.Fatalf("close error", err.Error())
	}
	if err := reviewsConn.Close(); err != nil {
		log.Fatalf("close error", err.Error())
	}
	if err := ordersConn.Close(); err != nil {
		log.Fatalf("close error", err.Error())
	}
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("shutdown error", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
