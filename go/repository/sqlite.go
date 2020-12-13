package repository

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/ThiagoRodriguesdeSantana/desafio_conductor/go-server-server/go/model"
	_ "github.com/mattn/go-sqlite3"
	uuidPkg "github.com/nu7hatch/gouuid"
)

//SqliteDb conection
type SqliteDb struct {
	dbInstanse *sql.DB
}

//CloseDb conection DB
func (db *SqliteDb) CloseDb() {
	db.dbInstanse.Close()
}

//InitDB initialize db
func InitDB() *SqliteDb {

	fileinfo, _ := os.Stat("sqlite-database.db")

	if fileinfo == nil {
		log.Println("Creating sqlite-database.db...")
		file, err := os.Create("sqlite-database.db")
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("sqlite-database.db created")
	}

	instanse, err := sql.Open("sqlite3", "./sqlite-database.db")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	db := SqliteDb{
		dbInstanse: instanse,
	}

	if fileinfo == nil {
		db.createTableAccount()
		db.createTableTransactions()
		// INSERT RECORDS
		db.loadFakeRegister()
	}

	return &db
}

func (db *SqliteDb) loadFakeRegister() {

	accountsID := make([]string, 0)

	status := []string{"ativo", "inativo"}

	for i := 0; i < 4; i++ {

		id, _ := uuidPkg.NewV4()

		idString := id.String()

		accountsID = append(accountsID, idString)

		account := model.Account{
			Id:     idString,
			Status: status[rand.Intn(1)],
		}

		db.insertAccount(account)

	}

	descriptions := []string{"Netflix", "Apple Store", "Amazon"}
	values := []float64{199.00, 10.00, 150.50, 45.90, 32.77}

	for i := 0; i < 10; i++ {

		id, _ := uuidPkg.NewV4()

		idString := id.String()

		accountID := accountsID[rand.Intn(3)]

		transaction := model.Transaction{
			Id:          idString,
			AccountId:   accountID,
			Description: descriptions[rand.Intn(2)],
			Value:       values[rand.Intn(4)],
		}

		db.insertTransaction(transaction)

	}

}

func (db *SqliteDb) createTableAccount() {
	createTableAccount := `CREATE TABLE accounts (
		"id" TEXT,		
		"status" TEXT		
	  );`

	log.Println("Create account table...")
	statement, err := db.dbInstanse.Prepare(createTableAccount)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("account table created")
}

func (db *SqliteDb) createTableTransactions() {
	createTableTransaction := `CREATE TABLE account_transaction (
		"id" TEXT,		
		"account_id" TEXT,
		"description" TEXT,
		"value" TEXT		
	  );`

	log.Println("Create account_transaction table...")
	statement, err := db.dbInstanse.Prepare(createTableTransaction)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("account account_transaction created")
}

func (db *SqliteDb) insertAccount(account model.Account) {
	log.Println("Inserting account record ...")
	insertStudentSQL := `INSERT INTO accounts(id, status) VALUES (?, ?)`
	statement, err := db.dbInstanse.Prepare(insertStudentSQL)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(account.Id, account.Status)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (db *SqliteDb) insertTransaction(transaction model.Transaction) {
	log.Println("Inserting account record ...")
	insertStudentSQL := `INSERT INTO account_transaction(id, account_id, description, value) VALUES (?, ?, ?, ?)`
	statement, err := db.dbInstanse.Prepare(insertStudentSQL)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(transaction.Id, transaction.AccountId, transaction.Description, transaction.Value)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

//FindAllAccounts get all accounts
func (db *SqliteDb) FindAllAccounts() ([]model.Account, error) {

	row, err := db.dbInstanse.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	accounts := make([]model.Account, 0)

	for row.Next() {
		var id string
		var status string

		row.Scan(&id, &status)

		account := model.Account{
			Id:     id,
			Status: status,
		}

		accounts = append(accounts, account)

		log.Println("Account: ", id)

	}

	return accounts, nil

}

//FindAccountByID find account by id
func (db *SqliteDb) FindAccountByID(id string) (*model.Account, error) {

	stm, err := db.dbInstanse.Prepare("Select * from accounts where id = ?")

	if err != nil {
		return nil, err
	}

	resp, _ := stm.Query(id)

	account := model.Account{}

	for resp.Next() {
		var id string
		var status string

		resp.Scan(&id, &status)

		account.Id = id
		account.Status = status

	}

	return &account, err

}

//FindTransactionByAccountID find transaction account by id
func (db *SqliteDb) FindTransactionByAccountID(id string) ([]model.Transaction, error) {

	stm, err := db.dbInstanse.Prepare("Select * from account_transaction where account_id = ?")

	if err != nil {
		return nil, err
	}

	resp, _ := stm.Query(id)

	transactions := make([]model.Transaction, 0)

	for resp.Next() {
		var id string
		var accountID string
		var description string
		var valueSt string

		resp.Scan(&id, &accountID, &description, &valueSt)

		value, err := strconv.ParseFloat(valueSt, 32)

		if err != nil {
			fmt.Println("invalid conversion", err)
			value = 0.0
		}

		transaction := model.Transaction{
			Id:          id,
			AccountId:   accountID,
			Description: description,
			Value:       value,
		}

		transactions = append(transactions, transaction)
	}

	return transactions, err

}
