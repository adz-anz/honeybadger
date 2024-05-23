# TODO list

# Documentation
- [x] Write `README.md`
- [ ] The `marker_settings` module in the API documentation is quite sparse. What do the functions actually do? Update the `Short` & `Long`

# Functionality
- [x] Implement import of API key from ENV var
- [x] Implement Logrus for error handling, info messages, etc
- [ ] Add `config` module for defining default datasets, config keys, etc
- [ ] Implement all API endpoints
    - [x] Auth
    - [x] Boards
	- [ ] Burn Alerts
	- [ ] Columns
	- [ ] Datasets
	- [ ] Dataset Definitions
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
- [ ] The `newBoardsCreateCmd` and `newBoardsUpdateCmd` function are complex due to accepting multiple queries, can this be simplified or the UX improved?

# Polish
- [ ] Banner
- [ ] Add `Example` strings
- [ ] Add custom error handling depending on the server response
	- [ ] Include errors based on the Honeycomb API Errors reference
	- [ ] Include checks for whether the provided API key has the right permissions before sending the request? Maybe?
- [ ] "Prettify" output from commands - colors and layout instead of just a `JSON` blob
