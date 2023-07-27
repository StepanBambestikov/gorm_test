package main

func main() {
	DB, _ := NewTestPostgresOrderStorage()
	defer DB.CleanDBAfterTest()

	err := DB.innerDB.Create(&validOrders[0]).Error
	if err != nil {
		return
	}

	err = DB.innerDB.Create(&validRequisites[0]).Error
	if err != nil {
		return
	}

	err = DB.innerDB.Create(&validOrders[0]).Error
	if err != nil {
		return
	}
}
