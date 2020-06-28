package services

import (
	"emotion-based-memes/structs"
	"io"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/preview/cognitiveservices/face"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/imagesearch"
	"github.com/Azure/go-autorest/autorest"
	"golang.org/x/net/context"
)

type AzureService struct {
	faceClient face.Client
	bingClient imagesearch.ImagesClient
}

func NewAzureService() *AzureService {
	authFace := autorest.NewCognitiveServicesAuthorizer(os.Getenv("FACE_KEY"))
	authBing := autorest.NewCognitiveServicesAuthorizer(os.Getenv("BING_KEY"))

	faceClient := face.NewClient(os.Getenv("FACE_ENDPOINT"))
	faceClient.Authorizer = authFace

	bingClient := imagesearch.ImagesClient{
		BaseClient: imagesearch.NewWithoutDefaults(os.Getenv("BING_ENDPOINT")),
	}
	bingClient.Authorizer = authBing

	return &AzureService{
		faceClient: faceClient,
		bingClient: bingClient,
	}
}

func (as *AzureService) DetectEmotion(image io.ReadCloser) (*face.DetectedFace, error) {
	returnFaceID := true
	returnFaceLandmarks := true
	returnRecognitionModel := true
	faceAttributeTypes := []face.AttributeType{face.AttributeTypeEmotion}

	result, err := as.faceClient.DetectWithStream(
		context.Background(),
		image,
		&returnFaceID,
		&returnFaceLandmarks,
		faceAttributeTypes,
		face.Recognition01,
		&returnRecognitionModel,
		face.Detection01,
	)

	if err != nil {
		log.Println("error on detect image", err)
		return nil, err
	}

	var resp face.DetectedFace
	for _, item := range *result.Value {
		resp = item
	}

	return &resp, nil
}

func (as *AzureService) SearchImages(search string) ([]structs.BingImage, error) {

	// this sdk client probally was written by a csharp dev
	images, err := as.bingClient.Search(
		context.Background(),         // context
		search+" meme",               // query keyword
		"",                           // Accept-Language header
		"",                           // User-Agent header
		"",                           // X-MSEdge-ClientID header
		"",                           // X-MSEdge-ClientIP header
		"",                           // X-Search-Location header
		imagesearch.All,              // image aspect
		imagesearch.ColorOnly,        // image color
		"",                           // country code
		nil,                          // count
		"All",                        // freshness
		nil,                          // height
		"",                           // ID
		imagesearch.ImageContent(""), // image content
		imagesearch.Photo,            // image type
		imagesearch.ImageLicenseAll,  // image license
		"",                           // market
		nil,                          // max file size
		nil,                          // max height
		nil,                          // max width
		nil,                          // min file size
		nil,                          // min height
		nil,                          // min width
		nil,                          // offset
		imagesearch.Strict,           // safe search
		imagesearch.ImageSizeAll,     // image size
		"",                           // set lang
		nil,                          // width
	)

	if err != nil {
		log.Println("err on search images", err)
	}

	imgs := []structs.BingImage{}

	for _, image := range *images.Value {
		img := structs.BingImage{
			ContentURL: *image.ContentURL,
			Name:       *image.Name,
		}
		imgs = append(imgs, img)
	}

	return imgs, nil
}
