package apps

import (
	"github.com/allegro/marathon-consul/tasks"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestParseApps(t *testing.T) {
	t.Parallel()

	appBlob, _ := ioutil.ReadFile("apps.json")

	expected := []*App{
		&App{
			HealthChecks: []HealthCheck{
				HealthCheck{
					Path:                   "/",
					PortIndex:              0,
					Protocol:               "HTTP",
					GracePeriodSeconds:     5,
					IntervalSeconds:        20,
					TimeoutSeconds:         20,
					MaxConsecutiveFailures: 3,
				},
			},
			ID: "/bridged-webapp",
			Tasks: []tasks.Task{
				tasks.Task{
					ID:                 "test.47de43bd-1a81-11e5-bdb6-e6cb6734eaf8",
					AppID:              "/test",
					Host:               "192.168.2.114",
					Ports:              []int{31315},
					HealthCheckResults: []tasks.HealthCheckResult{tasks.HealthCheckResult{Alive: true}},
				},
				tasks.Task{
					ID:    "test.4453212c-1a81-11e5-bdb6-e6cb6734eaf8",
					AppID: "/test",
					Host:  "192.168.2.114",
					Ports: []int{31797},
				},
			},
		},
	}
	apps, err := ParseApps(appBlob)
	assert.NoError(t, err)
	assert.Len(t, apps, 1)
	assert.Equal(t, expected, apps)
}

func TestParseApp(t *testing.T) {
	t.Parallel()

	appBlob, _ := ioutil.ReadFile("app.json")

	expected := &App{Labels: map[string]string{"consul": "true", "public": "tag"},
		HealthChecks: []HealthCheck{HealthCheck{Path: "/",
			PortIndex:              0,
			Protocol:               "HTTP",
			GracePeriodSeconds:     10,
			IntervalSeconds:        5,
			TimeoutSeconds:         10,
			MaxConsecutiveFailures: 3}},
		ID: "/myapp",
		Tasks: []tasks.Task{tasks.Task{
			ID:    "myapp.cc49ccc1-9812-11e5-a06e-56847afe9799",
			AppID: "/myapp",
			Host:  "10.141.141.10",
			Ports: []int{31678,
				31679,
				31680,
				31681},
			HealthCheckResults: []tasks.HealthCheckResult{tasks.HealthCheckResult{Alive: true}}},
			tasks.Task{
				ID:    "myapp.c8b449f0-9812-11e5-a06e-56847afe9799",
				AppID: "/myapp",
				Host:  "10.141.141.10",
				Ports: []int{31307,
					31308,
					31309,
					31310},
				HealthCheckResults: []tasks.HealthCheckResult{tasks.HealthCheckResult{Alive: true}}}}}

	app, err := ParseApp(appBlob)
	assert.NoError(t, err)
	assert.Equal(t, expected, app)
}