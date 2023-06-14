package backend

import (
	"context"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/guacsec/guac/internal/testing/ptrfrom"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

func (s *Suite) TestIsDependency() {
	type call struct {
		P1    *model.PkgInputSpec
		P2    *model.PkgInputSpec
		Input *model.IsDependencyInputSpec
	}
	tests := []struct {
		Name         string
		IngestPkg    []*model.PkgInputSpec
		Calls        []call
		Query        *model.IsDependencySpec
		Expected     []*model.IsDependency
		ExpIngestErr bool
		ExpQueryErr  bool
	}{
		{
			Name:      "HappyPath",
			IngestPkg: []*model.PkgInputSpec{p1, p2},
			Calls: []call{
				{
					P1: p1,
					P2: p2,
					Input: &model.IsDependencyInputSpec{
						Justification: "test justification",
					},
				},
			},
			Query: &model.IsDependencySpec{
				Justification: ptrfrom.String("test justification"),
			},
			Expected: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p2outName,
					Justification:    "test justification",
				},
			},
		},
		{
			Name:      "Ingest same",
			IngestPkg: []*model.PkgInputSpec{p1, p2},
			Calls: []call{
				{
					P1: p1,
					P2: p2,
					Input: &model.IsDependencyInputSpec{
						Justification: "test justification",
					},
				},
				{
					P1: p1,
					P2: p2,
					Input: &model.IsDependencyInputSpec{
						Justification: "test justification",
					},
				},
			},
			Query: &model.IsDependencySpec{
				Justification: ptrfrom.String("test justification"),
			},
			Expected: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p2outName,
					Justification:    "test justification",
				},
			},
		},
		{
			Name:      "Ingest same, different version",
			IngestPkg: []*model.PkgInputSpec{p1, p2, p3},
			Calls: []call{
				{
					P1: p1,
					P2: p2,
					Input: &model.IsDependencyInputSpec{
						Justification: "test justification",
					},
				},
				{
					P1: p1,
					P2: p3,
					Input: &model.IsDependencyInputSpec{
						Justification: "test justification",
					},
				},
			},
			Query: &model.IsDependencySpec{
				Justification: ptrfrom.String("test justification"),
			},
			Expected: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p2outName,
					Justification:    "test justification",
				},
			},
		},
		{
			Name:      "Query on Justification",
			IngestPkg: []*model.PkgInputSpec{p1, p2},
			Calls: []call{
				{
					P1: p1,
					P2: p2,
					Input: &model.IsDependencyInputSpec{
						Justification: "test justification one",
					},
				},
				{
					P1: p1,
					P2: p2,
					Input: &model.IsDependencyInputSpec{
						Justification: "test justification two",
					},
				},
			},
			Query: &model.IsDependencySpec{
				Justification: ptrfrom.String("test justification one"),
			},
			Expected: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p2outName,
					Justification:    "test justification one",
				},
			},
		},
		{
			Name:      "Query on pkg",
			IngestPkg: []*model.PkgInputSpec{p1, p2, p3},
			Calls: []call{
				{
					P1:    p1,
					P2:    p2,
					Input: &model.IsDependencyInputSpec{},
				},
				{
					P1:    p2,
					P2:    p3,
					Input: &model.IsDependencyInputSpec{},
				},
			},
			Query: &model.IsDependencySpec{
				Package: &model.PkgSpec{
					// ID: ptrfrom.String("5"),
					Type: ptrfrom.String("pypi"),
				},
			},
			Expected: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p2outName,
				},
			},
		},
		{
			Name:      "Query on dep pkg",
			IngestPkg: []*model.PkgInputSpec{p1, p2, p4},
			Calls: []call{
				{
					P1:    p2,
					P2:    p4,
					Input: &model.IsDependencyInputSpec{},
				},
				{
					P1:    p2,
					P2:    p1,
					Input: &model.IsDependencyInputSpec{},
				},
			},
			Query: &model.IsDependencySpec{
				DependentPackage: &model.PkgNameSpec{
					Name: ptrfrom.String("openssl"),
				},
			},
			Expected: []*model.IsDependency{
				{
					Package:          p2out,
					DependentPackage: p4outName,
				},
			},
		},
		{
			Name:      "Query on pkg multiple",
			IngestPkg: []*model.PkgInputSpec{p1, p2, p3},
			Calls: []call{
				{
					P1:    p1,
					P2:    p2,
					Input: &model.IsDependencyInputSpec{},
				},
				{
					P1:    p3,
					P2:    p2,
					Input: &model.IsDependencyInputSpec{},
				},
			},
			Query: &model.IsDependencySpec{
				Package: &model.PkgSpec{
					Type: ptrfrom.String("pypi"),
				},
			},
			Expected: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p1outName,
				},
				{
					Package:          p3out,
					DependentPackage: p1outName,
				},
			},
		},
		{
			Name:      "Query on both pkgs",
			IngestPkg: []*model.PkgInputSpec{p1, p2, p3, p4},
			Calls: []call{
				{
					P1:    p2,
					P2:    p1,
					Input: &model.IsDependencyInputSpec{},
				},
				{
					P1:    p3,
					P2:    p4,
					Input: &model.IsDependencyInputSpec{},
				},
			},
			Query: &model.IsDependencySpec{
				Package: &model.PkgSpec{
					Subpath: ptrfrom.String("saved_model_cli.py"),
				},
				DependentPackage: &model.PkgNameSpec{
					Name: ptrfrom.String("openssl"),
				},
			},
			Expected: []*model.IsDependency{
				{
					Package:          p3out,
					DependentPackage: p4outName,
				},
			},
		},
		{
			Name:      "Query none",
			IngestPkg: []*model.PkgInputSpec{p1, p2, p3},
			Calls: []call{
				{
					P1:    p1,
					P2:    p2,
					Input: &model.IsDependencyInputSpec{},
				},
				{
					P1:    p2,
					P2:    p3,
					Input: &model.IsDependencyInputSpec{},
				},
				{
					P1:    p1,
					P2:    p3,
					Input: &model.IsDependencyInputSpec{},
				},
			},
			Query: &model.IsDependencySpec{
				Package: &model.PkgSpec{
					Subpath: ptrfrom.String("asdf"),
				},
			},
			Expected: nil,
		},
		{
			Name:      "Query on ID",
			IngestPkg: []*model.PkgInputSpec{p1, p2, p3},
			Calls: []call{
				{
					P1:    p1,
					P2:    p2,
					Input: &model.IsDependencyInputSpec{},
				},
				{
					P1:    p2,
					P2:    p3,
					Input: &model.IsDependencyInputSpec{},
				},
				{
					P1:    p1,
					P2:    p3,
					Input: &model.IsDependencyInputSpec{},
				},
			},
			Query: &model.IsDependencySpec{
				ID: ptrfrom.String("9"),
			},
			Expected: []*model.IsDependency{
				{
					Package:          p2out,
					DependentPackage: p1outName,
				},
			},
		},
		{
			Name:      "Query on Range",
			IngestPkg: []*model.PkgInputSpec{p1, p2},
			Calls: []call{
				{
					P1: p1,
					P2: p1,
					Input: &model.IsDependencyInputSpec{
						VersionRange: "1-3",
					},
				},
				{
					P1: p2,
					P2: p1,
					Input: &model.IsDependencyInputSpec{
						VersionRange: "4-5",
					},
				},
			},
			Query: &model.IsDependencySpec{
				VersionRange: ptrfrom.String("1-3"),
			},
			Expected: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p1outName,
					VersionRange:     "1-3",
				},
			},
		},
		{
			Name:      "Query on DependencyType",
			IngestPkg: []*model.PkgInputSpec{p1, p2},
			Calls: []call{
				{
					P1: p1,
					P2: p1,
					Input: &model.IsDependencyInputSpec{
						DependencyType: model.DependencyTypeDirect,
					},
				},
				{
					P1: p2,
					P2: p1,
					Input: &model.IsDependencyInputSpec{
						DependencyType: model.DependencyTypeIndirect,
					},
				},
			},
			Query: &model.IsDependencySpec{
				DependencyType: (*model.DependencyType)(ptrfrom.String(string(model.DependencyTypeIndirect))),
			},
			Expected: []*model.IsDependency{
				{
					Package:          p2out,
					DependentPackage: p1outName,
					DependencyType:   model.DependencyTypeIndirect,
				},
			},
		},
		{
			Name:      "Ingest no P1",
			IngestPkg: []*model.PkgInputSpec{p2},
			Calls: []call{
				{
					P1:    p1,
					P2:    p2,
					Input: &model.IsDependencyInputSpec{},
				},
			},
			ExpIngestErr: true,
		},
		{
			Name:      "Ingest no P2",
			IngestPkg: []*model.PkgInputSpec{p1},
			Calls: []call{
				{
					P1:    p1,
					P2:    p4,
					Input: &model.IsDependencyInputSpec{},
				},
			},
			ExpIngestErr: true,
		},
		{
			Name:      "Query bad ID",
			IngestPkg: []*model.PkgInputSpec{p1, p2, p3},
			Calls: []call{
				{
					P1:    p1,
					P2:    p2,
					Input: &model.IsDependencyInputSpec{},
				},
				{
					P1:    p2,
					P2:    p3,
					Input: &model.IsDependencyInputSpec{},
				},
				{
					P1:    p1,
					P2:    p3,
					Input: &model.IsDependencyInputSpec{},
				},
			},
			Query: &model.IsDependencySpec{
				ID: ptrfrom.String("asdf"),
			},
			ExpQueryErr: true,
		},
	}
	ignoreID := cmp.FilterPath(func(p cmp.Path) bool {
		return strings.Compare(".ID", p[len(p)-1].String()) == 0
	}, cmp.Ignore())
	ctx := context.Background()
	for _, test := range tests {
		s.Run(test.Name, func() {
			b, err := GetBackend(s.Client)
			s.Require().NoError(err, "Could not instantiate testing backend")

			for _, a := range test.IngestPkg {
				if _, err := b.IngestPackage(ctx, *a); err != nil {
					s.Require().NoError(err, "Could not ingest pkg")
				}
			}
			for _, o := range test.Calls {
				_, err := b.IngestDependency(ctx, *o.P1, *o.P2, *o.Input)
				if (err != nil) != test.ExpIngestErr {
					s.T().Fatalf("did not get expected ingest error, want: %v, got: %v", test.ExpIngestErr, err)
				}
				if err != nil {
					return
				}
			}
			got, err := b.IsDependency(ctx, test.Query)
			if (err != nil) != test.ExpQueryErr {
				s.T().Fatalf("did not get expected query error, want: %v, got: %v", test.ExpQueryErr, err)
			}
			if err != nil {
				return
			}

			if diff := cmp.Diff(test.Expected, got, ignoreID); diff != "" {
				s.T().Errorf("Unexpected results. (-want +got):\n%s", diff)
			}
		})
	}
}
