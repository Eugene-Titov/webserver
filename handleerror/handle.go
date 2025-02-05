package handleerror

import "fmt"

func HandleErrorWithPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleErrorToOutput(err error) {
	fmt.Println(err)
}
