package uploadstorage

import (
	"context"
	"studyGoApp/common"
)

func (store *sqlStore) CreateImage(ctx context.Context, data *common.Image) error {
	db := store.db

	if _, err := db.Exec("INSERT INTO images (url, width, height, cloud_name, extension) VALUES (?, ?, ?, ?, ?)",
		data.Url, data.Width, data.Height, data.CloudName, data.Extension); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
