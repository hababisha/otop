package repository

import "github.com/hababisha/otop/utils"

func CleanOldOtps() error {
	query := `
		DELETE FROM otps WHERE created_at < NOW() - INTERVAL '24 hours'
	`

	_, err := utils.DB.Exec(query)
	return err
}