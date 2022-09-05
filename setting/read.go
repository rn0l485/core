package setting

import (
	"os"
	"io/ioutil"
	"encoding/json"  
)

func ReadJsonConfig(path string, data interface{}) error {
	jsonFile, err := os.Open(path)  
	if err != nil {  
		return err
	}  
	defer jsonFile.Close()  

	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, data); err != nil {
		return err
	}

	return nil
}