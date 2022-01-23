package requests

type Partner struct {
	Equipment                   string `json:"Equipment"`
	BusinessPartner             string `json:"BusinessPartner"`
	PartnerFunction             string `json:"PartnerFunction"`
	EquipmentPartnerObjectNmbr  string `json:"EquipmentPartnerObjectNmbr"`
	Partner                    *string `json:"Partner"`
	CreationDate               *string `json:"CreationDate"`
}
