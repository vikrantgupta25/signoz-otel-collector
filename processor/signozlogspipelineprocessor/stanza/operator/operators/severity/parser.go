// Brought in as is from opentelemetry-collector-contrib

package severity // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/severity"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Parser is an operator that parses severity from a field to an entry.
type Parser struct {
	helper.TransformerOperator
	helper.SeverityParser
}

// Process will parse severity from an entry.
func (p *Parser) Process(ctx context.Context, entry *entry.Entry) error {
	return p.ProcessWith(ctx, entry, p.Parse)
}
