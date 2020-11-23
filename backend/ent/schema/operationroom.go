package schema
import (
 "github.com/facebookincubator/ent"
 "github.com/facebookincubator/ent/schema/field"
 "github.com/facebookincubator/ent/schema/edge"
)
// Operationroom holds the schema definition for the Operationroom entity.
type Operationroom struct {
 ent.Schema
}
// Fields of the Operationroom.
func (Operationroom) Fields() []ent.Field {
 return []ent.Field{
 field.String("operationroom_name").NotEmpty(),
 }
}
// Edges of the Operationroom.
func (Operationroom) Edges() []ent.Edge {
 return []ent.Edge{
    edge.To("operationroom_id",Booking.Type),
 }
}