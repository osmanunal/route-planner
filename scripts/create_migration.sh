#!/bin/bash

MIGRATION_DIR="migration/migrations"
mkdir -p $MIGRATION_DIR

TIMESTAMP=$(date +%Y%m%d%H%M%S)
DESCRIPTION=$(echo $1 | tr ' ' '_' | tr '[:upper:]' '[:lower:]')
FILENAME="${MIGRATION_DIR}/${TIMESTAMP}_${DESCRIPTION}.go"

# Create migration file
cat > $FILENAME << EOF
package migrations

import (
	"context"
	"route-planner/migration"

	"github.com/uptrace/bun"
)

func up(ctx context.Context, db *bun.DB) error {
	return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		return nil
	})
}

func down(ctx context.Context, db *bun.DB) error {
	return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		return nil
	})
}
func init() {
	migration.Migrations.MustRegister(up, down)
}
EOF

echo "Yeni migration dosyası oluşturuldu: $FILENAME" 