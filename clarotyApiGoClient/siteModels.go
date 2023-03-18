package clarotyApiGoClient

type APISiteRequest struct {
	Rid__exact *string `json:"rid__exact"`
	Sort       *string `json:"sort"`
	Format     *string `json:"format"`
	Page       *int    `json:"page"`
	Per_page   *int    `json:"per_page"`
}

type APISiteResponse struct {
	CountFiltered int    `json:"count_filtered"`
	CountTotal    int    `json:"count_total"`
	CountInPage   int    `json:"count_in_page"`
	Sites         []Site `json:"objects"`
}

type Site struct {
	Address                      string `json:"address"`
	AssetsCount                  int    `json:"assets_count"`
	AutoTransitionToOperational  bool   `json:"auto_transition_to_operational"`
	Description                  string `json:"description"`
	ID                           int    `json:"id"`
	ImageURL                     string `json:"image_url"`
	InstalledDate                string `json:"installed_date"`
	IntegrityAlertsSeverity      int    `json:"integrity_alerts_severity"`
	IntegrityAlertsSeverityLabel string `json:"integrity_alerts_severity__"`
	IsConnected                  bool   `json:"is_connected"`
	Name                         string `json:"name"`
	Score                        int    `json:"score"`
	SecurityAlertsSeverity       int    `json:"security_alerts_severity"`
	SecurityAlertsSeverityLabel  string `json:"security_alerts_severity__"`
	TrainingMode                 bool   `json:"training_mode"`
	UnresolvedIntegrityAlerts    int    `json:"unresolved_integrity_alerts_count"`
	UnresolvedSecurityAlerts     int    `json:"unresolved_security_alerts_count"`
	Version                      string `json:"version"`
}
