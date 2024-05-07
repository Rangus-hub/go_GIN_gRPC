package repository


type Repository interface {
    
    // FetchData() ([]Data, error)
    // StoreData(data Data) error
    Deletedata(name string, id string) string
    AddData(name, id string) string
}
