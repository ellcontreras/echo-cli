package utils

import "log"

//CheckErr check if exists an error in the param err
func CheckErr(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}

//CheckNull check if a string is null or not
//func CheckNull(newApp string) bool {
//	return len(*newApp) > 0 && *newApp != ""
//}

//CheckNil check if non pointer of a string is null
func CheckNil(a string) bool {
	log.Println(a)
	return len(a) > 0 && a != ""
}
