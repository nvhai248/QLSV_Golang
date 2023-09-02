package uploadstorage

import (
	"context"
	"studyGoApp/common"
)

func (store *sqlStore) DeleteImages(ctx context.Context, ids []int) error {
	db := store.db

	if _, err := db.Exec("DELETE FROM images WHERE id in (?),", ids); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
