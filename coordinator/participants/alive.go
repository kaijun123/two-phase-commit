package participants

import "fmt"

type AliveMap struct {
	mapping map[string]bool
}

func CreateAliveMap(ipArray *[]string) *AliveMap {
	aliveMap := make(map[string]bool)

	for _, value := range *ipArray {
		fmt.Println("value", value)
		aliveMap[value] = true
	}

	return &AliveMap{
		mapping: aliveMap,
	}
}

func (a *AliveMap) UpdateMap(ip string) error {
	a.mapping[ip] = true
	return nil
}
