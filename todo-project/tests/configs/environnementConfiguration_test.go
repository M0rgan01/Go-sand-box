package configs

import (
	"github.com/morgan/Go-sand-box/todo-project/configs"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvConf(t *testing.T) {
	t.Run("Get Environnement with env", getEnvironnementWithEnv)
	t.Run("Get Environnement without env returns default Environnement", getEnvironnementWithoutEnvReturnsDefaultEnvironnement)
}

func getEnvironnementWithEnv(t *testing.T) {
	//given
	environnementValue := "production"
	_ = os.Setenv(configs.Environnement, environnementValue)

	//when
	environnement := configs.GetEnvironnement()

	//then
	assert.Equal(t, environnementValue, environnement)

	os.Clearenv()
}

func getEnvironnementWithoutEnvReturnsDefaultEnvironnement(t *testing.T) {
	//given
	environnementValue := "development"

	//when
	environnement := configs.GetEnvironnement()

	//then
	assert.Equal(t, environnementValue, environnement)
}
