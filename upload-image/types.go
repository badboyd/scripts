package main


// ImgPutRequest ...
type ImgPutRequest struct {
	Image     []byte `json:"image" validate:"required,v_image"`
	Thumbnail []byte `json:"thumbnail" validate:"required,v_thumbnail"`
	Width     int    `json:"width" validate:"v_integer"`
	Height    int    `json:"height" validate:"v_integer"`
	Length    int    `json:"length" validate:"v_integer"`
}

// ImgPutResponse ...
type ImgPutResponse struct {
	ImageName string `json:"file_name,omitempty"`
	Digest    string `json:"digest,omitempty"`
}