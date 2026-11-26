package astroServer

import (
	"fmt"
	"slices"
	"strings"
	"time"

	astroV1 "github.com/binuud/astro/gen/go/v1/astro"
	"github.com/binuud/astro/pkg/astro"
	datalayer "github.com/binuud/astro/pkg/dataLayer"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	CHARTS_DIRECTORY = "charts"
)

func CreateRecord(in *astroV1.UserCreateRequest) (*astroV1.AstroObject, error) {

	metaData := &astroV1.Metadata{
		Uid:  uuid.New().String(),
		Name: in.Metadata.Name,
		// Owner, err = authServer.GetUserUuidFromContext(ctx)
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
		ChartType: astroV1.Metadata_Global,
	}

	// logrus.Info("Received Create: ", in)

	chart, err := astro.NewAstroChart(&astroV1.BirthData{
		Tob: &astroV1.TobData{
			D:    in.BirthData.Tob.D,
			M:    in.BirthData.Tob.M,
			Y:    in.BirthData.Tob.Y,
			Hour: in.BirthData.Tob.Hour,
			Min:  in.BirthData.Tob.Min,
			Sec:  in.BirthData.Tob.Sec,
			Zone: in.BirthData.Tob.Zone,
		},
		Pob: &astroV1.PobData{
			City:    in.BirthData.Pob.City,
			State:   in.BirthData.Pob.State,
			Country: in.BirthData.Pob.Country,
			Lat:     in.BirthData.Pob.Lat,
			Long:    in.BirthData.Pob.Long,
		},
	})
	if err != nil {
		logrus.Errorf("Cannot prepare chart err(%v)", err)
	}
	// chart.PrintChartTable()

	astroObject := &astroV1.AstroObject{
		Metadata:  metaData,
		BirthData: chart.BirthData,
		Planets: &astroV1.Planets{
			PlanetsMap: chart.Planets.PlanetsMap,
		},
		Dashas: chart.Dashas,
	}

	err = SaveAstroChart(astroObject)
	if err != nil {
		return nil, err
	}

	return astroObject, nil
}

func UpdateRecord(in *astroV1.UserUpdateRequest) (*astroV1.AstroObject, error) {

	chart, err := astro.NewAstroChart(in.BirthData)
	if err != nil {
		return nil, err
	}

	err = SaveAstroChart(chart.AstroObject)
	if err != nil {
		return nil, err
	}

	return &astroV1.AstroObject{
		BirthData: chart.BirthData,
		Planets: &astroV1.Planets{
			PlanetsMap: chart.Planets.PlanetsMap,
		},
	}, nil
}

func UpSertRecord(in *astroV1.UserCreateRequest, filter bson.D) (*astroV1.AstroObject, error) {

	metaData := &astroV1.Metadata{
		Uid:  uuid.New().String(),
		Name: in.Metadata.Name,
		// Owner, err = authServer.GetUserUuidFromContext(ctx)
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
		ChartType: astroV1.Metadata_Global,
	}

	chart, err := astro.NewAstroChart(in.BirthData)
	if err != nil {
		return nil, err
	}

	astroObject := chart.AstroObject
	astroObject.Metadata = metaData

	err = SaveAstroChart(astroObject)
	if err != nil {
		return nil, err
	}

	return astroObject, nil
}

func CreateAstroObjectForTime(name string, t time.Time) (*astroV1.AstroObject, error) {

	birthData := CreateBirthDataForTime(t)
	return CreateAstroObject(name, birthData)

}

func CreateAstroObject(name string, birthData *astroV1.BirthData) (*astroV1.AstroObject, error) {
	metaData := &astroV1.Metadata{
		Uid:  uuid.New().String(),
		Name: name,
		// Owner, err = authServer.GetUserUuidFromContext(ctx)
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
		ChartType: astroV1.Metadata_Global,
	}

	chart, err := astro.NewAstroChart(birthData)
	if err != nil {
		return nil, err
	}

	astroObject := chart.AstroObject
	astroObject.Metadata = metaData

	return astroObject, nil
}

func CreateBirthDataForTime(t time.Time) *astroV1.BirthData {
	return &astroV1.BirthData{
		Tob: &astroV1.TobData{
			D:    int32(t.UTC().Day()),
			M:    int32(t.UTC().Month()),
			Y:    int32(t.UTC().Year()),
			Hour: int32(t.UTC().Hour()),
			Min:  int32(t.UTC().Minute()),
			Sec:  0,
			Zone: "UTC",
		},
		Pob: &astroV1.PobData{
			City: "Chennai",

			Lat:  13.079999923706055,
			Long: 80.69999694824219,
		},
	}
}

func ListRecords(pagination *astroV1.Pagination) ([]*astroV1.AstroObject, error) {

	charts := []*astroV1.AstroObject{}

	folderPath := fmt.Sprintf("%s/%s", datalayer.GetStoreDirectory(), CHARTS_DIRECTORY)
	files, err := datalayer.ListFolder(folderPath, ".yaml")
	if err != nil {
		return nil, err
	}
	slices.Sort(files)

	for _, item := range files {
		result := strings.Split(item, "--")
		if len(result) == 2 {
			astroObj := &astroV1.AstroObject{
				Metadata: &astroV1.Metadata{
					Uid:  result[1],
					Name: result[0],
				},
			}
			charts = append(charts, astroObj)
		}
	}
	return charts, nil
}

func SaveAstroChart(astroObject *astroV1.AstroObject) error {
	fileName := fmt.Sprintf("%s--%s.yaml", astroObject.Metadata.Name, astroObject.Metadata.Uid)
	filePath := datalayer.GetFileAbsolutePath(CHARTS_DIRECTORY, fileName)
	return datalayer.SaveProtoYAML(filePath, astroObject)
}

func GetRecord(uuid string, name string) (*astroV1.AstroObject, error) {
	fileName := fmt.Sprintf("%s--%s.yaml", name, uuid)
	filePath := datalayer.GetFileAbsolutePath(CHARTS_DIRECTORY, fileName)
	astroObject := &astroV1.AstroObject{}
	err := datalayer.ReadProtoYAML(filePath, astroObject)
	return astroObject, err
}
