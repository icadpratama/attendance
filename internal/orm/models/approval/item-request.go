package approval

import . "github.com/icadpratama/attendance/internal/orm/models"

type ItemRequest struct {
	BaseModelSoftDelete
	item     *string
	quantity *int
	details  *string
}
