package repositories

type BlacklistRepository interface {
	AddBlacklist(token string) (bool, error)
	IsBlacklist(token string) (bool, error)
}
