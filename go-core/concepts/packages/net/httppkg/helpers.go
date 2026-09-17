package main


func nextEmployeeID(emps []Employee) int {
	maxID := 0
	for _, emp := range emps {
		maxID = max(maxID, emp.ID)
	}
	newID := maxID + 1
	return newID
}
