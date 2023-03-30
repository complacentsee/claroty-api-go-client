// Copyright (c) 2023, Adam Traeger
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package clarotyApiGoClient

import "time"

type APIAssetRequest struct {
	Format                   *string `json:"format"`
	Page                     *int    `json:"page"`
	PerPage                  *int    `json:"per_page"`
	IPv4Exact                *string `json:"ipv4__exact"`
	IPv6Exact                *string `json:"ipv6__exact"`
	MacIContains             *string `json:"mac__icontains"`
	VlanExact                *string `json:"vlan__exact"`
	AddressExact             *string `json:"address__exact"`
	GatewayExact             *string `json:"gateway__exact"`
	AssetTypeExact           *string `json:"asset_type__exact"`
	HostNameExact            *string `json:"host_name__exact"`
	OsExact                  *string `json:"os__exact"`
	ModelIContains           *string `json:"model__icontains"`
	VendorIContains          *string `json:"vendor__icontains"`
	StateExact               *string `json:"state__exact"`
	DomainNamesExact         *string `json:"domain_names__exact"`
	FirmwareExact            *string `json:"firmware__exact"`
	SerialExact              *string `json:"serial__exact"`
	GenericIContains         *string `json:"generic__icontains"`
	DisplayNameIContains     *string `json:"display_name__icontains"`
	CriticalityExact         *string `json:"criticality__exact"`
	OldIpExact               *string `json:"old_ip__exact"`
	ProtocolExact            *string `json:"protocol__exact"`
	LastSeenExact            *string `json:"last_seen__exact"`
	QIContains               *string `json:"q__icontains"`
	AlertIdExact             *string `json:"alert_id__exact"`
	LastUpdatedGt            *string `json:"last_updated__gt"`
	BaselineExact            *string `json:"baseline__exact"`
	ArpBaselinesExact        *string `json:"arp_baselines__exact"`
	InsightStatusExact       *int    `json:"insight_status__exact"`
	InsightsInsightNameExact *string `json:"insights_insight_name__exact"`
	InsightTimestampGte      *string `json:"insight_timestamp__gte"`
	InsightTimestampLte      *string `json:"insight_timestamp__lte"`
	BaselineCategoryExact    *string `json:"baseline_category__exact"`
	BaselineAccessTypeExact  *string `json:"baseline_access_type__exact"`
	InsightNameExact         *string `json:"insight_name__exact"`
	InsightRowKeyExact       *string `json:"insight_row_key__exact"`
	GhostExact               *bool   `json:"ghost__exact"`
	TasksExact               *string `json:"tasks__exact"`
	ActiveQueriesExact       *string `json:"active_queries__exact"`
	SubnetTagExact           *string `json:"subnet_tag__exact"`
	CustomAttributesExact    *string `json:"custom_attributes__exact"`
	ClassTypeExact           *string `json:"class_type__exact"`
	DomainNameExact          *string `json:"domain_name__exact"`
	InvolvedInTagsExact      *string `json:"involved_in_tags__exact"`
	HostedTagsIContains      *string `json:"hosted_tags__icontains"`
	IdExact                  *int    `json:"id__exact"`
	SiteIdExact              *int    `json:"site_id__exact"`
	TimestampExact           *string `json:"timestamp__exact"`
	ApprovedExact            *bool   `json:"approved__exact"`
	ValidExact               *bool   `json:"valid__exact"`
	ParsedExact              *bool   `json:"parsed__exact"`
	SpecialHintExact         *int    `json:"special_hint__exact"`
	RiskLevelExact           *int    `json:"risk_level__exact"`
	NetworkIDExact           *int    `json:"network_id__exact"`
	VirtualZoneIDEact        *int    `json:"virtual_zone_id__exact"`
	SubnetIDExact            *int    `json:"subnet_id__exact"`
	PurdueLevelExact         *string `json:"purdue_level__exact"`
}

type APIAssetResponse struct {
	CountFiltered int     `json:"count_filtered"`
	CountInPage   int     `json:"count_in_page"`
	CountTotal    int     `json:"count_total"`
	Assets        []Asset `json:"objects"`
}

type Asset struct {
	ActiveQueriesNames     []string       `json:"active_queries_names"`
	ActiveTasksNames       []string       `json:"active_tasks_names"`
	Approved               bool           `json:"approved"`
	AssetType              int            `json:"asset_type"`
	AssetType_             string         `json:"asset_type__"`
	Children               []Children     `json:"children"`
	ClassType              string         `json:"class_type"`
	CodeSections           []CodeSections `json:"code_sections"`
	Criticality            int            `json:"criticality"`
	Criticality_           string         `json:"criticality__"`
	CustomAttributes       []Attribute    `json:"custom_attributes"`
	CustomInformations     []Information  `json:"custom_informations"`
	DefaultGateway         interface{}    `json:"default_gateway"`
	DisplayName            string         `json:"display_name"`
	DomainWorkgroup        interface{}    `json:"domain_workgroup"`
	EdgeID                 interface{}    `json:"edge_id"`
	EdgeLastRun            interface{}    `json:"edge_last_run"`
	Firmware               string         `json:"firmware"`
	FirstSeen              string         `json:"first_seen"`
	Ghost                  bool           `json:"ghost"`
	ID                     int            `json:"id"`
	InsightNames           []string       `json:"insight_names"`
	InstalledAntivirus     interface{}    `json:"installed_antivirus"`
	InstalledProgramsCount int            `json:"installed_programs_count"`
	Ipv4                   []string       `json:"ipv4"`
	LastSeen               string         `json:"last_seen"`
	LastUpdated            string         `json:"last_updated"`
	Mac                    []string       `json:"mac"`
	Model                  string         `json:"model"`
	Name                   string         `json:"name"`
	Network                Network        `json:"network"`
	NetworkID              int            `json:"network_id"`
	NumAlerts              int            `json:"num_alerts"`
	OsArchitecture         interface{}    `json:"os_architecture"`
	OsBuild                interface{}    `json:"os_build"`
	OsServicePack          interface{}    `json:"os_service_pack"`
	Parsed                 bool           `json:"parsed"`
	PatchCount             int            `json:"patch_count"`
	ProjectParsed          interface{}    `json:"project_parsed"`
	Protocol               []string       `json:"protocol"`
	PurdueLevel            float64        `json:"purdue_level"`
	ResourceID             string         `json:"resource_id"`
	RiskLevel              int            `json:"risk_level"`
	SerialNumber           string         `json:"serial_number"`
	SiteID                 int            `json:"site_id"`
	SiteName               string         `json:"site_name"`
	SpecialHint            int            `json:"special_hint"`
	SpecialHint_           string         `json:"special_hint__"`
	State                  interface{}    `json:"state"`
	Subnet                 Subnet         `json:"subnet"`
	SubnetID               int            `json:"subnet_id"`
	SubnetTag              string         `json:"subnet_tag"`
	SubnetType             int            `json:"subnet_type"`
	Timestamp              string         `json:"timestamp"`
	UsbDevicesCount        int            `json:"usb_devices_count"`
	Valid                  bool           `json:"valid"`
	VirtualZoneID          int            `json:"virtual_zone_id"`
	VirtualZoneName        string         `json:"virtual_zone_name"`
	VLAN                   []string       `json:"vlan"`
	PlcSlots               interface{}    `json:"plc_slots"`
	//PlcSlots               []PlcSlots     `json:"plc_slots"`
}

type Attribute struct {
	AssetID    int      `json:"asset_id"`
	Category   Category `json:"category"`
	ID         int      `json:"id"`
	ResourceID string   `json:"resource_id"`
	SiteID     int      `json:"site_id"`
	Value      string   `json:"value"`
}

type Category struct {
	Description string `json:"description"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceID  string `json:"resource_id"`
	SiteID      int    `json:"site_id"`
}

type Children struct {
	ActiveQueriesNames []any     `json:"active_queries_names"`
	ActiveTasksNames   []any     `json:"active_tasks_names"`
	Address            []string  `json:"address"`
	Approved           bool      `json:"approved"`
	AssetType          int       `json:"asset_type"`
	AssetType0         string    `json:"asset_type__"`
	ClassType          string    `json:"class_type"`
	Criticality        int       `json:"criticality"`
	Criticality0       string    `json:"criticality__"`
	CustomInformations []any     `json:"custom_informations"`
	DefaultGateway     any       `json:"default_gateway"`
	DisplayName        any       `json:"display_name"`
	DomainWorkgroup    any       `json:"domain_workgroup"`
	EdgeID             any       `json:"edge_id"`
	EdgeLastRun        any       `json:"edge_last_run"`
	Firmware           string    `json:"firmware,omitempty"`
	FirstSeen          time.Time `json:"first_seen"`
	Ghost              bool      `json:"ghost"`
	ID                 int       `json:"id"`
	InstalledAntivirus any       `json:"installed_antivirus"`
	LastSeen           time.Time `json:"last_seen"`
	LastUpdated        time.Time `json:"last_updated"`
	Model              string    `json:"model,omitempty"`
	Name               string    `json:"name"`
	Network            Network   `json:"network"`
	NetworkID          int       `json:"network_id"`
	OsArchitecture     any       `json:"os_architecture"`
	OsBuild            any       `json:"os_build"`
	OsServicePack      any       `json:"os_service_pack"`
	Parsed             bool      `json:"parsed"`
	ProjectParsed      any       `json:"project_parsed"`
	ResourceID         string    `json:"resource_id"`
	RiskLevel          int       `json:"risk_level"`
	SerialNumber       string    `json:"serial_number,omitempty"`
	SiteID             int       `json:"site_id"`
	SiteName           string    `json:"site_name"`
	SpecialHint        int       `json:"special_hint"`
	SpecialHint0       string    `json:"special_hint__"`
	State              any       `json:"state"`
	Subnet             Subnet    `json:"subnet"`
	SubnetID           int       `json:"subnet_id"`
	SubnetTag          string    `json:"subnet_tag"`
	SubnetType         int       `json:"subnet_type"`
	Timestamp          time.Time `json:"timestamp"`
	Vendor             string    `json:"vendor,omitempty"`
	VirtualZoneID      int       `json:"virtual_zone_id"`
	VirtualZoneName    string    `json:"virtual_zone_name"`
}

type Information struct {
	Category   int    `json:"category"`
	DisplayKey string `json:"display_key"`
	Key        string `json:"key"`
	Priority   int    `json:"priority"`
	Type       int    `json:"type"`
	Value      string `json:"val"`
}

type Network struct {
	ID         int
	Name       string
	ResourceID string
	SiteID     int
}

type Subnet struct {
	Name string `json:"name"`
}

type PLCInformation struct {
	Address         string `json:"address"`
	Description     string `json:"description"`
	FirmwareVersion string `json:"firmware_version"`
	InformationType int    `json:"information_type"`
	Name            string `json:"name"`
	OrderNumber     string `json:"order_number"`
	Priority        int    `json:"priority"`
	Product         string `json:"product"`
	SerialNumber    string `json:"serial_number"`
	Vendor          string `json:"vendor"`
}

type Value struct {
	PLCInformation PLCInformation `json:"PLCInformation"`
}

type PLCSlotInformation struct {
	Description     string `json:"description"`
	InformationType int    `json:"information_type"`
	Priority        int    `json:"priority"`
	Slot            int    `json:"slot"`
	Value           Value  `json:"value"`
}

type PlcSlots struct {
	PLCSlotInformation PLCSlotInformation `json:"PLCSlotInformation"`
}

type CodeSections struct {
	Filename string `json:"filename"`
	Rid      string `json:"rid"`
	Type     string `json:"type"`
}
