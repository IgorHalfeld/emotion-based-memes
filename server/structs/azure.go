package structs

type FaceAttributes struct {
	anger     float32
	content   float32
	disgust   float32
	feat      float32
	happiness float32
	neutral   float32
	sadness   float32
	surprise  float32
}

type FaceResponse struct {
	faceAttributes struct {
		emotion FaceAttributes
	}
}

type BingImage struct {
	ContentURL string `json:"content_url"`
	Name       string `json:"name"`
	BingID     string `json:"bing_id"`
}
