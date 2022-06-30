package helpers

func PanicOnErr(err error) {
    if err != nil {
        panic(err.Error())
    }
}