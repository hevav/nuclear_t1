package pipes

import (
	"hackathon/config"
	"hackathon/db"
	"hackathon/types"
	"math"
)

func UpdateAggSum(tx *types.Transaction) (*types.TransactionPartyGeoZoneVector, error) {
	txPartyGeoZoneVectors, err := db.GetGeoZonesByParty(tx.FromID)
	if err != nil {
		return nil, err
	}
	var nearestVector *types.TransactionPartyGeoZoneVector
	var minDistance float64 = math.MaxFloat64
	for _, vector := range txPartyGeoZoneVectors {
		distance := haversineDistance(tx.GeoLat, tx.GeoLon, vector.GeoLat, vector.GeoLon)
		if distance < minDistance && distance <= float64(vector.GeoRadius) {
			minDistance = distance
			nearestVector = &vector
		}
	}

	// TODO: вообще если несколько геозон, то надо их объединять вместе, это не сложно, но случается очень редко

	if nearestVector != nil {
		points, err := db.GetGeoZonePoints(nearestVector.ID)
		if err != nil {
			return nil, err
		}

		newLat, newLon := meanPoint(append(points, tx.GeoLat, tx.GeoLon))
		nearestVector.GeoRadius += uint64(haversineDistance(nearestVector.GeoLat, nearestVector.GeoLon, newLat, newLon))
		nearestVector.GeoLat = newLat
		nearestVector.GeoLat = newLon
		nearestVector.AggSum.Add(tx.Amount)

		return nearestVector, db.UpdateGeoZoneVector(nearestVector)
	}

	newVector := &types.TransactionPartyGeoZoneVector{
		TransactionPartyID: tx.FromID,
		GeoLat:             tx.GeoLat,
		GeoLon:             tx.GeoLon,
		GeoRadius:          config.Settings.Specific.GeoZone.DefaultRadiusMeters,
		AggSum: types.AggSum{
			Sum:         uint64(tx.Amount),
			AbsoluteSum: uint64(math.Abs(float64(tx.Amount))),
			Count:       1,
		},
	}

	return newVector, db.CreateNewGeoZoneVector(tx, newVector)
}

func meanPoint(points []float64) (float64, float64) {
	meanLat := points[0]
	meanLon := points[1]

	distances := make([]float64, len(points))
	for j := 0; j < len(points); j += 2 {
		distances[j] = 2 * math.Asin(math.Sin(math.Abs(points[j]-meanLat)/2)/math.Sin(math.Pi/2))
	}

	meanLat = 0
	meanLon = 0
	for j := 0; j < len(points); j += 2 {
		meanLat += points[j] * math.Cos(distances[j]/2)
		meanLon += points[j+1] * math.Cos(meanLat) / math.Sin(distances[j])
	}
	meanLat /= float64(len(points))
	meanLon /= float64(len(points))

	return meanLat, meanLon
}

func haversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000

	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c

	return distance
}

type Geozone struct{}

func (Geozone) Proceed(tx *types.Transaction) (float64, error) {
	geozone, err := UpdateAggSum(tx)
	if err != nil {
		return 0, err
	}

	return geozone.AggSum.GetScore(config.Weights.AggSumWeights.GeoZone), nil
}
