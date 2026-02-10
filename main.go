package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configurações do PABX
	pabxIP := "192.168.240.223"
	user := "freepbxuser"     // Usuário padrão (pode variar)
	pass := ""                // Senha (precisamos descobrir a do seu ambiente)
	dbname := "asteriskcdrdb"

	// DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, pabxIP, dbname)

	fmt.Printf("Conectando ao PABX em %s...
", pabxIP)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao abrir conexão:", err)
	}
	defer db.Close()

	// Testar conexão
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco do PABX (verifique IP/Porta/Credenciais):", err)
	}

	fmt.Println("Conectado com sucesso!")

	// Query de exemplo: últimas 5 chamadas
	rows, err := db.Query("SELECT calldate, src, dst, duration, billsec, disposition FROM cdr ORDER BY calldate DESC LIMIT 5")
	if err != nil {
		log.Fatal("Erro ao executar query:", err)
	}
	defer rows.Close()

	fmt.Println("
Últimas 5 chamadas:")
	fmt.Println("------------------------------------------------------------------")
	for rows.Next() {
		var calldate string
		var src, dst, disposition string
		var duration, billsec int

		err := rows.Scan(&calldate, &src, &dst, &duration, &billsec, &disposition)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Data: %s | De: %s -> Para: %s | Segundos: %d | Status: %s
", calldate, src, dst, billsec, disposition)
	}
}
