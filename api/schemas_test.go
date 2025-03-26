package api

import "testing"

func TestGetAppRole(t *testing.T) {
	jcc := JwtClaims{
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
	got_role, got_isTrainer := jcc.GetAppRole()
	want_role := "samo_trainee_app"
	want_isTrainer := true

	if got_role != want_role || got_isTrainer != want_isTrainer {
		t.Errorf("got %s | %t , want %s | %t", got_role, got_isTrainer, want_role, want_isTrainer)
	}
}

func TestGetAppRoleTrainee(t *testing.T) {
	jcc := JwtClaims{
		RealmAccess: map[string][]string{
			"roles": {
				"samo_trainee_app",
				"offline_access",
				"default-roles-trainer-helper",
				"uma_authorization",
			},
		},
	}
	got_role, got_isTrainer := jcc.GetAppRole()
	want_role := "samo_trainee_app"
	want_isTrainer := false

	if got_role != want_role || got_isTrainer != want_isTrainer {
		t.Errorf("got %s | %t , want %s | %t", got_role, got_isTrainer, want_role, want_isTrainer)
	}
}
