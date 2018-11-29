package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
)

var (
	globalUsers = []map[string]string{
		map[string]string{"name": "user0", "type": "type0"},
		map[string]string{"name": "user1", "type": "type0"},
		map[string]string{"name": "user2", "type": "type1"},
		map[string]string{"name": "user3", "type": "type1"},
	}

	baseURL = "http://localhost:5001/users"
)

func TestMain(m *testing.M) {
	time.Sleep(1 * time.Second) // one second for APIs to spawn
	os.Exit(m.Run())
}

func TestAllUsers(t *testing.T) {
	var allUsers []map[string]string
	err := get(baseURL, &allUsers)
	if err != nil {
		t.Errorf("Could not get all users: %s.", err)
		return
	}

	t.Run("equal users list result", testEqualUsersListResult(globalUsers, allUsers))
}

func testEqualUsersListResult(a, b []map[string]string) func(*testing.T) {
	return func(t *testing.T) {
		if len(a) != len(b) {
			t.Errorf("Wrong length. Expected: %d. Got: %d.", len(a), len(b))
		}
		for _, gu := range a {
			if !listContains(b, gu) {
				t.Errorf("Missing user in get all users: %v.", gu)
			}
		}
	}
}

func TestUsersByType(t *testing.T) {
	for _, test := range []struct {
		t    string
		l, h int
	}{
		{t: "type0", l: 0, h: 2},
		{t: "type1", l: 2, h: 4},
	} {
		var users []map[string]string
		err := get(baseURL+"?type="+test.t, &users)
		if err != nil {
			t.Errorf("Could not get users of type %s: %s.", test.t, err)
			continue
		}

		t.Run(fmt.Sprintf("equal users of type %s list result", test.t), testEqualUsersListResult(globalUsers[test.l:test.h], users))
	}
}

func TestUsersByName(t *testing.T) {
	for _, test := range []struct {
		n string
		i int
	}{
		{n: "user0", i: 0},
		{n: "user1", i: 1},
		{n: "user2", i: 2},
		{n: "user3", i: 3},
	} {
		var user map[string]string
		err := get(baseURL+"?name="+test.n, &user)
		if err != nil {
			t.Errorf("Could not get user with name %s: %s.", test.n, err)
			continue
		}

		t.Run(fmt.Sprintf("equal users with name %s list result", test.n), testEqualUsersListResult(globalUsers[test.i:test.i+1], []map[string]string{user}))
	}
}

func get(query string, result interface{}) error {
	res, err := http.Get(query)
	if err != nil {
		return errors.Wrap(err, "could not get users")
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&result); err != nil {
		return errors.Wrap(err, "could not decode users")
	}

	return nil
}

func listContains(users []map[string]string, user map[string]string) bool {
	for _, x := range users {
		if userEqual(x, user) {
			return true
		}
	}
	return false
}

func userEqual(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}

	return true
}
