func checkExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		log.Fatal(err)
		return false
	}
	if os.IsNotExist(err) {
		log.Fatal("Not Exists")
		return false
	}
	return true
}
