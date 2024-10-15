package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// SetupLogger configura o log com o nível de log informado
func SetupLogger() *logrus.Logger {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	return log
}
