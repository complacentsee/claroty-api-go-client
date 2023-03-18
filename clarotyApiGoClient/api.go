package clarotyApiGoClient

func NewClarotyAPI(config *APIConfiguration) ClarotyAPI {

	if config.MaxConcurrentRequests == nil {
		maxConcurrentRequests := 20
		config.MaxConcurrentRequests = &maxConcurrentRequests
	}

	semaphore := make(chan struct{}, *config.MaxConcurrentRequests)

	return ClarotyAPI{
		configuration:  *config,
		authentication: nil,
		semaphore:      &semaphore,
	}
}
