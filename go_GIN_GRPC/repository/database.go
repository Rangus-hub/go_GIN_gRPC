package repository

import (
	"database/sql"
)


type DatabaseRepository struct {
   
}


func NewDatabaseRepository() *DatabaseRepository {
    return &DatabaseRepository{}    
}

func connectToDB(){
	db, err := sql.Open("godror", ConnectionString)
    if err != nil {
        panic(err)
    }
    defer db.Close()
}

// func (r *DatabaseRepository) FetchData() ([]Data, error) {
   
// }

func (r *DatabaseRepository) Deletedata(name string, id string) string {
   return "data deleted for" + name + "and id " + id
} 
func (r *DatabaseRepository) AddData(name, id string) string {
    return "data added for" + name + "and id " + id
} 
