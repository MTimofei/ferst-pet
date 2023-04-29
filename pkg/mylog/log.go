package mylog


func MyCustomLogger(prefix string, flags int) *log.Logger {
    file, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    return log.New(file, prefix, flags)
}

myLogger := MyCustomLogger("MYAPP: ", log.LstdFlags)
myLogger.Println("Some log message")
