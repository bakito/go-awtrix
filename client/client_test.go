package client

import (
	"errors"
	"net/http"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var cl *client
	BeforeEach(func() {
		cl = &client{}
	})
	Context("client", func() {
		Context("handleOKResult", func() {
			var (
				resp  *resty.Response
				errIn error
			)
			BeforeEach(func() {
				resp = &resty.Response{
					RawResponse: &http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
					},
				}
				resp.SetBody([]byte("OK"))
			})
			It("should be successful", func() {
				err := cl.handleOKResult(resp, nil)
				Ω(err).ShouldNot(HaveOccurred())
			})
			It("should return an error if not nil", func() {
				errIn = errors.New("test")
				err := cl.handleOKResult(resp, errIn)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(Equal(errIn))
			})
			It("should return an error if status ode is not 200", func() {
				resp.RawResponse.StatusCode = http.StatusNotFound
				resp.RawResponse.Status = http.StatusText(http.StatusNotFound)
				err := cl.handleOKResult(resp, nil)
				Ω(err).Should(HaveOccurred())
				Ω(err.Error()).Should(Equal("unexpected status code: 404 (Not Found)"))
			})
			It("should return an error if body is not 'OK'", func() {
				resp.SetBody([]byte("foo"))
				err := cl.handleOKResult(resp, nil)
				Ω(err).Should(HaveOccurred())
				Ω(err.Error()).Should(Equal("unexpected response: 'foo'"))
			})
		})
	})
})
