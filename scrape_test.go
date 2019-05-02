package shovel

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func server() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(`
		<html>
			<body>
				<div class="item">
					<p class="name">Matt</p>
					<p class="age">33</p>
				</div>
				<div class="item">
					<p class="name">Bob</p>
					<p class="age">45</p>
				</div>
			</body>
		</html>`))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe("localhost:10123", nil)
}

func TestScrape(t *testing.T) {

	go server()

	manifest := Manifest{
		URL:                "http://localhost:10123",
		SingleItemSelector: ".item",
		AttributeSelectors: map[string]string{
			"name": ".name",
			"age":  ".age",
		},
	}

	result := ScrapePaginated(manifest, func(int, []map[string]string) (bool, string) { return true, "" })

	expected := []map[string]string{
		{
			"name": "Matt",
			"age":  "33",
		},
		{
			"name": "Bob",
			"age":  "45",
		},
	}

	if cmp.Equal(result, expected) != true {
		t.Fail()
	}
}
