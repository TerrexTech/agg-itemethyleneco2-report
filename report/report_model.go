package report

import (
	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

type SoldReport struct {
	ID           objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ReportID     uuuid.UUID        `bson:"reportID,omitempty" json:"reportID,omitempty"`
	SearchQuery  SoldItemParams    `bson:"searchQuery,omitempty" json:"searchQuery,omitempty"`
	ReportResult []ReportResult    `bson:"reportResult,omitempty" json:"reportResult,omitempty"`
}

type SoldReportBSON struct {
	ID           objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ReportID     string            `bson:"reportID,omitempty" json:"reportID,omitempty"`
	SearchQuery  SoldItemParams    `bson:"searchQuery,omitempty" json:"searchQuery,omitempty"`
	ReportResult []ReportResult    `bson:"reportResult,omitempty" json:"reportResult,omitempty"`
}

type ReportResult struct {
	SKU           string  `bson:"sku,omitempty" json:"sku,omitempty"`
	Name          string  `bson:"name,omitempty" json:"name,omitempty"`
	Ethylene      float64 `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	CarbonDioxide float64 `bson:"carbonDioxide,omitempty" json:"carbonDioxide,omitempty"`
	SoldWeight    float64 `bson:"soldWeight,omitempty" json:"soldWeight,omitempty"`
	WasteWeight   float64 `bson:"wasteWeight,omitempty" json:"wasteWeight,omitempty"`
	DonateWeight  float64 `bson:"donateWeight,omitempty" json:"donateWeight,omitempty"`
	TotalWeight   float64 `bson:"totalWeight,omitempty" json:"totalWeight,omitempty"`
}

func (s SoldReport) MarshalBSON() ([]byte, error) {
	sm := map[string]interface{}{
		"reportid":     s.ReportID.String(),
		"searchquery":  s.SearchQuery,
		"reportresult": s.ReportResult,
	}
	if s.ID != objectid.NilObjectID {
		sm["_id"] = s.ID
	}
	// Do more stuff

	if s.ReportID != (uuuid.UUID{}) {
		sm["reportID"] = s.ReportID.String()
	}

	return bson.Marshal(sm)
}

func (s *SoldReport) UnmarshalBSON(in []byte) error {
	sb := &SoldReportBSON{}
	err := bson.Unmarshal(in, sb)
	if err != nil {
		err = errors.Wrap(err, "UnmarshalBSON Error")
		return err
	}

	if sb.ID != objectid.NilObjectID {
		s.ID = sb.ID
	}
	reportID, err := uuuid.FromString(sb.ReportID)
	if err != nil {
		err = errors.Wrap(err, "UnmarshalBSON Error: Error parsing SaleID")
	}
	s.ReportID = reportID

	if s.ReportResult == nil {
		s.ReportResult = make([]ReportResult, 0)
	}
	for _, v := range sb.ReportResult {
		s.ReportResult = append(s.ReportResult, ReportResult{
			SKU:           v.SKU,
			Name:          v.Name,
			Ethylene:      v.Ethylene,
			CarbonDioxide: v.CarbonDioxide,
			SoldWeight:    v.SoldWeight,
			DonateWeight:  v.DonateWeight,
			WasteWeight:   v.WasteWeight,
			TotalWeight:   v.TotalWeight,
		})
	}
	return nil
}
