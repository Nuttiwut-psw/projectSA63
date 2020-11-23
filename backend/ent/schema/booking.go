package schema
import (
 "github.com/facebookincubator/ent"
 "github.com/facebookincubator/ent/schema/field"
 "github.com/facebookincubator/ent/schema/edge"
)
// Booking holds the schema definition for the Booking entity.
type Booking struct {
 ent.Schema
}
// Fields of the Booking.
func (Booking) Fields() []ent.Field {
 return []ent.Field{
    field.Time("date"),
 }
}
// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
 return []ent.Edge{
    edge.From("doctor_id",User.Type).Ref("doctor_id").Unique(),
    edge.From("patient_id",Patient.Type).Ref("patient_id").Unique(),
    edge.From("operationroom_id",Operationroom.Type).Ref("operationroom_id").Unique(),
 }
}