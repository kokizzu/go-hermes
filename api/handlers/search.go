package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	Hermes "github.com/realTristan/Hermes"
	Utils "github.com/realTristan/Hermes/api/utils"
)

// Search is a handler function that returns a fiber context handler function for searching the cache.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that searches the cache using the query, limit, strict, and schema parameters provided in the query string and returns a JSON-encoded string of the search results or an error message if the search fails or if the parameters are not provided.
func Search(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			strict bool
			query  string
			limit  int
		)

		// Get the query from the url params
		if query = ctx.Query("query"); len(query) == 0 {
			return ctx.Send(Utils.Error("query not provided"))
		}

		// Get the limit from the url params
		if err := Utils.GetLimitParam(ctx, &limit); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the strict from the url params
		if err := Utils.GetStrictParam(ctx, &strict); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Search for the query
		if res, err := c.Search(Hermes.SearchParams{
			Query:  query,
			Limit:  limit,
			Strict: strict,
		}); err != nil {
			return ctx.Send(Utils.Error(err))
		} else if data, err := json.Marshal(res); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(data)
		}
	}
}

// SearchOneWord is a handler function that returns a fiber context handler function for searching the cache for a single word.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that searches the cache for a single word using the query, limit, and strict parameters provided in the query string and returns a JSON-encoded string of the search results or an error message if the search fails or if the parameters are not provided.
func SearchOneWord(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			strict bool
			query  string
			limit  int
		)

		// Get the query from the url params
		if query = ctx.Query("query"); len(query) == 0 {
			return ctx.Send(Utils.Error("invalid query"))
		}

		// Get the limit from the url params
		if err := Utils.GetLimitParam(ctx, &limit); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the strict from the url params
		if err := Utils.GetStrictParam(ctx, &strict); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Search for the query
		if res, err := c.SearchOneWord(Hermes.SearchParams{
			Query:  query,
			Limit:  limit,
			Strict: strict,
		}); err != nil {
			return ctx.Send(Utils.Error(err))
		} else if data, err := json.Marshal(res); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(data)
		}
	}
}

// SearchValues is a handler function that returns a fiber context handler function for searching the cache for values.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that searches the cache for values using the query, limit, and schema parameters provided in the query string and returns a JSON-encoded string of the search results or an error message if the search fails or if the parameters are not provided.
func SearchValues(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			query  string
			limit  int
			schema map[string]bool
		)

		// Get the query from the url params
		if query = ctx.Query("query"); len(query) == 0 {
			return ctx.Send(Utils.Error("invalid query"))
		}

		// Get the limit from the url params
		if err := Utils.GetLimitParam(ctx, &limit); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Get the schema from the url params
		if err := Utils.GetSchemaParam(ctx, &schema); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Search for the query
		if res, err := c.SearchValues(Hermes.SearchParams{
			Query:  query,
			Limit:  limit,
			Schema: schema,
		}); err != nil {
			return ctx.Send(Utils.Error(err))
		} else if data, err := json.Marshal(res); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(data)
		}
	}
}

// SearchWithKey is a handler function that returns a fiber context handler function for searching the cache with a specific key.
// Parameters:
//   - c (*Hermes.Cache): A pointer to a Hermes.Cache struct.
//
// Returns:
//   - func(ctx *fiber.Ctx) error: A fiber context handler function that searches the cache with a specific key using the query and limit parameters provided in the query string and returns a JSON-encoded string of the search results or an error message if the search fails or if the parameters are not provided.
func SearchWithKey(c *Hermes.Cache) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var (
			key   string
			query string
			limit int
		)

		// Get the query from the url params
		if query = ctx.Query("query"); len(query) == 0 {
			return ctx.Send(Utils.Error("invalid query"))
		}

		// Get the key from the url params
		if key = ctx.Query("key"); len(key) == 0 {
			return ctx.Send(Utils.Error("invalid key"))
		}

		// Get the limit from the url params
		if err := Utils.GetLimitParam(ctx, &limit); err != nil {
			return ctx.Send(Utils.Error(err))
		}

		// Search for the query
		if res, err := c.SearchWithKey(Hermes.SearchParams{
			Key:   key,
			Query: query,
			Limit: limit,
		}); err != nil {
			return ctx.Send(Utils.Error(err))
		} else if data, err := json.Marshal(res); err != nil {
			return ctx.Send(Utils.Error(err))
		} else {
			return ctx.Send(data)
		}
	}
}
