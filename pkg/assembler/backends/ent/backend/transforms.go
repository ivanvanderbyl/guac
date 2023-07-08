package backend

import (
	"fmt"

	"github.com/guacsec/guac/pkg/assembler/backends/ent"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/dependency"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

func toModelArtifact(a *ent.Artifact) *model.Artifact {
	return &model.Artifact{
		ID:        nodeID(a.ID),
		Algorithm: a.Algorithm,
		Digest:    a.Digest,
	}
}

func toModelBuilder(b *ent.Builder) *model.Builder {
	return &model.Builder{
		ID:  nodeID(b.ID),
		URI: b.URI,
	}
}

func backReferencePackageVersion(pv *ent.PackageVersion) *ent.PackageType {
	if pv != nil && pv.Edges.Name != nil &&
		pv.Edges.Name.Edges.Namespace != nil &&
		pv.Edges.Name.Edges.Namespace.Edges.Package != nil {
		pn := pv.Edges.Name
		ns := pn.Edges.Namespace
		pt := ns.Edges.Package
		pn.Edges.Versions = []*ent.PackageVersion{pv}
		ns.Edges.Names = []*ent.PackageName{pn}
		pt.Edges.Namespaces = []*ent.PackageNamespace{ns}
		return pt
	}
	return nil
}

func backReferencePackageName(pn *ent.PackageName) *ent.PackageType {
	if pn.Edges.Namespace != nil &&
		pn.Edges.Namespace.Edges.Package != nil {
		ns := pn.Edges.Namespace
		pt := ns.Edges.Package
		ns.Edges.Names = []*ent.PackageName{pn}
		pt.Edges.Namespaces = []*ent.PackageNamespace{ns}
		return pt
	}
	return nil
}

func backReferenceSourceName(sn *ent.SourceName) *ent.SourceType {
	if sn.Edges.Namespace != nil {
		sns := sn.Edges.Namespace
		sns.Edges.Names = []*ent.SourceName{sn}
		st := sns.Edges.SourceType
		st.Edges.Namespaces = []*ent.SourceNamespace{sns}
		return st
	}
	return nil
}

func toModelPackage(p *ent.PackageType) *model.Package {
	if p == nil {
		return nil
	}
	return &model.Package{
		ID:         nodeID(p.ID),
		Type:       p.Type,
		Namespaces: collect(p.Edges.Namespaces, toModelNamespace),
	}
}

func toModelNamespace(n *ent.PackageNamespace) *model.PackageNamespace {
	if n == nil {
		return nil
	}
	return &model.PackageNamespace{
		ID:        nodeID(n.ID),
		Namespace: n.Namespace,
		Names:     collect(n.Edges.Names, toModelPackageName),
	}
}

func toModelPackageName(n *ent.PackageName) *model.PackageName {
	if n == nil {
		return nil
	}
	return &model.PackageName{
		ID:       nodeID(n.ID),
		Name:     n.Name,
		Versions: collect(n.Edges.Versions, toModelPackageVersion),
	}
}

func toModelSource(s *ent.SourceType) *model.Source {
	if s == nil {
		return nil
	}
	return &model.Source{
		ID:   nodeID(s.ID),
		Type: s.Type,
		Namespaces: collect(s.Edges.Namespaces, func(n *ent.SourceNamespace) *model.SourceNamespace {
			return &model.SourceNamespace{
				ID:        nodeID(n.ID),
				Namespace: n.Namespace,
				Names: collect(n.Edges.Names, func(n *ent.SourceName) *model.SourceName {
					return &model.SourceName{
						ID:     nodeID(n.ID),
						Name:   n.Name,
						Tag:    &n.Tag,
						Commit: &n.Commit,
					}
				}),
			}
		}),
	}
}

func toModelPackageVersion(v *ent.PackageVersion) *model.PackageVersion {

	return &model.PackageVersion{
		ID:         nodeID(v.ID),
		Version:    v.Version,
		Qualifiers: toPtrSlice(v.Qualifiers),
		Subpath:    v.Subpath,
	}
}

// collect is a simple helper to transform collections of a certain type to another type
// using the transform function func(T) R
func collect[T any, R any](items []T, transformer func(T) R) []R {
	if items == nil {
		return nil
	}
	out := make([]R, len(items))
	for i, item := range items {
		out[i] = transformer(item)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func nodeID(id int) string {
	return fmt.Sprintf("%d", id)
}

func stringOrEmpty(s *string) string {
	return valueOrDefault(s, "")
}

func valueOrDefault[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}

func toModelIsOccurrenceWithSubject(o *ent.Occurrence) *model.IsOccurrence {
	return &model.IsOccurrence{
		ID:            nodeID(o.ID),
		Subject:       toOccurrenceSubject(o),
		Artifact:      toModelArtifact(o.Edges.Artifact),
		Justification: o.Justification,
		Origin:        o.Origin,
		Collector:     o.Collector,
	}
}

func toModelIsOccurrence(o *ent.Occurrence, sub model.PackageOrSource) *model.IsOccurrence {
	return &model.IsOccurrence{
		ID:            nodeID(o.ID),
		Subject:       sub,
		Artifact:      toModelArtifact(o.Edges.Artifact),
		Justification: o.Justification,
		Origin:        o.Origin,
		Collector:     o.Collector,
	}
}

func toOccurrenceSubject(oc *ent.Occurrence) model.PackageOrSource {
	if oc.Edges.Package != nil {
		return toModelPackage(backReferencePackageVersion(oc.Edges.Package))
	} else if oc.Edges.Source != nil &&
		// FIXME: (ivanvanderbyl) Refactor into backReferenceSubject(...)
		oc.Edges.Source.Edges.Namespace != nil &&
		oc.Edges.Source.Edges.Namespace.Edges.SourceType != nil {

		// Manually construct back references to avoid another 3 queries
		s := oc.Edges.Source
		ns := s.Edges.Namespace
		ns.Edges.Names = []*ent.SourceName{s}
		st := ns.Edges.SourceType
		st.Edges.Namespaces = []*ent.SourceNamespace{ns}
		return toModelSource(st)
	}
	return nil
}

func pkgQualifierInputSpecToQuerySpec(input []*model.PackageQualifierInputSpec) []*model.PackageQualifierSpec {
	if input == nil {
		return nil
	}
	out := make([]*model.PackageQualifierSpec, len(input))
	for i, in := range input {
		out[i] = &model.PackageQualifierSpec{
			Key:   in.Key,
			Value: &in.Value,
		}
	}
	return out
}

func toModelIsDependencyWithBackrefs(id *ent.Dependency) *model.IsDependency {
	return toModelIsDependency(id, true)
}

func toModelIsDependencyWithoutBackrefs(id *ent.Dependency) *model.IsDependency {
	return toModelIsDependency(id, false)
}

func toModelIsDependency(id *ent.Dependency, backrefs bool) *model.IsDependency {
	var pkg *model.Package
	var depPkg *model.Package
	if backrefs {
		pkg = toModelPackage(backReferencePackageVersion(id.Edges.Package))
		depPkg = toModelPackage(backReferencePackageName(id.Edges.DependentPackage))
	} else {
		pkg = toModelPackage(id.Edges.Package.Edges.Name.Edges.Namespace.Edges.Package)
		depPkg = toModelPackage(id.Edges.DependentPackage.Edges.Namespace.Edges.Package)
	}

	return &model.IsDependency{
		ID:               nodeID(id.ID),
		Package:          pkg,
		DependentPackage: depPkg,
		VersionRange:     id.VersionRange,
		DependencyType:   dependencyTypeFromEnum(id.DependencyType),
		Justification:    id.Justification,
		Origin:           id.Origin,
		Collector:        id.Collector,
	}
}

func dependencyTypeFromEnum(t dependency.DependencyType) model.DependencyType {
	switch t {
	case dependency.DependencyTypeDIRECT:
		return model.DependencyTypeDirect
	case dependency.DependencyTypeINDIRECT:
		return model.DependencyTypeIndirect
	default:
		return model.DependencyTypeUnknown
	}
}

func toModelHasSBOM(sbom *ent.BillOfMaterials) *model.HasSbom {
	return &model.HasSbom{
		ID:               nodeID(sbom.ID),
		Subject:          toPackageOrArtifact(sbom.Edges.Package, sbom.Edges.Artifact),
		URI:              sbom.URI,
		Algorithm:        sbom.Algorithm,
		Digest:           sbom.Digest,
		DownloadLocation: sbom.DownloadLocation,
		Origin:           sbom.Origin,
		Collector:        sbom.Collector,
		Annotations:      toPtrSlice(sbom.Annotations),
	}
}

func toPackageOrArtifact(p *ent.PackageVersion, a *ent.Artifact) model.PackageOrArtifact {
	if p != nil {
		return toModelPackage(backReferencePackageVersion(p))
	} else if a != nil {
		return toModelArtifact(a)
	}
	return nil
}
