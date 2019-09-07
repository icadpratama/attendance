package claim

import . "github.com/icadpratama/attendance/internal/orm/models"

type FinancialClaim struct {
	BaseModelSoftDelete
	amount   *int
	reason   *string
	document string
}
