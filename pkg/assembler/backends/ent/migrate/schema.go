// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArtifactsColumns holds the columns for the "artifacts" table.
	ArtifactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "algorithm", Type: field.TypeString},
		{Name: "digest", Type: field.TypeString},
	}
	// ArtifactsTable holds the schema information for the "artifacts" table.
	ArtifactsTable = &schema.Table{
		Name:       "artifacts",
		Columns:    ArtifactsColumns,
		PrimaryKey: []*schema.Column{ArtifactsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "artifact_algorithm_digest",
				Unique:  true,
				Columns: []*schema.Column{ArtifactsColumns[1], ArtifactsColumns[2]},
			},
		},
	}
	// BuilderNodesColumns holds the columns for the "builder_nodes" table.
	BuilderNodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uri", Type: field.TypeString, Unique: true},
	}
	// BuilderNodesTable holds the schema information for the "builder_nodes" table.
	BuilderNodesTable = &schema.Table{
		Name:       "builder_nodes",
		Columns:    BuilderNodesColumns,
		PrimaryKey: []*schema.Column{BuilderNodesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "buildernode_uri",
				Unique:  true,
				Columns: []*schema.Column{BuilderNodesColumns[1]},
			},
		},
	}
	// IsDependenciesColumns holds the columns for the "is_dependencies" table.
	IsDependenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "version_range", Type: field.TypeString},
		{Name: "dependency_type", Type: field.TypeString},
		{Name: "justification", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "package_id", Type: field.TypeInt},
		{Name: "dependent_package_id", Type: field.TypeInt},
	}
	// IsDependenciesTable holds the schema information for the "is_dependencies" table.
	IsDependenciesTable = &schema.Table{
		Name:       "is_dependencies",
		Columns:    IsDependenciesColumns,
		PrimaryKey: []*schema.Column{IsDependenciesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "is_dependencies_package_versions_package",
				Columns:    []*schema.Column{IsDependenciesColumns[6]},
				RefColumns: []*schema.Column{PackageVersionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "is_dependencies_package_names_dependent_package",
				Columns:    []*schema.Column{IsDependenciesColumns[7]},
				RefColumns: []*schema.Column{PackageNamesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "isdependency_version_range_dependency_type_justification_origin_collector_package_id_dependent_package_id",
				Unique:  true,
				Columns: []*schema.Column{IsDependenciesColumns[1], IsDependenciesColumns[2], IsDependenciesColumns[3], IsDependenciesColumns[4], IsDependenciesColumns[5], IsDependenciesColumns[6], IsDependenciesColumns[7]},
			},
		},
	}
	// IsOccurrencesColumns holds the columns for the "is_occurrences" table.
	IsOccurrencesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "justification", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "collector", Type: field.TypeString},
		{Name: "package_id", Type: field.TypeInt, Nullable: true},
		{Name: "source_id", Type: field.TypeInt, Nullable: true},
		{Name: "artifact_id", Type: field.TypeInt},
	}
	// IsOccurrencesTable holds the schema information for the "is_occurrences" table.
	IsOccurrencesTable = &schema.Table{
		Name:       "is_occurrences",
		Columns:    IsOccurrencesColumns,
		PrimaryKey: []*schema.Column{IsOccurrencesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "is_occurrences_package_versions_package_version",
				Columns:    []*schema.Column{IsOccurrencesColumns[4]},
				RefColumns: []*schema.Column{PackageVersionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "is_occurrences_source_names_source",
				Columns:    []*schema.Column{IsOccurrencesColumns[5]},
				RefColumns: []*schema.Column{SourceNamesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "is_occurrences_artifacts_artifact",
				Columns:    []*schema.Column{IsOccurrencesColumns[6]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "occurrence_unique_package",
				Unique:  true,
				Columns: []*schema.Column{IsOccurrencesColumns[1], IsOccurrencesColumns[2], IsOccurrencesColumns[3], IsOccurrencesColumns[5], IsOccurrencesColumns[4], IsOccurrencesColumns[6]},
			},
		},
	}
	// PackageNamesColumns holds the columns for the "package_names" table.
	PackageNamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "namespace_id", Type: field.TypeInt},
	}
	// PackageNamesTable holds the schema information for the "package_names" table.
	PackageNamesTable = &schema.Table{
		Name:       "package_names",
		Columns:    PackageNamesColumns,
		PrimaryKey: []*schema.Column{PackageNamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "package_names_package_namespaces_names",
				Columns:    []*schema.Column{PackageNamesColumns[2]},
				RefColumns: []*schema.Column{PackageNamespacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "packagename_name_namespace_id",
				Unique:  true,
				Columns: []*schema.Column{PackageNamesColumns[1], PackageNamesColumns[2]},
			},
		},
	}
	// PackageNamespacesColumns holds the columns for the "package_namespaces" table.
	PackageNamespacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "namespace", Type: field.TypeString},
		{Name: "package_id", Type: field.TypeInt},
	}
	// PackageNamespacesTable holds the schema information for the "package_namespaces" table.
	PackageNamespacesTable = &schema.Table{
		Name:       "package_namespaces",
		Columns:    PackageNamespacesColumns,
		PrimaryKey: []*schema.Column{PackageNamespacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "package_namespaces_package_nodes_namespaces",
				Columns:    []*schema.Column{PackageNamespacesColumns[2]},
				RefColumns: []*schema.Column{PackageNodesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "packagenamespace_namespace_package_id",
				Unique:  true,
				Columns: []*schema.Column{PackageNamespacesColumns[1], PackageNamespacesColumns[2]},
			},
		},
	}
	// PackageNodesColumns holds the columns for the "package_nodes" table.
	PackageNodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, Unique: true},
	}
	// PackageNodesTable holds the schema information for the "package_nodes" table.
	PackageNodesTable = &schema.Table{
		Name:       "package_nodes",
		Columns:    PackageNodesColumns,
		PrimaryKey: []*schema.Column{PackageNodesColumns[0]},
	}
	// PackageVersionsColumns holds the columns for the "package_versions" table.
	PackageVersionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "version", Type: field.TypeString},
		{Name: "subpath", Type: field.TypeString},
		{Name: "qualifiers", Type: field.TypeString},
		{Name: "name_id", Type: field.TypeInt},
	}
	// PackageVersionsTable holds the schema information for the "package_versions" table.
	PackageVersionsTable = &schema.Table{
		Name:       "package_versions",
		Columns:    PackageVersionsColumns,
		PrimaryKey: []*schema.Column{PackageVersionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "package_versions_package_names_versions",
				Columns:    []*schema.Column{PackageVersionsColumns[4]},
				RefColumns: []*schema.Column{PackageNamesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "packageversion_version_subpath_qualifiers_name_id",
				Unique:  true,
				Columns: []*schema.Column{PackageVersionsColumns[1], PackageVersionsColumns[2], PackageVersionsColumns[3], PackageVersionsColumns[4]},
			},
		},
	}
	// SourcesColumns holds the columns for the "sources" table.
	SourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString},
		{Name: "namespace", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "tag", Type: field.TypeString, Nullable: true},
		{Name: "commit", Type: field.TypeString, Nullable: true},
	}
	// SourcesTable holds the schema information for the "sources" table.
	SourcesTable = &schema.Table{
		Name:       "sources",
		Columns:    SourcesColumns,
		PrimaryKey: []*schema.Column{SourcesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "source_type_namespace_name_tag_commit",
				Unique:  true,
				Columns: []*schema.Column{SourcesColumns[1], SourcesColumns[2], SourcesColumns[3], SourcesColumns[4], SourcesColumns[5]},
			},
		},
	}
	// SourceNamesColumns holds the columns for the "source_names" table.
	SourceNamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "commit", Type: field.TypeString, Nullable: true},
		{Name: "tag", Type: field.TypeString, Nullable: true},
		{Name: "namespace_id", Type: field.TypeInt},
	}
	// SourceNamesTable holds the schema information for the "source_names" table.
	SourceNamesTable = &schema.Table{
		Name:       "source_names",
		Columns:    SourceNamesColumns,
		PrimaryKey: []*schema.Column{SourceNamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "source_names_source_namespaces_namespace",
				Columns:    []*schema.Column{SourceNamesColumns[4]},
				RefColumns: []*schema.Column{SourceNamespacesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "sourcename_name_commit_tag",
				Unique:  true,
				Columns: []*schema.Column{SourceNamesColumns[1], SourceNamesColumns[2], SourceNamesColumns[3]},
			},
		},
	}
	// SourceNamespacesColumns holds the columns for the "source_namespaces" table.
	SourceNamespacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "namespace", Type: field.TypeString},
		{Name: "source_id", Type: field.TypeInt},
	}
	// SourceNamespacesTable holds the schema information for the "source_namespaces" table.
	SourceNamespacesTable = &schema.Table{
		Name:       "source_namespaces",
		Columns:    SourceNamespacesColumns,
		PrimaryKey: []*schema.Column{SourceNamespacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "source_namespaces_sources_source",
				Columns:    []*schema.Column{SourceNamespacesColumns[2]},
				RefColumns: []*schema.Column{SourcesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "sourcenamespace_namespace_source_id",
				Unique:  true,
				Columns: []*schema.Column{SourceNamespacesColumns[1], SourceNamespacesColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArtifactsTable,
		BuilderNodesTable,
		IsDependenciesTable,
		IsOccurrencesTable,
		PackageNamesTable,
		PackageNamespacesTable,
		PackageNodesTable,
		PackageVersionsTable,
		SourcesTable,
		SourceNamesTable,
		SourceNamespacesTable,
	}
)

func init() {
	IsDependenciesTable.ForeignKeys[0].RefTable = PackageVersionsTable
	IsDependenciesTable.ForeignKeys[1].RefTable = PackageNamesTable
	IsOccurrencesTable.ForeignKeys[0].RefTable = PackageVersionsTable
	IsOccurrencesTable.ForeignKeys[1].RefTable = SourceNamesTable
	IsOccurrencesTable.ForeignKeys[2].RefTable = ArtifactsTable
	PackageNamesTable.ForeignKeys[0].RefTable = PackageNamespacesTable
	PackageNamespacesTable.ForeignKeys[0].RefTable = PackageNodesTable
	PackageVersionsTable.ForeignKeys[0].RefTable = PackageNamesTable
	SourceNamesTable.ForeignKeys[0].RefTable = SourceNamespacesTable
	SourceNamespacesTable.ForeignKeys[0].RefTable = SourcesTable
}
