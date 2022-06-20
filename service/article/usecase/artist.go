package usecase

import (
	// "github.com/AndroX7/kumparan-assesment/lib/s3"
	"github.com/AndroX7/kumparan-assesment/lib/database_transaction"
	"github.com/AndroX7/kumparan-assesment/service/article"
)

type Usecase struct {
	transactionManager database_transaction.Client
	articleRepo        article.Repository
	// s3         s3.S3Client
}

func New(
	articleRepo article.Repository,
	transactionManager database_transaction.Client,
	// s3 s3.S3Client,
) article.Usecase {
	return &Usecase{
		transactionManager: transactionManager,
		articleRepo:        articleRepo,
		// s3:         s3,
	}
}
