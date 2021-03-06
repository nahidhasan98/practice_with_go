package main

import "fmt"

func main() {
	// //printing given data
	// fmt.Println("Given data:")
	// for key, val := range givenData {
	// 	fmt.Println(key, val)
	// }

	result := listAccount(givenData)

	//printing the result
	fmt.Println("Result data:")
	for key, val := range result {
		fmt.Println(key, val)
	}
}

func listAccount(data []map[string]interface{}) map[string][]interface{} {
	result := make(map[string][]interface{}) //defining a result variable

	for key := range data {
		mapKey := data[key]["account_type"].(string) //getting the key for given map data

		result[mapKey] = append(result[mapKey], data[key]["name"]) //setting the name value(got from given data) to result variable
	}

	return result
}

//this is given data
var givenData = []map[string]interface{}{
	0:  {"account_type": "Asset", "aid": "achead::7", "name": "Non Current Asset", "parent_id": "achead::2"},
	1:  {"account_type": "Asset", "aid": "achead::8", "name": "Machinery", "parent_id": "achead::7"},
	2:  {"account_type": "Asset", "aid": "achead::9", "name": "Current Asset", "parent_id": "achead::2"},
	3:  {"account_type": "Asset", "aid": "achead::10", "name": "Account Receivable", "parent_id": "achead::9"},
	4:  {"account_type": "Asset", "aid": "achead::11", "name": "Advance Deposit", "parent_id": "achead::9"},
	5:  {"account_type": "Asset", "aid": "achead::12", "name": "Cash at Bank", "parent_id": "achead::9"},
	6:  {"account_type": "Asset", "aid": "achead::13", "name": "Cash in Hand", "parent_id": "achead::9"},
	7:  {"account_type": "Asset", "aid": "achead::14", "name": "Inventory", "parent_id": "achead::9"},
	8:  {"account_type": "Liability", "aid": "achead::15", "name": "Longterm Liability", "parent_id": "achead::3"},
	9:  {"account_type": "Liability", "aid": "achead::16", "name": "Short-term Liability", "parent_id": "achead::3"},
	10: {"account_type": "Liability", "aid": "achead::17", "name": "Bank Loan", "parent_id": "achead::15"},
	11: {"account_type": "Liability", "aid": "achead::18", "name": "Current Liability", "parent_id": "achead::3"},
	12: {"account_type": "Liability", "aid": "achead::19", "name": "Account Payable", "parent_id": "achead::18"},
	13: {"account_type": "Liability", "aid": "achead::20", "name": "Duties & Taxes", "parent_id": "achead::18"},
	14: {"account_type": "Revenue", "aid": "achead::21", "name": "Sales Account", "parent_id": "achead::5"},
	15: {"account_type": "Revenue", "aid": "achead::22", "name": "Indirect Income", "parent_id": "achead::5"},
	16: {"account_type": "Equity", "aid": "achead::23", "name": "Capital Account", "parent_id": "achead::4"},
	17: {"account_type": "Equity", "aid": "achead::24", "name": "Drawings Account", "parent_id": "achead::23"},
	18: {"account_type": "Equity", "aid": "achead::25", "name": "Retained Earnings", "parent_id": "achead::4"},
	19: {"account_type": "Equity", "aid": "achead::26", "name": "Opening Balance", "parent_id": "achead::4"},
	20: {"account_type": "Expense", "aid": "achead::27", "name": "Direct Expense", "parent_id": "achead::6"},
	21: {"account_type": "Expense", "aid": "achead::28", "name": "Cost of Revenue", "parent_id": "achead::27"},
	22: {"account_type": "Expense", "aid": "achead::29", "name": "COGS", "parent_id": "achead::27"},
	23: {"account_type": "Expense", "aid": "achead::30", "name": "Indirect Expense", "parent_id": "achead::6"},
	24: {"account_type": "Expense", "aid": "achead::31", "name": "Administrative Expense", "parent_id": "achead::30"},
	25: {"account_type": "Expense", "aid": "achead::32", "name": "Salary & Allowance", "parent_id": "achead::31"},
	26: {"account_type": "Expense", "aid": "achead::33", "name": "Cash Discount", "parent_id": "achead::30"},
}

// //result data should be like this
// var returnData = map[string][]interface{}{
// 	"Asset":     {"Non Current Asset", "Machinery", "Current Asset", "Account Receivable", "Advance Deposit", "Cash at Bank", "Cash in Hand", "Inventory"},
// 	"Liability": {"Longterm Liability", "Short-term Liability", "Bank Loan", "Current Liability", "Account Payable", "Duties & Taxes"},
// 	"Revenue":   {"Sales Account", "Indirect Income"},
// 	"Equity":    {"Capital Account", "Drawings Account", "Retained Earnings", "Opening Balance"},
// 	"Expense":   {"Direct Expense", "Cost of Revenue", "COGS", "Indirect Expense", "Administrative Expense", "Salary & Allowance", "Cash Discount"},
// }
