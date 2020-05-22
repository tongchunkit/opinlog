package opinlog

// Message is a function that returns a string message with the fields
// How to use sample:
// func UnmarshalError(obj interface{}, err error) opinlog.Message {
//   return func() (string, []opinlog.Field) {
//     return "unmarshal error",
//       opinlog.NewFields(opinlog.NewField("err", err), opinlog.NewField("obj", obj))
//   }
// }
type Message = func() (string, []Field)
