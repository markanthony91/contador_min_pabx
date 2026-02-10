package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	PabxIP   string
	PabxPort string
	PabxUser string
	PabxPass string
	PabxDB   string
	ShopID   string
}

func loadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema.")
	}

	return Config{
		PabxIP:   getEnv("PABX_IP", "127.0.0.1"),
		PabxPort: getEnv("PABX_PORT", "3306"),
		PabxUser: getEnv("PABX_USER", "freepbxuser"),
		PabxPass: getEnv("PABX_PASS", ""),
		PabxDB:   getEnv("PABX_DB", "asteriskcdrdb"),
		ShopID:   getEnv("SHOP_ID", "UNKNOWN_SHOP"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	config := loadConfig()

	fmt.Printf("--- Contador de Minutos PABX [%s] ---\n", config.ShopID)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
		config.PabxUser, config.PabxPass, config.PabxIP, config.PabxPort, config.PabxDB)

	fmt.Printf("Conectando ao banco CDR em %s:%s...\n", config.PabxIP, config.PabxPort)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar ao banco do PABX: %v (Verifique IP/Credenciais)", err)
	}

	fmt.Println("✅ Conectado com sucesso ao MySQL!")

	// TODO: Implementar lógica de coleta incremental (não pegar duplicados)
	fetchLatestCalls(db)
}

func fetchLatestCalls(db *sql.DB) {
	query := "SELECT calldate, src, dst, duration, billsec, disposition FROM cdr ORDER BY calldate DESC LIMIT 5"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Erro ao executar query: %v", err)
	}
	defer rows.Close()

	fmt.Println("\nÚltimas 5 chamadas identificadas:")
	fmt.Println("------------------------------------------------------------------")
	for rows.Next() {
		var calldate string
		var src, dst, disposition string
		var duration, billsec int

		err := rows.Scan(&calldate, &src, &dst, &duration, &billsec, &disposition)
		if err != nil {
			log.Printf("Erro ao ler linha: %v", err)
			continue
		}
		fmt.Printf("Data: %s | De: %s -> Para: %s | Segundos: %d | Status: %s\n", 
			calldate, src, dst, billsec, disposition)
	}
}