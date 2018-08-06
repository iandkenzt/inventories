package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger ...
var Logger = logrus.New()

// InitLogger ...
func InitLogger() {
	Logger.Out = os.Stdout
}
