package main

import (
	"encoding/json"
	"strconv"

	"github.com/alexmorten/smag-mvp/config"
	"github.com/alexmorten/smag-mvp/elastic"
	"github.com/alexmorten/smag-mvp/elastic/indexer"
	esModels "github.com/alexmorten/smag-mvp/elastic/models"
	"github.com/alexmorten/smag-mvp/insta/models"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/kafka/changestream"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	i := indexer.New(elastic.FacesIndex, elastic.FacesIndexMapping, kafka.TopicNamePGFaces, "face_indexer", conf, indexFace)

	service.CloseOnSignal(i)
	waitUntilDone := i.Start()

	waitUntilDone()
}

func indexFace(m *changestream.ChangeMessage) (*indexer.BulkIndexDoc, error) {

	switch m.Payload.Op {
	case "r", "u", "c":
		break
	default:
		return nil, nil
	}

	face := &models.FaceData{}
	err := json.Unmarshal(m.Payload.After, face)
	if err != nil {
		return nil, err
	}

	return createBulkIndexOperation(face)
}

func createBulkIndexOperation(face *models.FaceData) (*indexer.BulkIndexDoc, error) {
	bulkOperation := `{ "index": {}  }` + "\n"

	doc, err := esModels.FaceDocFromFaceData(face)
	if err != nil {
		return nil, err
	}

	docJSON, err := json.Marshal(doc)

	if err != nil {
		return nil, err
	}

	docJSON = append(docJSON, "\n"...)

	bulkUpsertBody := bulkOperation + string(docJSON)

	return &indexer.BulkIndexDoc{DocumentId: strconv.Itoa(int(face.ID)), BulkOperation: bulkUpsertBody}, err

}
