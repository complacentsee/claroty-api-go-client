package clarotyApiGoClient

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func (request *APIAssetRequest) SetFormatAssetList() {
	format := "asset_list"
	request.Format = &format
}

func (request *APIAssetRequest) SetFormatInsightAssets() {
	format := "insight_assets"
	request.Format = &format
}

func (request *APIAssetRequest) SetFormatRids() {
	format := "rids"
	request.Format = &format
}

func (request *APIAssetRequest) SetPage(page int) {
	request.Page = &page
}

func (request *APIAssetRequest) SetPerPage(perPage int) {
	request.PerPage = &perPage
}

func (request *APIAssetRequest) SetIPv4Exact(ipv4Exact string) {
	request.IPv4Exact = &ipv4Exact
}

func (request *APIAssetRequest) SetIPv6Exact(ipv6Exact string) {
	request.IPv6Exact = &ipv6Exact
}

func (request *APIAssetRequest) SetMacIContains(macIContains string) {
	request.MacIContains = &macIContains
}

func (request *APIAssetRequest) SetVlanExact(vlanExact string) {
	request.VlanExact = &vlanExact
}

func (request *APIAssetRequest) SetAddressExact(addressExact string) {
	request.AddressExact = &addressExact
}

func (request *APIAssetRequest) SetGatewayExact(gatewayExact string) {
	request.GatewayExact = &gatewayExact
}

func (request *APIAssetRequest) SetAssetTypeExact(assetTypeExact string) {
	request.AssetTypeExact = &assetTypeExact
}

func (request *APIAssetRequest) SetHostNameExact(hostNameExact string) {
	request.HostNameExact = &hostNameExact
}

func (request *APIAssetRequest) SetOsExact(osExact string) {
	request.OsExact = &osExact
}

func (request *APIAssetRequest) SetModelIContains(modelIContains string) {
	request.ModelIContains = &modelIContains
}

func (request *APIAssetRequest) SetVendorIContains(vendorIContains string) {
	request.VendorIContains = &vendorIContains
}

func (request *APIAssetRequest) SetStateExact(stateExact string) {
	request.StateExact = &stateExact
}

func (request *APIAssetRequest) SetDomainNamesExact(domainNamesExact string) {
	request.DomainNamesExact = &domainNamesExact
}

func (request *APIAssetRequest) SetFirmwareExact(firmwareExact string) {
	request.FirmwareExact = &firmwareExact
}

func (request *APIAssetRequest) SetSerialExact(serialExact string) {
	request.SerialExact = &serialExact
}

func (request *APIAssetRequest) SetGenericIContains(genericIContains string) {
	request.GenericIContains = &genericIContains
}

func (request *APIAssetRequest) SetDisplayNameIContains(displayNameIContains string) {
	request.DisplayNameIContains = &displayNameIContains
}

func (request *APIAssetRequest) SetCriticalityExact(criticalityExact string) {
	request.CriticalityExact = &criticalityExact
}

func (request *APIAssetRequest) SetOldIpExact(oldIpExact string) {
	request.OldIpExact = &oldIpExact
}

func (request *APIAssetRequest) SetProtocolExact(protocolExact string) {
	request.ProtocolExact = &protocolExact
}

func (request *APIAssetRequest) SetLastSeenExact(lastSeenExact string) {
	request.LastSeenExact = &lastSeenExact
}

func (request *APIAssetRequest) SetQIContains(qIContains string) {
	request.QIContains = &qIContains
}

func (request *APIAssetRequest) SetAlertIdExact(alertIdExact string) {
	request.AlertIdExact = &alertIdExact
}

func (request *APIAssetRequest) SetLastUpdatedGt(lastUpdatedGt string) {
	request.LastUpdatedGt = &lastUpdatedGt
}

func (request *APIAssetRequest) SetBaselineExact(baselineExact string) {
	request.BaselineExact = &baselineExact
}

func (request *APIAssetRequest) SetArpBaselinesExact(arpBaselinesExact string) {
	request.ArpBaselinesExact = &arpBaselinesExact
}

func (request *APIAssetRequest) SetInsightStatusOpen() {
	insightStatusOpen := 0
	request.InsightStatusExact = &insightStatusOpen
}

func (request *APIAssetRequest) SetInsightStatusHidden() {
	insightStatusHidden := 1
	request.InsightStatusExact = &insightStatusHidden
}

func (request *APIAssetRequest) SetInsightStatusCompleted() {
	insightStatusCompleted := 2
	request.InsightStatusExact = &insightStatusCompleted
}

func (request *APIAssetRequest) SetInsightsInsightNameExact(insightsInsightNameExact string) {
	request.InsightsInsightNameExact = &insightsInsightNameExact
}

func (request *APIAssetRequest) SetInsightTimestampGte(insightTimestampGte string) {
	request.InsightTimestampGte = &insightTimestampGte
}

func (request *APIAssetRequest) SetInsightTimestampLte(insightTimestampLte string) {
	request.InsightTimestampLte = &insightTimestampLte
}

func (request *APIAssetRequest) SetBaselineCategoryExact(baselineCategoryExact string) {
	request.BaselineCategoryExact = &baselineCategoryExact
}

func (request *APIAssetRequest) SetBaselineAccessTypeExact(baselineAccessTypeExact string) {
	request.BaselineAccessTypeExact = &baselineAccessTypeExact
}

func (request *APIAssetRequest) SetInsightNameExact(insightNameExact string) {
	request.InsightNameExact = &insightNameExact
}

func (request *APIAssetRequest) SetInsightRowKeyExact(insightRowKeyExact string) {
	request.InsightRowKeyExact = &insightRowKeyExact
}

func (request *APIAssetRequest) SetGhostExact(ghostExact bool) {
	request.GhostExact = &ghostExact
}

func (request *APIAssetRequest) SetTasksExact(tasksExact string) {
	request.TasksExact = &tasksExact
}

func (request *APIAssetRequest) SetActiveQueriesExact(activeQueriesExact string) {
	request.ActiveQueriesExact = &activeQueriesExact
}

func (request *APIAssetRequest) SetSubnetTagExact(subnetTagExact string) {
	request.SubnetTagExact = &subnetTagExact
}

func (request *APIAssetRequest) SetCustomAttributesExact(customAttributesExact string) {
	request.CustomAttributesExact = &customAttributesExact
}

func (request *APIAssetRequest) SetClassTypeExact(classTypeExact string) {
	request.ClassTypeExact = &classTypeExact
}

func (request *APIAssetRequest) SetDomainNameExact(domainNameExact string) {
	request.DomainNameExact = &domainNameExact
}

func (request *APIAssetRequest) SetInvolvedInTagsExact(involvedInTagsExact string) {
	request.InvolvedInTagsExact = &involvedInTagsExact
}

func (request *APIAssetRequest) SetHostedTagsIContains(hostedTagsIContains string) {
	request.HostedTagsIContains = &hostedTagsIContains
}

func (request *APIAssetRequest) SetIdExact(idExact int) {
	request.IdExact = &idExact
}

func (request *APIAssetRequest) SetSiteIdExact(siteIdExact int) {
	request.SiteIdExact = &siteIdExact
}

func (request *APIAssetRequest) SetTimestampExact(timestampExact string) {
	request.TimestampExact = &timestampExact
}

func (request *APIAssetRequest) SetValidExact(validExact bool) {
	request.ValidExact = &validExact
}

func (request *APIAssetRequest) SetParsedExact(parsedExact bool) {
	request.ParsedExact = &parsedExact
}

func (request *APIAssetRequest) SetIsVirtSetSpecialHintExact0ualExact(specialHint int) {
	request.SpecialHintExact = &specialHint
}

func (request *APIAssetRequest) SetSpecialHintExact0() {
	val := 0
	request.SpecialHintExact = &val
}

func (request *APIAssetRequest) SetSpecialHintExact1() {
	val := 1
	request.SpecialHintExact = &val
}

func (request *APIAssetRequest) SetSpecialHintExact2() {
	val := 2
	request.SpecialHintExact = &val
}

func (request *APIAssetRequest) SetSpecialHintExact3() {
	val := 3
	request.SpecialHintExact = &val
}

func (request *APIAssetRequest) SetSpecialHintExact4() {
	val := 4
	request.SpecialHintExact = &val
}

func (request *APIAssetRequest) SetRiskLevelExact(riskLevelExact int) {
	request.RiskLevelExact = &riskLevelExact
}

func (request *APIAssetRequest) SetNetworkIDExact(networkIDExact int) {
	request.NetworkIDExact = &networkIDExact
}

func (request *APIAssetRequest) SetVirtualZoneIDExact(virtualZoneIDExact int) {
	request.VirtualZoneIDEact = &virtualZoneIDExact
}

func (request *APIAssetRequest) SetSubnetIDExact(subnetIDExact int) {
	request.SubnetIDExact = &subnetIDExact
}

func (request *APIAssetRequest) SetPurdueLevelExact(purdueLevelExact string) {
	request.PurdueLevelExact = &purdueLevelExact
}

func (request *APIAssetRequest) GetQueryParameters() string {
	values := url.Values{}

	// Default values on nil pointers
	// these are defaults for the API
	if request.Page == nil {
		request.Page = new(int)
		*request.Page = 1
	}

	if request.PerPage == nil {
		request.PerPage = new(int)
		*request.PerPage = 50
	}

	if request.Format == nil {
		fmt.Println("Query format not specified. If not provided, all asset properties are returned and could affect performance.")
	}

	if request.SiteIdExact == nil {
		fmt.Println("Query site ID not specified. If not provided, all assets are returned and could affect performance.")
	}

	value := reflect.ValueOf(request).Elem()
	st := reflect.TypeOf(request).Elem()

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldName := st.Field(i).Tag.Get("json")

		if fieldValue.IsNil() {
			continue
		}

		switch fieldValue.Interface().(type) {
		case *string:
			values.Set(fieldName, *fieldValue.Interface().(*string))
		case *int:
			values.Set(fieldName, strconv.Itoa(*fieldValue.Interface().(*int)))
		case *bool:
			values.Set(fieldName, strconv.FormatBool(*fieldValue.Interface().(*bool)))
		default:
			// handle other types if needed
		}
	}

	return values.Encode()
}

func (request *APIAssetRequest) GetResponseCount() (int, error) {
	if request.Page != nil && request.PerPage != nil {
		return *request.Page * *request.PerPage, nil
	}
	err := fmt.Errorf("page or PerPage is nil")
	return 0, err
}
