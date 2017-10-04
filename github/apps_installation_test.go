// Copyright 2016 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAppsService_ListRepos(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/installation/repositories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"repositories": [{"id":1}]}`)
	})

	opt := &ListOptions{Page: 1, PerPage: 2}
	repositories, _, err := client.Apps.ListRepos(context.Background(), opt)
	if err != nil {
		t.Errorf("Apps.ListRepos returned error: %v", err)
	}

	want := []*Repository{{ID: Int(1)}}
	if !reflect.DeepEqual(repositories, want) {
		t.Errorf("Apps.ListRepos returned %+v, want %+v", repositories, want)
	}
}

func TestAppsService_AddRepo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/app/installations/1/repositories/1", func(w http.ResponseWriter, r *http.Request) {
		// testMethod(t, r, "PUT")
		// testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
		v := new(Repository)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		fmt.Fprint(w, `{"id":1}`)
	})

	repo, _, err := client.Apps.AddRepo(context.Background(), 1, 1)
	if err != nil {
		t.Errorf("Apps.AddRepo returned error: %v", err)
	}

	want := []*Repository{{ID: Int(1)}}
	// want := `{"repositories": [{"id":1}]}`
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("AddRepo returned %+v, want %+v", repo, want)
	}
}

// func TestRepositoriesService_Create_org(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	input := &Repository{Name: String("n")}

// 	mux.HandleFunc("/orgs/o/repos", func(w http.ResponseWriter, r *http.Request) {
// 		v := new(Repository)
// 		json.NewDecoder(r.Body).Decode(v)

// 		testMethod(t, r, "POST")
// 		if !reflect.DeepEqual(v, input) {
// 			t.Errorf("Request body = %+v, want %+v", v, input)
// 		}

// 		fmt.Fprint(w, `{"id":1}`)
// 	})

// 	repo, _, err := client.Repositories.Create(context.Background(), "o", input)
// 	if err != nil {
// 		t.Errorf("Repositories.Create returned error: %v", err)
// 	}

// 	want := &Repository{ID: Int(1)}
// 	if !reflect.DeepEqual(repo, want) {
// 		t.Errorf("Repositories.Create returned %+v, want %+v", repo, want)

func TestAppsService_RemoveRepo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/app/installations/1/repositories/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "Accept", mediaTypeIntegrationPreview)
	})

	_, err := client.Apps.RemoveRepo(context.Background(), 1, 1)
	if err != nil {
		t.Errorf("Apps.RemoveRepo returned error: %v", err)
	}
}
