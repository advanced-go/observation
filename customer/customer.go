package customer

// Customer - customer information
type Entry struct {
	Id    int32 // Unique customerv1 id
	OrgId string

	// TODO: What if we create/manage the Elastic cluster, and just configure a way for the customerv1
	//       to access the Elastic cluster via Kibana, and also allow programmatic query access?
	//       This would add tremendous value as canned querys and visualizations could be created.

	TrackActivity bool // Replicate CDC SLOEntry and Customer activity to Elastic
	Topography    []CustomerLocality
}

// CustomerLocality - contains all the physical origins for customerv1 traffic
type CustomerLocality struct {
	Id         int32
	CustomerId int32
	Locality   Locality
}

type CustomerProcessing struct {
	CustomerId int32
	Locality   []Locality
}

// CustomerMetric - provided tracking of customerv1 metric information
type CustomerMetric struct {
	Id          int32
	CustomerId  int32
	Application string
	RouteName   string
	Locality    Locality
	Name        string
	Value       int32
}
