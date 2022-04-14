func checkExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		log.Fatal(err)
		return false
	}
	if os.IsNotExist(err) {
		log.Fatal("Path does not Exists")
		return false
	}
	return true
}


func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}
