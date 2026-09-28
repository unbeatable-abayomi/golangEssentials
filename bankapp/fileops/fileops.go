package fileops


import (
   "fmt"
   "os"
   "strconv"
   "errors"
)





func GetFloatFromFile(fileName string) (float64, error){
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Falied To find  File")
	}
	valueText := string(data)
	value, err:= strconv.ParseFloat(valueText, 64)

		if err != nil {
		return 1000, errors.New("Falied To Parse Stored Value")
	}

	return value, nil

}

func WriteFloatFile(value float64, fileName string){
  balanceText := fmt.Sprint(value)
  os.WriteFile(fileName, []byte(balanceText), 0644)

}