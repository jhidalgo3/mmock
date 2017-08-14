package translate

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jmartin82/mmock/definition"
)

//HTTP is and adaptor beteewn the http and mock definition.
type HTTP struct {
}

//BuildRequestDefinitionFromHTTP Read the request definition and return a mock request.
func (t HTTP) BuildRequestDefinitionFromHTTP(req *http.Request) definition.Request {

	res := definition.Request{}
	res.Host = strings.Split(req.Host, ":")[0]
	res.Method = req.Method
	res.Path = req.URL.Path
	res.Headers = make(definition.Values)
	for header, values := range req.Header {
		if header != "Cookie" {
			res.Headers[header] = values
		}
	}

	res.Cookies = make(definition.Cookies)
	for _, cookie := range req.Cookies() {
		res.Cookies[cookie.Name] = cookie.Value
	}

	res.QueryStringParameters = make(definition.Values)
	for name, values := range req.URL.Query() {
		res.QueryStringParameters[name] = values
	}

	body, _ := ioutil.ReadAll(req.Body)
	res.Body = string(body)
	return res
}

//WriteHTTPResponseFromDefinition read a mock response and write a http response.
func (t HTTP) WriteHTTPResponseFromDefinition(fr *definition.Response, w http.ResponseWriter, r *http.Request) {
	fmt.Println(fr.Body)
	if strings.HasPrefix(fr.Body, "@") {
		file := fr.Body[1:]
		fmt.Println(file)
		serveFile(file, w, r)
	} else {
		for header, values := range fr.Headers {
			for _, value := range values {
				w.Header().Add(header, value)
			}

		}
		if len(fr.Cookies) > 0 {
			cookies := []string{}
			for cookie, value := range fr.Cookies {
				cookies = append(cookies, fmt.Sprintf("%s=%s", cookie, value))
			}
			w.Header().Add("Set-Cookie", strings.Join(cookies, ";"))
		}

		w.WriteHeader(fr.StatusCode)

		fmt.Printf("BODY %s \n", fr.Body)
		io.WriteString(w, fr.Body)
	}

}

func serveFile(f string, w http.ResponseWriter, r *http.Request) {
	/*data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+"fileName.here")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	http.ServeContent(w, r,f, time.Now(), bytes.NewReader(data))*/
	http.ServeFile(w, r, f)
}
