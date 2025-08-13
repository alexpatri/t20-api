package bootstrap

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"t20-api/internal/api"
	"t20-api/internal/database"
	"t20-api/internal/utils"
	"t20-api/internal/utils/config"

	"github.com/rs/zerolog/log"
)

type Flags struct {
	Env *string
}

// InitializeFlags trata as flags de execução
func InitializeFlags() Flags {
	var flags Flags

	flags.Env = flag.String("vars", "file", "Defines from where to load env vars: file or exported")

	flag.Parse()
	log.Info().Str("vars-from", *flags.Env).Msg("Flags initialized successfully.")

	return flags
}

// HandleMode trata o comportamento para o modo `file`
func HandleMode(env string) {
	if env == "file" {
		utils.ClearTerminal()
	}
}

// CreateConfigContext cria o contexto de configuração
func CreateConfigContext(flags Flags) *config.Config {
	log.Info().Msg("Creating configuration context...")

	cfg := config.LoadConfig(*flags.Env)

	log.Info().Msg("Configuration context created successfully.")

	return cfg
}

// ConnectToDatabase conecta ao banco de dados
func ConnectToDatabase(dbConfig *config.DatabaseConfig) {
	log.Info().Msg("Connecting to the database...")

	db, err := database.Connect(*dbConfig)
	if err != nil {
		log.Panic().Err(err).Msg("Error connecting to the database.")
	}

	log.Info().Msg("Successfully connected to the database.")

	database.Database = db

}

// CreateServer cria e inicializa o servidor
func CreateServer(cfg *config.Config) {
	log.Info().Msg("Starting the API service...")

	server := api.NewServer(cfg)

	// Canal para receber sinal de interrupção
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Iniciar servidor em uma goroutine
	go func() {
		if err := server.Start(); err != nil {
			log.Fatal().Err(err).Msg("Failed to start server.")
		}
	}()

	log.Info().Str("porta", cfg.Server.Port).Msg("Server started.")

	// Aguardar sinal para encerrar
	<-quit
	log.Info().Msg("Shutting down server...")

	err := database.Disconnect(database.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to disconnect from database.")
	}

	log.Info().Msg("Disconnected from database.")
}
