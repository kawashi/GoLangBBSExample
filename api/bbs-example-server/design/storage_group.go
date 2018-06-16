package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("bbs_example", func() {
	Description("BBS Example Model.")
	Store("MySQL", gorma.MySQL, func() {
		Model("UserPost", func() {
			RendersTo(UserPostMedia)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("message", gorma.String)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
		})
	})
})
