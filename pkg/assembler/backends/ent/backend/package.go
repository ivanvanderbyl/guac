package backend

import (
	"context"
	"log"

	"entgo.io/ent/dialect/sql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagenamespace"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagenode"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/pkg/errors"
)

func (b *EntBackend) Packages(ctx context.Context, pkgSpec *model.PkgSpec) ([]*model.Package, error) {
	query := b.client.PackageNode.Query().Order(ent.Asc(packagenode.FieldType))

	paths := getPreloads(ctx)
	if len(paths) > 0 {
		log.Println("Preloading Packages", "paths", paths)
	}

	if pkgSpec == nil {
		pkgSpec = &model.PkgSpec{}
	}

	query.Where(optionalPredicate(pkgSpec.Type, packagenode.TypeEQ))

	if PathContains(paths, "namespaces") {
		query.WithNamespaces(func(q *ent.PackageNamespaceQuery) {
			q.Order(ent.Asc(packagenamespace.FieldNamespace))
			q.Where(optionalPredicate(pkgSpec.Namespace, packagenamespace.NamespaceEQ))

			if PathContains(paths, "namespaces.names") {
				q.WithNames(func(q *ent.PackageNameQuery) {
					q.Order(ent.Asc(packagename.FieldName))
					q.Where(optionalPredicate(pkgSpec.Name, packagename.NameEQ))

					if PathContains(paths, "namespaces.names.versions") {
						q.WithVersions(func(q *ent.PackageVersionQuery) {
							q.Order(ent.Asc(packageversion.FieldVersion))
							q.Where(optionalPredicate(pkgSpec.Version, packageversion.VersionEQ))
						})
					}
				})
			}
		})
	}

	// FIXME: (ivanvanderbyl) This could be much more compact and use a single query as above.
	if pkgSpec != nil {
		query.Where(optionalPredicate(pkgSpec.ID, IDEQ))
	} else {
		query.Limit(100)
	}

	pkgs, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	return collect(pkgs, toModelPackage), nil
}

func (b *EntBackend) IngestPackage(ctx context.Context, pkg model.PkgInputSpec) (*model.Package, error) {
	pvID, err := WithinTX(ctx, b.client, func(ctx context.Context) (*int, error) {
		client := ent.FromContext(ctx)
		pvID, err := ingestPackage(ctx, client, pkg)
		if err != nil {
			return nil, err
		}

		return &pvID, nil
	})
	if err != nil {
		return nil, err
	}

	record, err := b.client.PackageVersion.Query().
		Where(packageversion.ID(*pvID)).QueryName().QueryNamespace().QueryPackage().
		WithNamespaces(func(q *ent.PackageNamespaceQuery) {
			q.Order(ent.Asc(packagenamespace.FieldNamespace))
			q.WithNames(func(q *ent.PackageNameQuery) {
				q.Order(ent.Asc(packagename.FieldName))
				q.WithVersions(func(q *ent.PackageVersionQuery) {
					q.Order(ent.Asc(packageversion.FieldVersion))
				})
			})
		}).
		Only(ctx)

	// TODO: Figure out if we need to preload the edges from the graphql query
	// record, err := b.client.PackageNode.Query().Where(packagenode.ID(*pvID)).
	// 	WithNamespaces(func(q *PackageNamespaceQuery) {
	// 		q.Order(Asc(packagenamespace.FieldNamespace))
	// 		q.WithNames(func(q *PackageNameQuery) {
	// 			q.Order(Asc(packagename.FieldName))
	// 			q.WithVersions(func(q *PackageVersionQuery) {
	// 				q.Order(Asc(packageversion.FieldVersion))
	// 			})
	// 		})
	// 	}).
	// 	Only(ctx)
	if err != nil {
		return nil, err
	}

	return toModelPackage(record), nil
}

func ingestPackage(ctx context.Context, client *ent.Client, pkg model.PkgInputSpec) (int, error) {
	// ingestPackage is used in multiple places, so we extract it to a function.
	pkgID, err := client.PackageNode.Create().SetType(pkg.Type).
		OnConflict(sql.ConflictColumns(packagenode.FieldType)).UpdateNewValues().ID(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "upsert package node")
	}

	if pkg.Namespace == nil {
		empty := ""
		pkg.Namespace = &empty
	}

	nsID, err := client.PackageNamespace.Create().SetPackageID(pkgID).SetNamespace(*pkg.Namespace).
		OnConflict(sql.ConflictColumns(packagenamespace.FieldNamespace, packagenamespace.FieldPackageID)).UpdateNewValues().ID(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "upsert package namespace")
	}

	nameID, err := client.PackageName.Create().SetNamespaceID(nsID).SetName(pkg.Name).
		OnConflict(sql.ConflictColumns(packagename.FieldName, packagename.FieldNamespaceID)).UpdateNewValues().ID(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "upsert package name")
	}

	if pkg.Version == nil {
		empty := ""
		pkg.Version = &empty
	}
	pvID, err := client.PackageVersion.Create().
		SetNameID(nameID).
		SetVersion(*pkg.Version).
		SetSubpath(valueOrDefault(pkg.Subpath, "")).
		SetQualifiers(qualifiersToString(pkg.Qualifiers)).
		OnConflict(
			sql.ConflictColumns(
				packageversion.FieldVersion,
				packageversion.FieldSubpath,
				packageversion.FieldQualifiers,
				packageversion.FieldNameID,
			),
		).
		UpdateNewValues().ID(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "upsert package version")
	}
	return pvID, nil
}
