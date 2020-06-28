package handlers

import (
	"bytes"
	"emotion-based-memes/container"
	"emotion-based-memes/services"
	"emotion-based-memes/structs"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/profiles/preview/cognitiveservices/face"
	renderPkg "github.com/unrolled/render"
)

var render = renderPkg.New()

type FaceHandler struct {
	azureService *services.AzureService
}

func NewFaceHandler(container container.ServiceContainer) *FaceHandler {
	return &FaceHandler{
		azureService: container.AzureService,
	}
}

func (fh *FaceHandler) Analyze(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Emotion string              `json:"emotion"`
		Images  []structs.BingImage `json:"images"`
	}

	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("error on get file from form")
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		log.Println("error on create temp file")
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("error on read buff")
	}
	tempFile.Write(fileBytes)

	log.Println("Processing...")

	faceAnalysis, _ := fh.azureService.DetectEmotion(ioutil.NopCloser(bytes.NewReader(fileBytes)))
	search := findEmotionPredominat(*faceAnalysis.FaceAttributes)

	v, _ := fh.azureService.SearchImages(search)

	render.JSON(w, 200, response{
		Emotion: search,
		Images:  v,
	})
}

func findEmotionPredominat(faceAttributes face.Attributes) string {
	type emotion struct {
		Name    string
		Percent float64
	}

	lastEmotion := emotion{
		Name:    "raiva",
		Percent: *faceAttributes.Emotion.Anger,
	}

	if *faceAttributes.Emotion.Contempt > lastEmotion.Percent {
		lastEmotion.Name = "desprezo"
		lastEmotion.Percent = *faceAttributes.Emotion.Contempt
	}

	if *faceAttributes.Emotion.Disgust > lastEmotion.Percent {
		lastEmotion.Name = "nojo"
		lastEmotion.Percent = *faceAttributes.Emotion.Disgust
	}

	if *faceAttributes.Emotion.Fear > lastEmotion.Percent {
		lastEmotion.Name = "medo"
		lastEmotion.Percent = *faceAttributes.Emotion.Fear
	}

	if *faceAttributes.Emotion.Happiness > lastEmotion.Percent {
		lastEmotion.Name = "felicidade"
		lastEmotion.Percent = *faceAttributes.Emotion.Happiness
	}

	if *faceAttributes.Emotion.Neutral > lastEmotion.Percent {
		lastEmotion.Name = "neutro"
		lastEmotion.Percent = *faceAttributes.Emotion.Neutral
	}

	if *faceAttributes.Emotion.Sadness > lastEmotion.Percent {
		lastEmotion.Name = "tristeza"
		lastEmotion.Percent = *faceAttributes.Emotion.Sadness
	}

	if *faceAttributes.Emotion.Surprise > lastEmotion.Percent {
		lastEmotion.Name = "surpresa"
		lastEmotion.Percent = *faceAttributes.Emotion.Surprise
	}

	return lastEmotion.Name
}
