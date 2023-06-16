/*
 */
package main

func main() {
	// ctx := context.TODO()

	// Create a Subject.

	// mySenzing := &senzing.SenzingImpl{
	// 	GrpcDialOptions: []grpc.DialOption{},
	// 	GrpcTarget:      "",
	// 	ObserverOrigin:  "",
	// 	Observers:       []observer.Observer{},
	// }

	// G2 style.

	// g2Engine := mySenzing.GetG2Engine(ctx)
	// jsonString := ``
	// result, err := g2Engine.AddRecord(jsonString)

	// Senzing style.

	// engine := mySenzing.GetEngine(ctx)
	// engineAddRecordRequest := &senzing.EngineAddRecordRequest{} // Uses existing import of senzing
	// engineAddRecordRequest := &engine.AddRecordRequest{}        // Requires an import of senzing/engine

	// engineAddRecordResponse, err := engine.AddRecord(engineAddRecordRequest)
	// x := engineAddRecordResponse.Bob

}
