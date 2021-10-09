package servers

var (
	ExchangeSvr *Exchange
)

func InitServers(){
	ExchangeSvr=NewExchange()

}
