package utils

import "log"

//CheckErr check if exists an error in the param err
func CheckErr(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}

//CheckNull check if a string is null or not
<<<<<<< HEAD
func CheckNull(newApp *string) bool {
	return len(*newApp) > 0 && *newApp != ""
}

//CheckNil check if non pointer of a string is null
func CheckNil(a string) bool {
=======
//func CheckNull(newApp string) bool {
//	return len(*newApp) > 0 && *newApp != ""
//}

//CheckNil check if non pointer of a string is null
func CheckNil(a string) bool {
	log.Println(a)
>>>>>>> 15d7c53f62fd6f82c9fd105b3d417c0e95038f43
	return len(a) > 0 && a != ""
}
