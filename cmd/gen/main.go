package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/suremarc/go-acr"
)

func main() {
	fi, err := os.Open(os.Args[1])
	if err != nil {
		logrus.WithError(err).Fatal("open file")
	}

	db, err := acr.NewDB(fi)
	if err != nil {
		logrus.WithError(err).Fatal("load db")
	}

	logrus.Info(
		db.Table,
		db.Field,
		db.Template,
		db.Xref,
	)
}
