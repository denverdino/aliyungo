package pvtz

import (
	"testing"
)


func TestDescribeRegions(t *testing.T) {
	client := NewTestClient()

	regions, err := client.DescribeRegions()

	t.Logf("regions: %v, %v", regions, err)
}

func TestAddZone(t *testing.T) {
	client := NewTestClient()

	response, err := client.AddZone(&AddZoneArgs{
		ZoneName:"demo.com",
	})

	t.Logf("AddZone: %++v, %v", response, err)
	TestDescribeZones(t)

	zoneId := response.ZoneId

	testDescribeZoneRecords(t, zoneId)

	testDeleteZone(t, zoneId)
}

func testDeleteZone(t *testing.T, zoneId string) {
	client := NewTestClient()

	err := client.DeleteZone(&DeleteZoneArgs{
		ZoneId: zoneId,
	})
	t.Logf("DeleteZone: %v", err)
}

func TestDescribeZones(t *testing.T) {
	client := NewTestClient()

	zones, err := client.DescribeZones(&DescribeZonesArgs{})

	t.Logf("zones: %v, %v", zones, err)
}

func testDescribeZoneRecords(t *testing.T, zoneId string) {
	client := NewTestClient()

	response, err := client.AddZoneRecord(&AddZoneRecordArgs{
		ZoneId: zoneId,
		Rr: "www",
		Type: "A",
		Ttl: 60,
		Value: "1.1.1.1",
	})

	t.Logf("AddZoneRecord: %v, %v", response, err)

	if err != nil {
		return
	}

	recordId := response.RecordId
	records, err := client.DescribeZoneRecords(&DescribeZoneRecordsArgs{
		ZoneId: zoneId,
	})

	t.Logf("records: %v, %v", records, err)


	err = client.DeleteZoneRecord(&DeleteZoneRecordArgs{
		RecordId: recordId,
	})
	t.Logf("DeleteZoneRecord: %v", err)
}