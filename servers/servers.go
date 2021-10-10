package servers

var (
	ExchangeSvr *Exchange
	UserSvr    *User
)

func InitServers(){
	ExchangeSvr=NewExchange()
	UserSvr =NewUser()

}
