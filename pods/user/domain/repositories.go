package domain

// copy from https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f

// FIXME 或许应该放在interfaces.go 里面？

//Repository repository interface
type UserRepository interface {
	// Find(id entity.ID) (*User, error)
	FindByUsername(username string) (*User, error)
	// FindByEmail(email string) (*User, error)
	// FindByChangePasswordHash(hash string) (*User, error)
	// FindByValidationHash(hash string) (*User, error)
	// FindAll() ([]*User, error)
	Update(user *User) error
	// Store(user *User) (entity.ID, error)
	// AddCompany(id entity.ID, company *Company) error
	// AddInvite(userID entity.ID, companyID entity.ID) error
}
