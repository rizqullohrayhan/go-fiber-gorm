package seeders

func SeederInit() error {
	if err := SeedRoles(); err != nil {
		return err
	}
	if err := SeedUsers(); err != nil {
		return err
	}
	if err := SeedStatusBorrowHistories(); err != nil {
		return err
	}
	return nil
}