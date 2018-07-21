package gateway

func GetPlayers() ([]byte, error) {
	url := fomoHost() + "/api/players"
	return HttpGet(url)
}

func GetNames() ([]byte, error) {
	url := fomoHost() + "/api/names"
	return HttpGet(url)
}

func GetKeyHolderStats(addr string) ([]byte, error) {
	url := fomoHost() + "/api/keyHolderStats?address=" + addr
	return HttpGet(url)
}
