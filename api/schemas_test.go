package api

import "testing"

func jccFixtureTrainee() JwtClaims {
	return JwtClaims{
		RealmAccess: map[string][]string{
			"roles": {
				"samo_trainee_app",
				"offline_access",
				"default-roles-trainer-helper",
				"uma_authorization",
			},
		},
	}
}

func jccFixtureTrainer() JwtClaims {
	return JwtClaims{
		RealmAccess: map[string][]string{
			"roles": {
				"samo_trainer_app",
				"samo_trainee_app",
				"offline_access",
				"default-roles-trainer-helper",
				"uma_authorization",
			},
		},
	}
}

func jccFixtureTrainerLater() JwtClaims {
	return JwtClaims{
		RealmAccess: map[string][]string{
			"roles": {
				"samo_trainee_app",
				"samo_trainer_app",
				"offline_access",
				"default-roles-trainer-helper",
				"uma_authorization",
			},
		},
	}
}

func TestAppRole(t *testing.T) {
	var test = []struct {
		name          string
		jcc           JwtClaims
		wantIsTrainer bool
		wantRole      string
	}{
		{"trainer", jccFixtureTrainer(), true, "samo_trainer_app"},
		{"trainee", jccFixtureTrainee(), false, "samo_trainee_app"},
		{"empty", JwtClaims{}, false, ""},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			gotRole, gotIsTrainer := tt.jcc.AppRole()

			if gotRole != tt.wantRole || gotIsTrainer != tt.wantIsTrainer {
				t.Errorf("got %s | %t , want %s | %t", gotRole, gotIsTrainer, tt.wantRole, tt.wantIsTrainer)
			}
		})
	}
}

func TestAppTraineeRole(t *testing.T) {
	var test = []struct {
		name     string
		jcc      JwtClaims
		wantRole string
	}{
		{"trainer", jccFixtureTrainer(), "samo_trainee_app"},
		{"trainerLater", jccFixtureTrainerLater(), "samo_trainee_app"},
		{"trainee", jccFixtureTrainee(), "samo_trainee_app"},
		{"empty", JwtClaims{}, ""},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			gotRole, _ := tt.jcc.AppTraineeRole()

			if gotRole != tt.wantRole {
				t.Errorf("got %s want %s ", gotRole, tt.wantRole)
			}
		})
	}
}

func TestIsTrainer(t *testing.T) {
	var test = []struct {
		name string
		jcc  JwtClaims
		want bool
	}{
		{"trainer", jccFixtureTrainer(), true},
		{"trainee", jccFixtureTrainee(), false},
		{"empty", JwtClaims{}, false},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got_val := tt.jcc.IsTrainer()

			if got_val != tt.want {
				t.Errorf("got %t, want %t", got_val, tt.want)
			}

		})
	}
}
