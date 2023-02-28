package server

func Init() error {
	router, err := NewRouter()
	if err != nil {
		return err
	}
	router.Run()
	return nil
}
