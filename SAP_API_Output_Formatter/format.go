package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-equipment-master-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToEquipment(raw []byte, l *logger.Logger) (*Equipment, error) {
	pm := &responses.Equipment{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Equipment. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	Equipment := &Equipment{
		Equipment:                     data.Equipment,
		ValidityEndDate:               data.ValidityEndDate,
		ValidityEndTime:               data.ValidityEndTime,
		ValidityStartDate:             data.ValidityStartDate,
		EquipmentName:                 data.EquipmentName,
		EquipmentCategory:             data.EquipmentCategory,
		TechnicalObjectType:           data.TechnicalObjectType,
		GrossWeight:                   data.GrossWeight,
		GrossWeightUnit:               data.GrossWeightUnit,
		SizeOrDimensionText:           data.SizeOrDimensionText,
		InventoryNumber:               data.InventoryNumber,
		OperationStartDate:            data.OperationStartDate,
		AcquisitionValue:              data.AcquisitionValue,
		Currency:                      data.Currency,
		AcquisitionDate:               data.AcquisitionDate,
		AssetManufacturerName:         data.AssetManufacturerName,
		ManufacturerPartTypeName:      data.ManufacturerPartTypeName,
		ManufacturerCountry:           data.ManufacturerCountry,
		ConstructionYear:              data.ConstructionYear,
		ConstructionMonth:             data.ConstructionMonth,
		ManufacturerPartNmbr:          data.ManufacturerPartNmbr,
		ManufacturerSerialNumber:      data.ManufacturerSerialNumber,
		MaintenancePlant:              data.MaintenancePlant,
		AssetLocation:                 data.AssetLocation,
		AssetRoom:                     data.AssetRoom,
		PlantSection:                  data.PlantSection,
		WorkCenter:                    data.WorkCenter,
		WorkCenterPlant:               data.WorkCenterPlant,
		CompanyCode:                   data.CompanyCode,
		BusinessArea:                  data.BusinessArea,
		MasterFixedAsset:              data.MasterFixedAsset,
		FixedAsset:                    data.FixedAsset,
		CostCenter:                    data.CostCenter,
		WBSElementExternalID:          data.WBSElementExternalID,
		SettlementOrder:               data.SettlementOrder,
		MaintenancePlanningPlant:      data.MaintenancePlanningPlant,
		MaintenancePlannerGroup:       data.MaintenancePlannerGroup,
		MainWorkCenter:                data.MainWorkCenter,
		MainWorkCenterPlant:           data.MainWorkCenterPlant,
		CatalogProfile:                data.CatalogProfile,
		FunctionalLocation:            data.FunctionalLocation,
		SuperordinateEquipment:        data.SuperordinateEquipment,
		EquipInstallationPositionNmbr: data.EquipInstallationPositionNmbr,
		TechnicalObjectSortCode:       data.TechnicalObjectSortCode,
		ConstructionMaterial:          data.ConstructionMaterial,
		Material:                      data.Material,
		EquipmentIsAvailable:          data.EquipmentIsAvailable,
		EquipmentIsInstalled:          data.EquipmentIsInstalled,
		EquipIsAllocToSuperiorEquip:   data.EquipIsAllocToSuperiorEquip,
		EquipHasSubOrdinateEquipment:  data.EquipHasSubOrdinateEquipment,
		CreationDate:                  data.CreationDate,
		LastChangeDateTime:            data.LastChangeDateTime,
		EquipmentIsMarkedForDeletion:  data.EquipmentIsMarkedForDeletion,
	}

	return Equipment, nil
}

func ConvertToPartner(raw []byte, l *logger.Logger) (*Partner, error) {
	pm := &responses.Partner{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Partner. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	partner := &Partner{
		Equipment:                  data.Equipment,
		BusinessPartner:            data.BusinessPartner,
		PartnerFunction:            data.PartnerFunction,
		EquipmentPartnerObjectNmbr: data.EquipmentPartnerObjectNmbr,
		Partner:                    data.Partner,
		CreationDate:               data.CreationDate,
	}

	return partner, nil
}
