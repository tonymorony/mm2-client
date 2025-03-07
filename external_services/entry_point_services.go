package external_services

import (
	"github.com/kpango/glg"
	"mm2_client/constants"
)

func LaunchPriceServices() {
	glg.Info("Starting binance websocket service")
	go StartBinanceWebsocketService()
	glg.Info("Starting coingecko price service")
	go StartCoingeckoService()
	glg.Info("Starting coinpaprika price service")
	go StartCoinpaprikaService()
	constants.GPricesServicesRunning = true
}

func LaunchMessagesService(kind string, target string) {
	go StartNotifierMessagesService(kind, target)
}

func LaunchServices(kind string, target string) {
	glg.Info("Launching price services")
	LaunchPriceServices()
	glg.Info("Launching extra services")
	LaunchMessagesService(kind, target)
}
