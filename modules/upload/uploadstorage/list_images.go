package uploadstorage

import (
	"context"
	"studyGoApp/common"
)

func (store *sqlStore) ListImages(ctx context.Context) ([]common.Image, error) {
	db := store.db

	var images []common.Image

	if err := db.Select(&images, "SELECT * FROM images"); err != nil {
		return nil, common.ErrDB(err)
	}

	return images, nil
}
