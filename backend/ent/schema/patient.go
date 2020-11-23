package schema
import (
 "github.com/facebookincubator/ent"
 "github.com/facebookincubator/ent/schema/field"
 "github.com/facebookincubator/ent/schema/edge"
)
// Patient holds the schema definition for the Patient entity.
type Patient struct {
 ent.Schema
}
// Fields of the Patient.
func (Patient) Fields() []ent.Field {
 return []ent.Field{
 field.String("patient_name").NotEmpty(),
 field.Int("patient_age"),
 }
}
// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
 return []ent.Edge{
    edge.To("patient_id",Booking.Type),
 }
}