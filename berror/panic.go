package berror

import "log"

// import "log"
//
//	defer func() {
//		log.Println("done")
//		// Println executes normally even if there is a panic
//		if err := recover(); err != nil {
//			log.Printf("run time panic: %v", err)
//		}
//	}()

func PanicRecover() {
	if err := recover(); err != nil {
		log.Printf("runtime panic:%v", err)
	}
}
