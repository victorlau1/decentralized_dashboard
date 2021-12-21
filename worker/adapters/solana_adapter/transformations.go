package solanaadapter

import (
	sw "github.com/victorlau1/solanaclient"
	"github.com/victorlau1/worker/models"
)

// Transformations are the transformations specified to alter the source data into the
// ElasticSearch required data structure.
type Transformations interface {
	HandleClientTransformation(interface{}) (models.ClientDecentralization, error)
}

type solanaBeachTransformations struct {
}

type solanaBeachNonValidatorModel struct {
	*sw.InlineResponse20015
}

func (sb *solanaBeachTransformations) HandleClientTransformation(solanaResponse interface{}) (models.ClientDecentralization, error) {
	nm := models.ClientDecentralization{}
	// res, ok := solanaResponse.(solanaBeachNonValidatorModel); ok {
	// 	res.Location.
	// }
	// return nm, err
}
