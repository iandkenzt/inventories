package restapi

// ValidateAppSecretKey check given key and compare with app key from env
func ValidateAppSecretKey(myKey string) bool {
	apiSecretKey := Conf.AppSecretKey
	if myKey == apiSecretKey {
		// fmt.Println("Your token is valid.  I like your style.")
		return true
	}

	// fmt.Println("This token is terrible!  I cannot accept this.")
	return false
}
