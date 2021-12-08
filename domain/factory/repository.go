package factory

import "github.com/mahgtavares/fullstack-fullcicle-immersion/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
