package xsd

import (
	"fmt"
	"os"
	"path/filepath"
)

// ResolveIncludes recursively resolves xs:include directives in this schema
// and its included schemas. Content from included schemas (elements, types,
// attributes, groups, annotations) is merged into this schema. Cycle
// detection prevents infinite recursion when schemas reference each other.
//
// selfPath is the path of the file (or containing WSDL file) that this schema
// was loaded from. It's used both to mark the root as visited (preventing
// cycles back to it) and to resolve relative schemaLocation values in child
// includes.
//
// After a successful call, Includes is cleared to signal that no unresolved
// includes remain.
//
// Only xs:include is resolved here. xs:include requires the included schema
// to share the parent's target namespace (or have no target namespace, the
// "chameleon include" case); both produce a single merged namespace. xs:import,
// which brings in a different namespace, is not handled by this method and
// should be resolved separately if needed.
func (s *Schema) ResolveIncludes(selfPath string) error {
	visited := map[string]bool{}
	if selfPath != "" {
		abs, err := filepath.Abs(selfPath)
		if err != nil {
			return fmt.Errorf("resolve self path %q: %w", selfPath, err)
		}
		visited[abs] = true
	}
	return s.resolveIncludes(filepath.Dir(selfPath), visited)
}

func (s *Schema) resolveIncludes(baseDir string, visited map[string]bool) error {
	for _, inc := range s.Includes {
		if inc.SchemaLocation == "" {
			continue
		}
		loc := inc.SchemaLocation
		if !filepath.IsAbs(loc) {
			loc = filepath.Join(baseDir, loc)
		}
		abs, err := filepath.Abs(loc)
		if err != nil {
			return fmt.Errorf("xs:include %q: %w", inc.SchemaLocation, err)
		}
		if visited[abs] {
			continue
		}
		visited[abs] = true

		f, err := os.Open(abs)
		if err != nil {
			return fmt.Errorf("xs:include %q: %w", inc.SchemaLocation, err)
		}
		included, err := Parse(f)
		closeErr := f.Close()
		if err != nil {
			return fmt.Errorf("xs:include %q: parse: %w", inc.SchemaLocation, err)
		}
		if closeErr != nil {
			return fmt.Errorf("xs:include %q: close: %w", inc.SchemaLocation, closeErr)
		}

		// Recurse with the included file's directory as the new base, so its
		// own xs:include directives resolve relative to itself.
		if err := included.resolveIncludes(filepath.Dir(abs), visited); err != nil {
			return err
		}

		// Merge the included schema's content into this schema. xs:include
		// semantics put everything in the parent's target namespace.
		s.Elements = append(s.Elements, included.Elements...)
		s.ComplexTypes = append(s.ComplexTypes, included.ComplexTypes...)
		s.SimpleTypes = append(s.SimpleTypes, included.SimpleTypes...)
		s.Attributes = append(s.Attributes, included.Attributes...)
		s.AttributeGroups = append(s.AttributeGroups, included.AttributeGroups...)
		s.Groups = append(s.Groups, included.Groups...)
		s.Annotations = append(s.Annotations, included.Annotations...)

		// Preserve any xs:import the included schema had; they may need
		// separate resolution by a future xs:import pass.
		s.Imports = append(s.Imports, included.Imports...)
	}

	s.Includes = nil
	return nil
}
