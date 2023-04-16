package hermes

import (
	"strings"
)

// SearchWithSpaces function with lock
func (fts *FTS) SearchWithSpaces(query string, limit int, strict bool, keys map[string]bool) ([]map[string]string, []int) {
	fts.mutex.RLock()
	defer fts.mutex.RUnlock()
	return fts._SearchWithSpaces(query, limit, strict, keys)
}

// Search for multiple words
func (fts *FTS) _SearchWithSpaces(query string, limit int, strict bool, keys map[string]bool) ([]map[string]string, []int) {
	// Split the query into words
	var words []string = strings.Split(strings.TrimSpace(query), " ")

	// If the words array is empty
	switch {
	case len(words) == 0:
		return []map[string]string{}, []int{}
	case len(words) == 1:
		return fts.Search(words[0], limit, strict)
	}

	// Create an array to store the result
	var result []map[string]string = []map[string]string{}

	// Loop through the words and get the indices that are common
	for i := 0; i < len(words); i++ {
		// Search for the query inside the cache
		var queryResult, _ = fts.Search(words[i], limit, strict)

		// Iterate over the result
		for j := 0; j < len(queryResult); j++ {
			// Iterate over the keys and values for the json data for that index
			for key, value := range queryResult[j] {
				switch {
				case !keys[key]:
					continue
				case _Contains(value, query):
					result = append(result, queryResult[j])
				}
			}
		}
	}

	// Return the result
	return result, []int{}
}

// SearchInJsonWithKey function with lock
func (fts *FTS) SearchInJsonWithKey(query string, key string, limit int) []map[string]string {
	fts.mutex.RLock()
	defer fts.mutex.RUnlock()
	return fts._SearchInJsonWithKey(query, key, limit)
}

// SearchInJsonWithKey function
func (fts *FTS) _SearchInJsonWithKey(query string, key string, limit int) []map[string]string {
	// Define variables
	var result []map[string]string = []map[string]string{}

	// Iterate over the query result
	for i := 0; i < len(fts.json); i++ {
		switch {
		case _ContainsIgnoreCase(fts.json[i][key], query):
			result = append(result, fts.json[i])
		}
	}

	// Return the result
	return result
}

// SearchInJson function with lock
func (fts *FTS) SearchInJson(query string, limit int, keys map[string]bool) []map[string]string {
	fts.mutex.RLock()
	defer fts.mutex.RUnlock()
	return fts._SearchInJson(query, limit, keys)
}

// _SearchInJson function
func (fts *FTS) _SearchInJson(query string, limit int, keys map[string]bool) []map[string]string {
	// Define variables
	var result []map[string]string = []map[string]string{}

	// Iterate over the query result
	for i := 0; i < len(fts.json); i++ {
		// Iterate over the keys and values for the json data for that index
		for key, value := range fts.json[i] {
			switch {
			case !keys[key]:
				continue
			case _ContainsIgnoreCase(value, query):
				result = append(result, fts.json[i])
			}
		}
	}

	// Return the result
	return result
}

// Search function with lock
func (fts *FTS) Search(query string, limit int, strict bool) ([]map[string]string, []int) {
	fts.mutex.RLock()
	defer fts.mutex.RUnlock()
	return fts._Search(query, limit, strict)
}

// Search for a single query
func (fts *FTS) _Search(query string, limit int, strict bool) ([]map[string]string, []int) {
	// If the query is empty
	if len(query) == 0 {
		return []map[string]string{}, []int{}
	}

	// Define variables
	var (
		result  []map[string]string = []map[string]string{}
		indices []int               = make([]int, len(fts.json))
	)

	// If the user wants a strict search, just return the result
	// straight from the cache
	if strict {
		// Check if the query is in the cache
		if _, ok := fts.cache[query]; !ok {
			return result, indices
		}

		// Loop through the indices
		indices = fts.cache[query]
		for i := 0; i < len(indices); i++ {
			result = append(result, fts.json[indices[i]])
		}

		// Return the result
		return result, indices
	}

	// Loop through the cache keys
	for i := 0; i < len(fts.keys); i++ {
		switch {
		case len(result) >= limit:
			return result, indices
		case !_Contains(fts.keys[i], query):
			continue
		}

		// Loop through the cache indices
		for j := 0; j < len(fts.cache[fts.keys[i]]); j++ {
			var index int = fts.cache[fts.keys[i]][j]
			if indices[index] == -1 {
				continue
			}

			// Else, append the index to the result
			result = append(result, fts.json[index])
			indices[index] = -1
		}
	}

	// Return the result
	return result, indices
}
