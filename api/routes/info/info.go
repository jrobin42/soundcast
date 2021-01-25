package info

import (
	"soundcast/api/interfaces/data"

	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

type jsonResponse map[string]interface{}

// matchUserAgent return true when one of the regex specified in
// DbElement matches the user agent passed as parameter
func matchUserAgent(ua string) data.Matches {
	return func(element data.DbElement) bool {
		value, found := element["user_agents"]
		if found {
			switch typedValue := value.(type) {
			case []interface{}:
				for _, regex := range typedValue {
					re := regexp2.MustCompile(regex.(string), 0)
					matched, _ := re.MatchString(ua)
					if matched {
						return true
					}
				}
			}
		}
		return false
	}
}

// generateResponse update the jsonRespose template by looking for its keys in dbElement, and update its values if found
func generateResponse(template jsonResponse, data data.DbElement) jsonResponse {
	for tKey := range template {
		value, found := data[tKey]
		if found {
			template[tKey] = value
		}
	}
	return template
}

// RequestHandler for route "/info?ua={url_encoded_ua}"
// Search in the database the fist entry that match the user agent given as parameter
// It then returns the informations looked for from this entry
func RequestHandler(c *gin.Context) {
	ua := c.Query("ua")
	df := c.MustGet("dataFinder").(data.Finder)

	result := df.First(matchUserAgent(ua))
	if result == nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "user agent does not match any known entity",
		})
		return
	}
	template := jsonResponse{
		"app":    "unknown",
		"device": "unknown",
		"bot":    false,
	}
	response := generateResponse(template, result)
	c.JSON(200, response)
}
