package service

import (
	"fmt"
	"os"
)

const MODE_PROD = "PROD"
const MODE_TEST = "TEST"

var AppMode string = func() string {
	mode := os.Getenv("REFLOREST_APP_MODE")
	fmt.Println("from env:" + mode)
	switch mode {
	case MODE_PROD:
		return MODE_PROD
	case MODE_TEST:
		return MODE_TEST
	default:
		return MODE_PROD
	}
}()

// Set the Application mode at runtime.
func SetAppMode(mode string) error {
	switch mode {
	case MODE_PROD:
		AppMode = MODE_PROD
		return nil
	case MODE_TEST:
		AppMode = MODE_TEST
		return nil
	default:
		return fmt.Errorf("Not a Valid APP Mode Supplied")
	}
}
