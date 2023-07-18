func checkFileExists(path string) (bool error) {
        _, err := os.Stat(path)
        if err == nil {
                log.Fatal(err)
                return false, err
        }
        if os.IsNotExist(err) {
                log.Fatal("File path does not exist")
                return false, err
        }
        return true, nil
}


func isDirectory(path string) (bool, error) {
        fileInfo, err := os.Stat(path)
        if err != nil {
                return false, err
        }

        return fileInfo.IsDir(), err
}
