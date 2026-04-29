package linux

import (
	"strings"
)

func parseID(id string) string {
	ids := strings.Fields(id)
	firstID := ids[0]
	cleanedID := strings.ReplaceAll(firstID, "\"", "")
	id = cleanedID
	return id
}

func GetID(osMap map[string]string) string {
	id := osMap["ID"]

	if len(id) > 1 {
		id = parseID(id)
	}
	return id
}

func GetIDLike(osMap map[string]string) string {
	idLike := osMap["ID_LIKE"]

	if idLike == "" {
		idLike = GetID(osMap)
	}
	return idLike
}
