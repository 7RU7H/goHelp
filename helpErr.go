import (
	"log"
	"runtime"
)

func ErrorHandling(err error, level int) error {
	switch level {
	case 0:
		log.Printf("Non fatal error: %v \n", err)
		break
	case 1:
		defer panicRecovery(&err)
		return err
	case 2:
		log.Fatalln(err)
		panic(err)
	}

	return nil
}

func panicRecovery(err *error) error {
	if r := recover(); r != nil {
		if _, ok := r.(runtime.Error); ok {
			log.Println("Panicing!!")
			panic(r)
		} else {
			*err = r.(error)
			return nil
		}
	}
	return nil
}
