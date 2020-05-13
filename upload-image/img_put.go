package main

// ImgPut insert new image
func (c *Client) ImgPut(in *ImgPutRequest, out *ImgPutResponse) error {
	req, err := c.newRequest("POST", "/v1/img", in)
	if err != nil {
		return err
	}

	_, err = c.do(req, out)
	return err
}
