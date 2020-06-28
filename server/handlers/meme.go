package handlers

import (
	"bytes"
	"emotion-based-memes/container"
	"emotion-based-memes/services"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/cognitiveservices/face"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

type MemeHandler struct {
	azureService *services.AzureService
}

func NewMemeHandler(container container.ServiceContainer) *MemeHandler {
	return &MemeHandler{
		azureService: container.AzureService,
	}
}

func (mh *MemeHandler) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("file")
	memeURL := r.FormValue("url")
	if err != nil {
		log.Println("error file form")
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

	fmt.Println("Processing...")

	path, err := downloadImage(memeURL)
	if err != nil {
		fmt.Println("error on download image")
	}

	meme, err := os.Open(path)
	if err != nil {
		log.Println("error on open file")
	}
	defer meme.Close()
	memeBytes, _ := ioutil.ReadAll(meme)

	photoAnalysis, err := mh.azureService.DetectEmotion(ioutil.NopCloser(bytes.NewReader(fileBytes)))

	if err != nil {
		log.Println("error found photo analysis", err)
		render.JSON(w, http.StatusBadRequest, nil)
	}

	memeAnalysis, err := mh.azureService.DetectEmotion(ioutil.NopCloser(bytes.NewReader(memeBytes)))
	if err != nil {
		log.Println("error found meme analysis", err)
		render.JSON(w, http.StatusBadRequest, nil)
	}

	meme.Seek(0, 0)
	_, err = createNewPicture(meme, file, memeAnalysis.FaceRectangle, photoAnalysis.FaceRectangle)

	render.JSON(w, http.StatusNoContent, nil)
}

func createNewPicture(meme io.Reader, photo multipart.File, memeCoords *face.Rectangle, photoCoords *face.Rectangle) (string, error) {
	output := "./temp-images/output.png"

	photo.Seek(0, 0)
	target, _, err := image.Decode(photo)
	if err != nil {
		log.Println("error on decoding image", err)
	}

	cropped, err := cutter.Crop(target, cutter.Config{
		Width:  int(*photoCoords.Width),
		Height: int(*photoCoords.Height),
		Anchor: image.Point{
			int(*photoCoords.Left),
			int(*photoCoords.Top),
		},
		Mode:    cutter.TopLeft,
		Options: cutter.Copy,
	})
	if err != nil {
		fmt.Println("error on crop", err)
	}

	resized := resize.Resize(uint(*memeCoords.Width), uint(*memeCoords.Height), cropped, resize.Lanczos3)

	memeImage, _, err := image.Decode(meme)
	if err != nil {
		log.Println("error on convert meme image", err)
	}

	destImage := image.NewRGBA(image.Rect(0, 0, memeImage.Bounds().Dx(), memeImage.Bounds().Dy()))

	draw.Draw(destImage, memeImage.Bounds(), memeImage, image.Point{0, 0}, draw.Src)

	draw.Draw(
		destImage,
		destImage.Bounds(),
		resized,
		image.Point{
			-int(*memeCoords.Left),
			-int(*memeCoords.Top),
		},
		draw.Src,
	)

	f, err := os.Create(output)
	if err != nil {
		log.Println("error on create file")
	}
	defer f.Close()
	err = png.Encode(f, destImage)
	if err != nil {
		log.Println("error on encode")
	}

	return output, nil
}

func downloadImage(url string) (string, error) {
	us := strings.Split(url, ".")
	t := us[len(us)-1]

	path := "./temp-images/image." + t

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return path, nil
}
