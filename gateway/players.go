package gateway

func GetPlayers() []byte {
	return Players()
}

func GetNames() []byte {
	return Names()
}

func GetKeyHolderStats(addr string) ([]byte, error) {
	url := fomoHost() + "/api/keyHolderStats?address=" + addr
	return HttpGet(url)
}

func getPlayers() ([]byte, error) {
	url := fomoHost() + "/api/players"
	return HttpGet(url)
}

func getNames() ([]byte, error) {
	url := fomoHost() + "/api/names"
	return HttpGet(url)
}
