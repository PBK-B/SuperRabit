// Code generated by entc, DO NOT EDIT.

package app

import (
	"time"
	"yayar/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BundleID applies equality check predicate on the "bundle_id" field. It's identical to BundleIDEQ.
func BundleID(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBundleID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Logo applies equality check predicate on the "logo" field. It's identical to LogoEQ.
func Logo(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// LatestVersion applies equality check predicate on the "latest_version" field. It's identical to LatestVersionEQ.
func LatestVersion(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatestVersion), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// BundleIDEQ applies the EQ predicate on the "bundle_id" field.
func BundleIDEQ(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBundleID), v))
	})
}

// BundleIDNEQ applies the NEQ predicate on the "bundle_id" field.
func BundleIDNEQ(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBundleID), v))
	})
}

// BundleIDIn applies the In predicate on the "bundle_id" field.
func BundleIDIn(vs ...string) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBundleID), v...))
	})
}

// BundleIDNotIn applies the NotIn predicate on the "bundle_id" field.
func BundleIDNotIn(vs ...string) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBundleID), v...))
	})
}

// BundleIDGT applies the GT predicate on the "bundle_id" field.
func BundleIDGT(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBundleID), v))
	})
}

// BundleIDGTE applies the GTE predicate on the "bundle_id" field.
func BundleIDGTE(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBundleID), v))
	})
}

// BundleIDLT applies the LT predicate on the "bundle_id" field.
func BundleIDLT(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBundleID), v))
	})
}

// BundleIDLTE applies the LTE predicate on the "bundle_id" field.
func BundleIDLTE(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBundleID), v))
	})
}

// BundleIDContains applies the Contains predicate on the "bundle_id" field.
func BundleIDContains(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBundleID), v))
	})
}

// BundleIDHasPrefix applies the HasPrefix predicate on the "bundle_id" field.
func BundleIDHasPrefix(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBundleID), v))
	})
}

// BundleIDHasSuffix applies the HasSuffix predicate on the "bundle_id" field.
func BundleIDHasSuffix(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBundleID), v))
	})
}

// BundleIDEqualFold applies the EqualFold predicate on the "bundle_id" field.
func BundleIDEqualFold(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBundleID), v))
	})
}

// BundleIDContainsFold applies the ContainsFold predicate on the "bundle_id" field.
func BundleIDContainsFold(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBundleID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// LogoEQ applies the EQ predicate on the "logo" field.
func LogoEQ(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// LogoNEQ applies the NEQ predicate on the "logo" field.
func LogoNEQ(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogo), v))
	})
}

// LogoIn applies the In predicate on the "logo" field.
func LogoIn(vs ...string) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLogo), v...))
	})
}

// LogoNotIn applies the NotIn predicate on the "logo" field.
func LogoNotIn(vs ...string) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLogo), v...))
	})
}

// LogoGT applies the GT predicate on the "logo" field.
func LogoGT(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogo), v))
	})
}

// LogoGTE applies the GTE predicate on the "logo" field.
func LogoGTE(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogo), v))
	})
}

// LogoLT applies the LT predicate on the "logo" field.
func LogoLT(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogo), v))
	})
}

// LogoLTE applies the LTE predicate on the "logo" field.
func LogoLTE(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogo), v))
	})
}

// LogoContains applies the Contains predicate on the "logo" field.
func LogoContains(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogo), v))
	})
}

// LogoHasPrefix applies the HasPrefix predicate on the "logo" field.
func LogoHasPrefix(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogo), v))
	})
}

// LogoHasSuffix applies the HasSuffix predicate on the "logo" field.
func LogoHasSuffix(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogo), v))
	})
}

// LogoEqualFold applies the EqualFold predicate on the "logo" field.
func LogoEqualFold(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogo), v))
	})
}

// LogoContainsFold applies the ContainsFold predicate on the "logo" field.
func LogoContainsFold(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogo), v))
	})
}

// LatestVersionEQ applies the EQ predicate on the "latest_version" field.
func LatestVersionEQ(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatestVersion), v))
	})
}

// LatestVersionNEQ applies the NEQ predicate on the "latest_version" field.
func LatestVersionNEQ(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLatestVersion), v))
	})
}

// LatestVersionIn applies the In predicate on the "latest_version" field.
func LatestVersionIn(vs ...string) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLatestVersion), v...))
	})
}

// LatestVersionNotIn applies the NotIn predicate on the "latest_version" field.
func LatestVersionNotIn(vs ...string) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLatestVersion), v...))
	})
}

// LatestVersionGT applies the GT predicate on the "latest_version" field.
func LatestVersionGT(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLatestVersion), v))
	})
}

// LatestVersionGTE applies the GTE predicate on the "latest_version" field.
func LatestVersionGTE(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLatestVersion), v))
	})
}

// LatestVersionLT applies the LT predicate on the "latest_version" field.
func LatestVersionLT(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLatestVersion), v))
	})
}

// LatestVersionLTE applies the LTE predicate on the "latest_version" field.
func LatestVersionLTE(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLatestVersion), v))
	})
}

// LatestVersionContains applies the Contains predicate on the "latest_version" field.
func LatestVersionContains(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLatestVersion), v))
	})
}

// LatestVersionHasPrefix applies the HasPrefix predicate on the "latest_version" field.
func LatestVersionHasPrefix(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLatestVersion), v))
	})
}

// LatestVersionHasSuffix applies the HasSuffix predicate on the "latest_version" field.
func LatestVersionHasSuffix(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLatestVersion), v))
	})
}

// LatestVersionIsNil applies the IsNil predicate on the "latest_version" field.
func LatestVersionIsNil() predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLatestVersion)))
	})
}

// LatestVersionNotNil applies the NotNil predicate on the "latest_version" field.
func LatestVersionNotNil() predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLatestVersion)))
	})
}

// LatestVersionEqualFold applies the EqualFold predicate on the "latest_version" field.
func LatestVersionEqualFold(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLatestVersion), v))
	})
}

// LatestVersionContainsFold applies the ContainsFold predicate on the "latest_version" field.
func LatestVersionContainsFold(v string) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLatestVersion), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.App {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.App(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.App {
	return predicate.App(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVersions applies the HasEdge predicate on the "versions" edge.
func HasVersions() predicate.App {
	return predicate.App(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VersionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, VersionsTable, VersionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVersionsWith applies the HasEdge predicate on the "versions" edge with a given conditions (other predicates).
func HasVersionsWith(preds ...predicate.Version) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VersionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, VersionsTable, VersionsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.App) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.App) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.App) predicate.App {
	return predicate.App(func(s *sql.Selector) {
		p(s.Not())
	})
}
