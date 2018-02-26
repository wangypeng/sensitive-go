package conf

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var ConfigMap = map[string]string{}

var configPath = "./server.conf"

func InitConf() {
	bytes, err := ioutil.ReadFile(configPath)
	
    if err != nil {
        fmt.Println("ReadFile: ", err.Error())
    }

    if err := json.Unmarshal(bytes, &ConfigMap); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
	}
}