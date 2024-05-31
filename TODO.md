# TODO list

# Documentation
- [x] Write `README.md`
	- [x] Add Types to flag documentation
- [ ] The `marker_settings` module in the API documentation is quite sparse. What do the functions actually do? Update the `Short` & `Long`
- [ ] Add `Example` strings to `README.md` and `cobra` commands

# Functionality
- [x] Implement import of API key from ENV var
- [x] Implement Logrus for error handling, info messages, etc
- [x] Add "dry run" functionality
- [ ] Add `config` module for defining default datasets, config keys, etc
- [ ] Implement all API endpoints
    - [x] Auth
    - [x] Boards
	- [ ] Burn Alerts
	- [ ] Columns
	- [x] Datasets
	- [x] Dataset Definitions
	- [ ] Derived Columns
	- [ ] Events
	- [ ] Kinesis Events
	- [x] Markers
	- [x] Marker Settings
	- [ ] Queries
	- [ ] Query Annotations
	- [ ] Query Data
	- [ ] Recipients
	- [ ] SLOs
	- [ ] Triggers

# Improvements
- [x] The `newBoardsCreateCmd` and `newBoardsUpdateCmd` function are complex due to accepting multiple queries, can this be simplified or the UX improved?
	> [!NOTE]
	> The `create` and `update` commands now no longer allow CRUD actions for queries. Queries can be affected using the `add_query`, `update_query`, and `delete_query` commands.
	- [x] Write `newBoardsAddQueryCmd()`
	- [ ] Write `newBoardsUpdateQueryCmd()`
	- [ ] Write `newBoardsDeleteQueryCmd()`

# Polish
- [X] Banner
- [x] Implement an improved version of the `payload.Execute()` function to handle marshalling to intended types.
- [ ] Implement `enum` checks, for `strings` that must be one of X values
- [ ] Implement `string` length checks for `strings` that must be of a specific length
- [ ] Add custom error handling depending on the server response
	- [ ] Include errors based on the Honeycomb API Errors reference
	- [ ] Include checks for whether the provided API key has the right permissions before sending the request? Maybe?
- [ ] "Prettify" output from commands - colors and layout instead of just a `JSON` blob
- [ ] Add interactive terminal for all CRUD activities.
