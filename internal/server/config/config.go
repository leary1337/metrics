package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	ServerAddr string
}

const DefaultServerAddr = `localhost:8080`

func NewConfig() *Config {
	config := &Config{}
	flag.StringVar(&config.ServerAddr, "a", DefaultServerAddr, "server address (default localhost:8080)")

	// Создаем кастомный флаг сет для обработки неизвестных флагов
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.CommandLine.SetOutput(os.Stderr)
	flag.CommandLine.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// Проверка на неизвестные флаги
	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	return config
}
