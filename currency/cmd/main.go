package main

import (
	"context"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"currency/internal/controller/grpc"
	_http "currency/internal/controller/http"
	"currency/internal/infrastructure/currency/binance"
	"currency/internal/infrastructure/currency/coingecko"
	"currency/internal/infrastructure/currency/currencyapi"
	"currency/internal/infrastructure/events/kafka"
	"currency/internal/services/currency"
	echoserver "currency/pkg/echo"
	"currency/pkg/grpc/server"
	"currency/pkg/http/client"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("starting application...")

	fx.New(CreateApp()).Run()

	logrus.Info("application stopped.")
}

func CreateApp() fx.Option { //nolint: ireturn // ok
	return fx.Options(
		fx.Provide(
			NewHTTPServices,
			NewGRPCServices,
			NewBinanceConfig,
			NewEchoConfig,
			NewCurrencyapiConfig,
			NewCoingeckoConfig,
			NewGrpcConfig,
			createChan,
			func() *http.Client {
				return client.New(client.WithLogger(logrus.StandardLogger()))
			},
			func(cfg *currencyapi.Config, cli *http.Client) *currencyapi.CurrencyAPI {
				return currencyapi.NewCurrencyAPI(
					cfg,
					currencyapi.WithClient(cli),
				)
			},
			func(cfg *binance.Config, cli *http.Client) *binance.BinanceAPI {
				return binance.NewBinanceApi(
					cfg,
					binance.WithClient(cli),
				)
			},
			func(cfg *coingecko.Config, cli *http.Client) *coingecko.CoingeckoAPI {
				return coingecko.NewCoingeckoApi(
					cfg,
					coingecko.WithClient(cli),
				)
			},
			fx.Annotate(
				func(c *currencyapi.CurrencyAPI) *currencyapi.CurrencyAPI {
					return c
				},
				fx.As(new(currency.ICurrencyAPI)),
			),
			kafka.CreateKafka,
			currency.NewCurrencyService,
			server.NewServer,
			echoserver.New,
		),
		fx.Invoke(
			startErrorHandling,
			registerCryptoChain,
			registerHttpHandlers,
			registerGRPCHandlers,
			startHTTPServer,
			startGRPCServer,
			makeRequest,
		),
	)
}

func createChan() chan error {
	return make(chan error)
}

func makeRequest(k *kafka.Kafka) {
	go func(k *kafka.Kafka) {
		for {
			time.Sleep(time.Second * 4)
			logrus.Info("kafka send error:", k.Publish())
		}
	}(k)
}

func startHTTPServer(srv *echoserver.Server, ls fx.Lifecycle, errChan chan error) {
	ls.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				go srv.Start(errChan)
				return nil
			},
			OnStop: func(_ context.Context) error {
				return srv.Stop()
			},
		},
	)
}

func startGRPCServer(srv *server.Server, ls fx.Lifecycle, errChan chan error) {
	ls.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				srv.Run(errChan)
				return nil
			},
			OnStop: func(_ context.Context) error {
				srv.Stop()
				return nil
			},
		},
	)
}

func startErrorHandling(shutdowner fx.Shutdowner, errChan chan error) {
	go func(ch chan error) {
		err := <-ch
		logrus.Error(err)
		shutdowner.Shutdown() //nolint: errcheck// no need here to check
	}(errChan)
}

func registerCryptoChain(
	cu *currencyapi.CurrencyAPI,
	b *binance.BinanceAPI,
	co *coingecko.CoingeckoAPI,
) {
	cu.SetNext(b)
	b.SetNext(co)
}

func registerHttpHandlers(srv *_http.Services, e *echoserver.Server) {
	_http.RegisterHandlers(e.GetEchoServer(), srv)
}

func registerGRPCHandlers(srv *grpc.Services, s *server.Server) {
	grpc.RegisterHandlers(s.Server, srv)
}

func NewHTTPServices(
	c *currency.Service,
) *_http.Services {
	return &_http.Services{
		CurrencyService: c,
	}
}

func NewGRPCServices(
	c *currency.Service,
) *grpc.Services {
	return &grpc.Services{
		CurrencyService: c,
	}
}
