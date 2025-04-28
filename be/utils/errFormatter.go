package utils

func HandleErr(msg string, err error) {
	if err != nil {
		panic(msg + err.Error())
	}
}
