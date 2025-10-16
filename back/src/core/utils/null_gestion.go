// Package utils contient des fonctions utilitaires globales pour le projet.
//
//revive:disable:var-naming
package utils

import (
	"database/sql"
	"time"
)

func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func NullStringToStringPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func NullTimeToPointer(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

func NullInt64ToPointer(nt sql.NullInt64) *int64 {
	if nt.Valid {
		return &nt.Int64
	}
	return nil
}

func NullFloat64ToPointer(nt sql.NullFloat64) *float64 {
	if nt.Valid {
		return &nt.Float64
	}
	return nil
}

// Contains retourne si le liste contient l'élèment.
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func AddString(updates map[string]interface{}, key, value string) {
	if value != "" {
		updates[key] = value
	}
}

func AddBoolPtr(updates map[string]interface{}, key string, value *bool) {
	if value != nil {
		updates[key] = *value
	}
}

func AddFloat(updates map[string]interface{}, key string, value float64) {
	if value != 0 {
		updates[key] = value
	}
}

func AddDuration(updates map[string]interface{}, key string, d time.Duration) {
	if d != 0 {
		updates[key] = d
	}
}

func NullStringValidation(s *string) sql.NullString {
	if s != nil && *s != "" {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{Valid: false}
}

func NullFloat64Validation(f *float64) sql.NullFloat64 {
	if f != nil && *f != 0 {
		return sql.NullFloat64{Float64: *f, Valid: true}
	}
	return sql.NullFloat64{Valid: false}
}

func NullInt64Validation(i *int64) sql.NullInt64 {
	if i != nil {
		return sql.NullInt64{Int64: *i, Valid: true}
	}
	return sql.NullInt64{Valid: false}
}

func NullTimeValidation(t *time.Time) sql.NullTime {
	if t != nil {
		return sql.NullTime{Time: *t, Valid: true}
	}
	return sql.NullTime{Valid: false}
}
