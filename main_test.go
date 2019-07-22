package main

import ( 	//"fmt"
			"strings"
			"io/ioutil"
			"testing"
			"net/http"
			"net/http/httptest"
		)

func TestArticlePage(t *testing.T) {
	
	//Create a recorder for capturing results of our test requests
	recorder := httptest.NewRecorder()
	
	//Simulate a PUT request to add testable article data
	client := &http.Client{}
	request, err := http.NewRequest("PUT", "http://127.0.0.1:9090/articles/testarticle", strings.NewReader("Sample article content test"))
	
	//Error check on PUT request setup
	if err != nil {
		t.Fatal(err)
	}
	
	request.SetBasicAuth("admin", "admin")
	request.ContentLength = 27
	response, err := client.Do(request)
	
	//Error check on PUT request execution
	if err != nil {
		t.Fatal(err)
	} else {
		//defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		
		//Error check on response body read-in
		if err != nil {
			t.Fatal(err)
		}
		
		if(len(string(contents)) != 0) {
			t.Errorf("Expected an empty return for PUT request.")
		}
		if status := response.StatusCode; status != http.StatusOK && status != http.StatusCreated { 
			t.Errorf("Received PUT request HTTP status code: %v, but expected %v or %v", status, http.StatusOK, http.StatusCreated);
		}
	}
	
	//GET request to main articles page to ensure proper status response
	request, err = http.NewRequest("GET", "http://127.0.0.1:9090/articles/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ArticlePage)
	handler.ServeHTTP(recorder, request)
	
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Articles Main: Wrong Initial GET HTTP status code: Received %v expected %v", status, http.StatusOK)
	}
	
	//GET Request on individual article page to ensure proper status response
	//Request to be passed to the handler
	request, err = http.NewRequest("GET", "/articles/testarticle", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Article Individual: Wrong Initial GET HTTP status code: Received %v expected %v", status, http.StatusOK)
	}
	
}

