package paginator

import (
	"testing"

	"github.com/HongJaison/go-admin4/modules/config"
	"github.com/HongJaison/go-admin4/plugins/admin/modules/parameter"
	_ "github.com/HongJaison/themes4/sword"
)

func TestGet(t *testing.T) {
	config.Initialize(&config.Config{Theme: "sword"})
	Get(Config{
		Size:         105,
		Param:        parameter.BaseParam().SetPage("7"),
		PageSizeList: []string{"10", "20", "50", "100"},
	})
}
