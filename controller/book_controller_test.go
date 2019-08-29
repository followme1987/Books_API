package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/followme1987/bookAPI/controller"
	"github.com/followme1987/bookAPI/model"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"

	"strings"

	_ "github.com/lib/pq"
)

var _ = Describe("BookController", func() {
	Context("when send a request to retrieve all the books", func() {
		var server *httptest.Server
		BeforeEach(func() {
			c := controller.Controller{}
			server = httptest.NewServer(c.GetBooks(nil))
		})
		AfterEach(func() {
			server.Close()
		})
		It("should return HTTP 200 if the request path is correct", func() {
			subUrl := "/books"
			url := server.URL + subUrl

			reader := strings.NewReader("")
			req, err := http.NewRequest("GET", url, reader)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res, err := http.DefaultClient.Do(req)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer res.Body.Close()
			gomega.Expect(res.StatusCode).To(gomega.BeNumerically("==", http.StatusOK))
		})
	})

	Context("when send a request to retrieve a book by Id", func() {
		var server *httptest.Server
		BeforeEach(func() {
			c := controller.Controller{}
			server = httptest.NewServer(c.GetBookById(nil))
		})
		AfterEach(func() {
			server.Close()
		})
		It("should return HTTP 200 if the request path is correct", func() {
			subUrl := "/book/1"
			url := server.URL + subUrl

			reader := strings.NewReader("")
			req, err := http.NewRequest("GET", url, reader)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res, err := http.DefaultClient.Do(req)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer res.Body.Close()
			gomega.Expect(res.StatusCode).To(gomega.BeNumerically("==", http.StatusOK))
		})

		It("should return HTTP 400 if the request id is not valid", func() {
			subUrl := "/book/test-id"
			url := server.URL + subUrl

			reader := strings.NewReader("")
			req, err := http.NewRequest("GET", url, reader)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res, err := http.DefaultClient.Do(req)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer res.Body.Close()
			gomega.Expect(res.StatusCode).To(gomega.BeNumerically("==", http.StatusBadRequest))
		})
	})

	Context("when send a request to delete a book by Id", func() {
		var server *httptest.Server
		BeforeEach(func() {
			c := controller.Controller{}
			server = httptest.NewServer(c.DeleteBookById(nil))
		})
		AfterEach(func() {
			server.Close()
		})
		It("should return HTTP 200 if the request path is correct", func() {
			subUrl := "/book/1"
			url := server.URL + subUrl

			reader := strings.NewReader("")
			req, err := http.NewRequest("DELETE", url, reader)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res, err := http.DefaultClient.Do(req)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer res.Body.Close()
			gomega.Expect(res.StatusCode).To(gomega.BeNumerically("==", http.StatusOK))
		})

		It("should return HTTP 400 if the request id is not valid", func() {
			subUrl := "/book/test-id"
			url := server.URL + subUrl

			reader := strings.NewReader("")
			req, err := http.NewRequest("DELETE", url, reader)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res, err := http.DefaultClient.Do(req)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer res.Body.Close()
			gomega.Expect(res.StatusCode).To(gomega.BeNumerically("==", http.StatusBadRequest))
		})
	})

	Context("when send a request to update books", func() {
		var server *httptest.Server
		BeforeEach(func() {
			c := controller.Controller{}
			server = httptest.NewServer(c.UpdateBooks(nil))
		})
		AfterEach(func() {
			server.Close()
		})
		It("should return HTTP 200 if the request path and body payload are correct", func() {
			subUrl := "/books"
			url := server.URL + subUrl

			book := model.Book{
				Id:    1,
				Title: "test book",
				Year:  "1900",
			}

			requestBytes, err := json.Marshal(&book)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBytes))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res, err := http.DefaultClient.Do(req)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer res.Body.Close()
			gomega.Expect(res.StatusCode).To(gomega.BeNumerically("==", http.StatusOK))
		})

		It("should return HTTP 400 if the payload is not correct", func() {
			subUrl := "/book"
			url := server.URL + subUrl
			reader := strings.NewReader("")
			req, err := http.NewRequest("PUT", url, reader)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res, err := http.DefaultClient.Do(req)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer res.Body.Close()
			gomega.Expect(res.StatusCode).To(gomega.BeNumerically("==", http.StatusBadRequest))
		})
	})

	Context("when send a request to add a book", func() {
		var server *httptest.Server
		BeforeEach(func() {
			c := controller.Controller{}
			server = httptest.NewServer(c.AddBook(nil))
		})
		AfterEach(func() {
			server.Close()
		})
		It("should return HTTP 200 if the request path and body payload are correct", func() {
			subUrl := "/book"
			url := server.URL + subUrl

			book := model.Book{
				Title: "test book",
				Year:  "1900",
			}

			requestBytes, err := json.Marshal(&book)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBytes))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res, err := http.DefaultClient.Do(req)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer res.Body.Close()
			gomega.Expect(res.StatusCode).To(gomega.BeNumerically("==", http.StatusOK))
		})

		It("should return HTTP 400 if the payload is not correct", func() {
			subUrl := "/book"
			url := server.URL + subUrl
			reader := strings.NewReader("")
			req, err := http.NewRequest("POST", url, reader)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res, err := http.DefaultClient.Do(req)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer res.Body.Close()
			gomega.Expect(res.StatusCode).To(gomega.BeNumerically("==", http.StatusBadRequest))
		})
	})
})
