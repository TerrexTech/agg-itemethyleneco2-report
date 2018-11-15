package report

import (
	util "github.com/TerrexTech/go-commonutils/commonutil"
	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

type MetricSoldItem struct {
	ID            objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ItemID        uuuid.UUID        `bson:"itemID,omitempty" json:"itemID,omitempty"`
	SaleID        uuuid.UUID        `bson:"saleID,omitempty" json:"saleID,omitempty"`
	SKU           string            `bson:"sku,omitempty" json:"sku,omitempty"`
	Name          string            `bson:"name,omitempty" json:"name,omitempty"`
	Lot           string            `bson:"lot,omitempty" json:"lot,omitempty"`
	SoldWeight    float64           `bson:"soldWeight,omitempty" json:"soldWeight,omitempty"`
	WasteWeight   float64           `bson:"wasteWeight,omitempty" json:"wasteWeight,omitempty"`
	DonateWeight  float64           `bson:"donateWeight,omitempty" json:"donateWeight,omitempty"`
	TotalWeight   float64           `bson:"totalWeight,omitempty" json:"totalWeight,omitempty"`
	Timestamp     int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Ethylene      float64           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	CarbonDioxide float64           `bson:"carbonDioxide,omitempty" json:"carbonDioxide,omitempty"`
}

type SoldItemParams struct {
	Timestamp *Comparator `json:"timestamp,omitempty"`
}

func (s MetricSoldItem) MarshalBSON() ([]byte, error) {
	si := map[string]interface{}{
		"itemID":        s.ItemID.String(),
		"saleID":        s.SaleID.String(),
		"lot":           s.Lot,
		"name":          s.Name,
		"sku":           s.SKU,
		"soldWeight":    s.SoldWeight,
		"donateWeight":  s.DonateWeight,
		"wasteWeight":   s.WasteWeight,
		"timestamp":     s.Timestamp,
		"totalWeight":   s.TotalWeight,
		"ethylene":      s.Ethylene,
		"carbonDioxide": s.CarbonDioxide,
	}

	if s.ID != objectid.NilObjectID {
		si["_id"] = s.ID
	}
	return bson.Marshal(si)
}

func (s MetricSoldItem) UnmarshalBSON(in []byte) error {
	m := make(map[string]interface{})
	err := bson.Unmarshal(in, m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	err = s.unmarshalFromMap(m)
	return err
}

func (s MetricSoldItem) unmarshalFromMap(m map[string]interface{}) error {
	var err error
	var assertOK bool

	if m["_id"] != nil {
		s.ID, assertOK = m["_id"].(objectid.ObjectID)
		if !assertOK {
			s.ID, err = objectid.FromHex(m["_id"].(string))
			if err != nil {
				err = errors.Wrap(err, "Error while asserting ObjectID")
				return err
			}
		}
	}

	if m["itemID"] != nil {
		s.ItemID, err = uuuid.FromString(m["itemID"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error while asserting ItemID")
			return err
		}
	}

	if m["saleID"] != nil {
		s.SaleID, err = uuuid.FromString(m["saleID"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error while asserting DeviceID")
			return err
		}
	}

	if m["lot"] != nil {
		s.Lot, assertOK = m["lot"].(string)
		if !assertOK {
			return errors.New("Error while asserting Lot")
		}
	}

	if m["name"] != nil {
		s.Name, assertOK = m["name"].(string)
		if !assertOK {
			return errors.New("Error while asserting Name")
		}
	}

	if m["sku"] != nil {
		s.SKU, assertOK = m["sku"].(string)
		if !assertOK {
			return errors.New("Error while asserting Sku")
		}
	}
	if m["soldWeight"] != nil {
		s.SoldWeight, err = util.AssertFloat64(m["soldWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting SoldWeight")
			return err
		}
	}
	if m["donateWeight"] != nil {
		s.DonateWeight, err = util.AssertFloat64(m["donateWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting DonateWeight")
			return err
		}
	}
	if m["wasteWeight"] != nil {
		s.WasteWeight, err = util.AssertFloat64(m["wasteWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting WasteWeight")
			return err
		}
	}
	if m["timestamp"] != nil {
		s.Timestamp, err = util.AssertInt64(m["timestamp"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting Timestamp")
			return err
		}
	}
	if m["totalWeight"] != nil {
		s.TotalWeight, err = util.AssertFloat64(m["totalWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting TotalWeight")
			return err
		}
	}
	if m["ethylene"] != nil {
		s.Ethylene, err = util.AssertFloat64(m["ethylene"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting Ethylene")
			return err
		}
	}
	if m["carbonDioxide"] != nil {
		s.CarbonDioxide, err = util.AssertFloat64(m["carbonDioxide"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting carbonDioxide")
			return err
		}
	}
	return nil
}
